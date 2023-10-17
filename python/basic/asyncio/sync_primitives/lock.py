import asyncio

lock = asyncio.Lock()

async def main():
    # ... later
    await lock.acquire()
    try:
        # access shared state
        pass
    finally:
        lock.release()

asyncio.run(main())