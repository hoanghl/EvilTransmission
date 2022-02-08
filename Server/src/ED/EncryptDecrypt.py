
from cryptography.hazmat.primitives.ciphers import Cipher, algorithms, modes
from cryptography.hazmat.primitives.asymmetric import rsa
from cryptography.hazmat.primitives import serialization

class EncryptDecrypt:
    def __init__(self) -> None:
        pass

    def encryptRSA(self):
        pass
    
    def decryptRSA(self):
        pass

    def getServerRSAPublicKey(self):
        """Return server's RSA public key
        """
        pass

    def loadClientRSAPublicKey(self, clientPubKey: str):
        """Load client's RSA public key. Input must obey PEM format.

        Args:
            clientPubKey (str): String sent by client
        """

    def createRSAKey(self):

    def encryptAES(self, data: bytes) -> str:
        """Encrypt

        Returns:
            str: [description]
        """
        pass

    def decryptAES(self):
        pass
