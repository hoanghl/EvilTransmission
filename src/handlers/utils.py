from datetime import datetime
from pathlib import Path
from typing import Literal, Union

import numpy as np
import onnxruntime
from loguru import logger
from PIL import Image

# TODO: HoangLe [Aug-03]: Replace this hardcode by ENV variable
PATH_DIR_FEAT_IMG = "feats/imgage"
PATH_DIR_FEAT_VID = "feats/video"
PATH_ONNX_MODEL = "res/resnet34_custom_infer_quan.onnx"

ort_session = onnxruntime.InferenceSession(PATH_ONNX_MODEL)

# transform = transforms.Compose(
#     [
#         transforms.Resize((224, 224)),
#         transforms.ToTensor(),
#         transforms.Normalize(
#             mean=[0.4914, 0.4822, 0.4465],
#             std=[0.2023, 0.1994, 0.2010],
#         ),
#     ]
# )


def transform_img(img, dtype: type = np.float16):
    img = img.resize((224, 224), Image.LANCZOS)

    img = np.asarray(img).astype(dtype)

    mean = [0.485, 0.456, 0.406]
    std = [0.229, 0.224, 0.225]
    img = img.transpose([2, 0, 1])
    for channel in range(img.shape[0]):
        img[channel, :, :] = (img[channel, :, :] / 255 - mean[channel]) / std[channel]
    img = np.expand_dims(img, axis=0)

    return img


def extract_feat(img_trans: np.ndarray) -> np.ndarray:
    ort_inps = {ort_session.get_inputs()[0].name: img_trans}
    ort_outs = ort_session.run(None, ort_inps)

    return ort_outs[0]


def _gen_random_name():
    return datetime.now().strftime(r"%Y%m%d_%H%M%S")


class Features:
    def __init__(self) -> None:
        self._feats_img = [np.load(path) for path in Path(PATH_DIR_FEAT_IMG).glob("*.npy")]
        self._feats_vid = [np.load(path) for path in Path(PATH_DIR_FEAT_VID).glob("*.npy")]

    def add_feat(self, feat: np.ndarray, feat_type: Literal["image", "video"], file_name: Union[str, None]):
        if feat_type == "image":
            self._feats_img.append(feat)
        elif feat_type == "video":
            self._feats_vid.append(feat)
        else:
            raise NotImplementedError()

        if file_name is None:
            file_name = _gen_random_name()
        else:
            file_name = Path(file_name).stem
        path_feat = Path(f"feats/{feat_type}/{file_name}.npy")
        path_feat.parent.mkdir(parents=True, exist_ok=True)
        np.save(path_feat, feat)

    def detect_match(self, feat: np.ndarray, feat_type: Literal["image", "video"], theta: float = 0.9) -> bool:
        def _get_norm(x: np.ndarray, eps: float = 1e-8):
            normed = np.clip(np.linalg.norm(x, axis=-1, keepdims=True), eps, None)

            return normed

        res = self._feats_img if feat_type == "image" else self._feats_vid
        if res == []:
            return False

        res_stacked = np.stack(res)
        res_norm = res_stacked / _get_norm(res_stacked)

        # Get cosine similarity
        feat_norm = feat / _get_norm(feat)
        if feat.ndim < 2:
            feat = feat[None, :]

        logger.debug(f"shape: feat_norm: {feat_norm.shape}")
        logger.debug(f"shape: res_norm: {res_norm.shape}")

        sim = np.squeeze(feat_norm @ res_norm.T)

        # Detect matching
        is_matching = np.sum(sim >= theta).item() > 0

        return is_matching
