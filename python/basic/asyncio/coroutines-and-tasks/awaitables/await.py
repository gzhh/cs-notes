import asyncio

async def nested():
    return 42

async def main():
    # RuntimeWarning: coroutine 'nested' was never awaited
    # nested()

    print(await nested())

asyncio.run(main())
