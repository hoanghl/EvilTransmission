from src import *

from .Protocol import Protocol_3Vil


class Server(asyncio.Protocol):
    def __init__(self, storage) -> None:
        super().__init__()

        self.storage = storage
        self.protocol_processor = Protocol_3Vil()

        self.buffer = None
        self.transport = None
        self.payload = None
        # self.total_received = 0
        # self.total_packets = 0

    def connection_made(self, transport):
        logger.info(f"Connected by: {transport.get_extra_info('peername')}")

        self.transport = transport

    def connection_lost(self, exc: Exception) -> None:

        self.transport = None
        self.buffer = None
        self.payload = None

    def response(self, buffer):
        self.transport.write(buffer)

    def close(self):
        logger.info(f"Close connection from: {self.transport.get_extra_info('peername')}")
        self.transport.close()

    def data_received(self, data):

        if self.buffer is None:
            self.buffer = Buffer()

        self.buffer += data
        # self.total_received += len(data)
        # print(f"total received: {self.total_received}")

        self.digest_buffer()

    def digest_buffer(self):
        try:
            recv_packets = self.protocol_processor.decompose_buff(self.buffer)
            if len(recv_packets) == 0:
                return

            for packet in recv_packets:
                flags, payload = packet

                print(f"hash payload: {get_sha256(payload)}")

                if flags["packet_type"] == Hash.code:

                    logger.info(f"Hash code: {payload.hex()}")

                    ## Check hash
                    if self.storage.check_hash(payload):
                        resp_buffer = self.protocol_processor.create_packet(
                            Response, permitted=False
                        )
                    else:
                        resp_buffer = self.protocol_processor.create_packet(
                            Response, permitted=True
                        )

                    self.response(resp_buffer)

                elif flags["packet_type"] == Data.code:
                    if self.payload is None:
                        self.payload = []
                    self.payload += [payload]

                else:
                    logger.error(f"flags['packet_type'] = {flags['packet_type']}")
                    raise NotImplementedError()

                if flags["is_sucessive"] is False:
                    self.close()

                    if flags["packet_type"] == Data.code:
                        self.storage.store_file(b"".join(self.payload))

                # self.total_packets += 1
                # print(f"total packets: {self.total_packets}")

        except Exception:
            self.response(self.protocol_processor.create_packet(Response, internal_error=True))
            traceback.print_exc()
