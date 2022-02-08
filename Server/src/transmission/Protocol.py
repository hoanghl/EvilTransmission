from src import *


class Protocol_3Vil:
    ## PROCESS INCOMING REQUEST ############################

    def __init__(self) -> None:
        self.component = 0  ## mark which part of packet has been processed
        self.payload_size = None  ## this will
        self.flags = None

    def decompose_flag(self, byte_flag: bytes):
        byte_flag = int.from_bytes(byte_flag, "big")

        is_sucessive = None
        packet_type = None
        response_type = None

        ## Determine whether there is sucessive packets
        packet_type = (byte_flag & packet_type_filter) >> Shifts[PacketType]

        ## Determine packet type
        is_sucessive = byte_flag & sucessive_filter == 1

        ## If response, determine response type
        if packet_type == Response.code:
            response_type = (byte_flag & response_type_filter) >> Shifts[ResponseType]

        return {
            "is_sucessive": is_sucessive,
            "packet_type": packet_type,
            "response_type": response_type,
        }

    def decompose_size(self, byte_size: bytes):
        byte_size = int.from_bytes(byte_size, "big")

        return byte_size

    def decompose_buff(self, buffer: Buffer) -> list:
        """Decompose buffer into packets

        Args:
            buffer (Buffer): received buffer

        Returns:
            list: list of tuple of flags and payload of packets
        """
        recv_packets = []

        while True:
            if self.component == 0:
                if len(buffer) < PacketType.LEN_FLAG:
                    break

                ## Step 1: Read byte Flag
                self.component += 1

                byte_flag = buffer[: PacketType.LEN_FLAG]
                self.flags = self.decompose_flag(byte_flag)

            elif self.component == 1:
                if len(buffer) < PacketType.LEN_SIZE:
                    break

                ## Step 2: Read bytes Size
                self.component += 1

                byte_size = buffer[: PacketType.LEN_SIZE]
                self.payload_size = self.decompose_size(byte_size)
            elif self.component == 2:
                if len(buffer) < self.payload_size:
                    break

                ## Step 3: Read payload
                payload = buffer[: self.payload_size]

                self.component = 0
                recv_packets.append((self.flags, payload))

                if self.flags["is_sucessive"] is False:
                    break
            else:
                break

        return recv_packets

    ## PROCESS OUTGOING REQUEST ############################

    def create_packet(
        self, packet_type, payload: bytes = None, permitted=True, internal_error=False
    ):
        """Create packet (request or response)

        Args:
            packet_type (PacketType): Type of packet, a subclass of PacketType
            payload (bytes, optional): payload to be sent in case Hash/payload request. Defaults to None.
            permitted (bool, optional): Flag for response. Defaults to True.
            internal_error (bool, optional): Flag for response. Defaults to False.

        Returns:
            bytes: buffer containing bytes to be sent
        """

        buffer = bytes()

        if packet_type in [Hash, Data]:
            assert payload is not None

            accum = 0
            while accum < len(payload):
                ## Craft byte_flag and add to buffer

                # Craft packet_type
                byte_flag = packet_type.code << Shifts[packet_type]

                # Craft sucessive
                sucessive = int(len(payload) - accum > MAX_SIZE_PACKET)
                byte_flag += sucessive

                byte_flag = byte_flag.to_bytes(1, "big")
                buffer += byte_flag

                ## Add byte_size to buffer
                data_size = min(len(payload) - accum, MAX_SIZE_PACKET)
                byte_size = data_size.to_bytes(2, "big")
                buffer += byte_size

                ## Add payload chunk to buffer
                buffer += payload[accum : accum + data_size]

                accum += data_size

        elif packet_type is Response:
            if internal_error is True:
                response_type = InternalErr
            elif permitted is True:
                response_type = Permitted
            elif permitted is False:
                response_type = NotPermitted

            byte_flag = (response_type.code << Shifts[response_type]) + (
                packet_type.code << Shifts[packet_type]
            )
            byte_flag = byte_flag.to_bytes(1, "big")
            buffer += byte_flag

            byte_size = int(0).to_bytes(2, "big")
            buffer += byte_size

        return buffer
