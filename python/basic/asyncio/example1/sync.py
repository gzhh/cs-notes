import asyncio

async def waiter(secs):
    print("One", secs)
    await asyncio.sleep(secs)
    print("Two", secs)

async def main():
    print("main start")

    # sync
    await waiter(1)
    await waiter(2)
    await waiter(3)

    # async
    # await asyncio.gather(waiter(1), waiter(2), waiter(3))
    print("main end")

asyncio.run(main())
