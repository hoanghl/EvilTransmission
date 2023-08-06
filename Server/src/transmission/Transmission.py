import socket
from typing import Tuple

from src.ED.utils import *


class transmission:
    def __init__(self) -> None:
        pass

    def send_file(self, secret: bytes, pack_type: int, data: bytes, conn: socket.socket):
        """Send file to client

        Args:
            pack_type (int): [description]
            data (bytes): [description]
            conn (socket.socket): [description]
        """
        ########
        # Establish packet
        ########
        packet = bytearray()

        # Add DDos key
        for letter in get_sha256(secret):
            packet.append(ord(letter))

        # Add packet type
        packet.append(pack_type)

        # Add data
        packet.extend(data)

        ########
        # Send packet
        ########
        conn.sendall(packet)

    def unpack(data: bytes) -> tuple:
        """Verify and unpack data

        Args:
            data (bytes): data received from client

        Returns:
            tuple: tuple of packet type, data
            None: otherwise
        """
        ########
        # Verify: min len of packet, valid DDoS string
        ########
        if len(data) < 65:
            return None

        hash_val = data[0:64].decode()
        pack_type = int.from_bytes(data[64:65], "little")
        data = data[65:].decode()
