
from cryptography.hazmat.primitives import serialization
from cryptography.hazmat.primitives.asymmetric import rsa
from cryptography.hazmat.primitives.ciphers import Cipher, algorithms, modes


class EncryptDecrypt:
    def __init__(self) -> None:
        self.rsa_pubkey, self.rsa_privatekey = None, None
        self.fernetkey = None

        self._gen_keypairRSA()

    def _rsa_gen_keypair(self):
        """Create public and private key for RSA-384 and update it to self.rsa_pubkey and self.rsa_privatekey
        """
        pass

    def _fernet_gen_key(self):
        """Generate symmetric key for Fernet. Only run if self.fernetkey is None
        """

        if self.fernetkey is None:
            ## Generate Fernet key
            pass

    def _update_fernet_key(self, key):
        pass

