import asyncio
import time

async def say_after(delay, what):
    print(delay)
    await asyncio.sleep(delay)
    print(what)

async def main():
    print(f"started at {time.strftime('%X')}")

    await asyncio.gather(say_after(1, 'hello'), say_after(2, 'world'))

    print(f"finished at {time.strftime('%X')}")

asyncio.run(main())
