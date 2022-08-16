import hashlib


class Buffer:
    def __init__(self):
        self._buff = bytes()

    def __add__(self, item: bytes):

        self._buff += item

        return self

    def __len__(self):
        return len(self._buff)

    def __getitem__(self, idx: slice):
        ret = None

        ret = self._buff[idx]
        self._buff = self._buff[idx.stop :]

        return ret


def get_sha256(secret: bytes) -> str:
    sha256 = hashlib.sha256()

    sha256.update(secret)

    return sha256.hexdigest()
