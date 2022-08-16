import json

from src import *


class Storage:
    def __init__(self, root="tests/") -> None:
        logger.info("Initialize: Storage")
        logger.info("== Load saved")

        self.hashes = self.load_hash()
        self.root = root

    def load_hash(self, path="saved_hash.json") -> dict:
        """Load saved hash strings

        Args:
            path (str, optional): Path to file containing saved hashes. Defaults to "saved_hash.json".

        Returns:
            dict: dict of hash of file and corresponding path
        """

        if not osp.isfile(path):
            logger.error(f"Saved file not found: {path}")

            return dict()

        with open(path) as f:
            hashes = json.load(f)

        return hashes

    async def back_up(self, path="saved_hash.json"):
        """Back up hash periodically

        Args:
            path (str, optional): path of backed up file. Defaults to "saved_hash.json".
        """
        while True:
            await asyncio.sleep(60)

            logger.info("Periodically back up")

            ## Back up to file
            if not osp.isfile(path):
                logger.error(f"Old file not found, create new: {path}")

            with open(path, "w+") as f:
                json.dump(self.hashes, f, indent=2)

    def check_hash(self, hash_bytes: bytes):
        return hash_bytes.hex() in self.hashes

    def store_file(self, data: bytes):
        hash_str = get_sha256(data)

        print(f"hash file: {hash_str}")

        file_path = osp.join(self.root, f"{len(self.hashes)}.jpg")

        ## Add hash to current self.hashes
        self.hashes[hash_str] = file_path

        ## Save file
        with open(file_path, "wb+") as f:
            f.write(data)
