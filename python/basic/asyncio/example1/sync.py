import asyncio
import time

async def waiter(secs):
    print("One", secs)
    time.sleep(secs)
    print("Two", secs)

async def main():
    print("main start")
    await waiter(1)
    await waiter(2)
    await waiter(3)
    print("main end")

asyncio.run(main())