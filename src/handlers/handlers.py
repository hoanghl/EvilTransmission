import base64
import tempfile
import traceback

import cv2
import numpy as np
from fastapi import HTTPException, UploadFile
from loguru import logger
from PIL import Image

from . import utils

THETA = 0.9

feature_matcher = utils.Features()


def handle_image(file: UploadFile) -> dict:
    """Extract feature from image

    Args:
        file (UploadFile): Requested file

    Returns:
        dict: Return info
    """

    payload = file.file

    img_trans = utils.transform_img(Image.open(payload))
    feat = utils.extract_feat(img_trans)
    feat = feat.squeeze()

    is_matching = feature_matcher.detect_match(feat, "image", THETA)

    if is_matching is False:
        feature_matcher.add_feat(feat, "image", file.filename)

    feat_encoded = base64.b64encode(feat.tobytes())

    return {
        "ismatching": is_matching,
        "embedding": feat_encoded,
    }


def handle_video(file: UploadFile, delta: int = 20) -> dict:
    """Extract feature from video

    Args:
        file (bytes): _description_

    Returns:
        dict: _description_
    """

    payload = file.file

    with tempfile.NamedTemporaryFile() as file_tmp:
        # Write data to temporary file
        file_tmp.write(payload.read())
        file_tmp.seek(0)

        logger.debug(f"File name: {file_tmp.name}")

        # Load video using opencv
        frames = []
        try:
            cap = cv2.VideoCapture(file_tmp.name)

            n = 0
            while cap.isOpened():
                if n % delta == 0:
                    _, frame = cap.read()
                    if frame is None:
                        break

                    img = Image.fromarray(frame)
                    img_trans = utils.transform_img(img)

                    frames.append(img_trans.squeeze(0))
                    if len(frames) > 20:
                        cap.release()

                        break

                n = (n + 1) % delta
        except Exception:
            raise HTTPException(500, detail=traceback.format_exc())
        finally:
            if "cap" in locals():
                cap.release()

    video_trans = np.stack(frames)
    final_feat = utils.extract_feat(video_trans)
    final_feat = np.mean(final_feat, axis=0)
    logger.debug(f"shape: final_feat: {final_feat.shape}")

    is_matching = feature_matcher.detect_match(final_feat, "video", THETA)

    if is_matching is False:
        feature_matcher.add_feat(final_feat, "video", file.filename)

    final_feat_encoded = base64.b64encode(final_feat.tobytes())

    return {
        "ismatching": is_matching,
        "embedding": final_feat_encoded,
    }
