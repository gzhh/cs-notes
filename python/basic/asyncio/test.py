import asyncio

# python3.11
async def hello():
    print("Hello world!")
    await asyncio.sleep(1)
    print("Hello again!")

asyncio.run(hello())