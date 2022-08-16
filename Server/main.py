from src import *
from src.storage import Storage
from src.transmission import Server


async def start_server(storage):
    logger.info("Initialize: Server")
    logger.info(f"== Server operates at: {HOST} - {PORT}")
    loop = asyncio.get_running_loop()

    server = await loop.create_server(lambda: Server(storage), HOST, PORT)

    async with server:
        await server.serve_forever()


async def main():
    storage = Storage()

    await asyncio.gather(
        storage.back_up(),
        start_server(storage),
    )


if __name__ == "__main__":
    asyncio.run(main())
