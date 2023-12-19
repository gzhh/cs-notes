import asyncio
import json
import time
import threading

async def task():
    print('before %s' % threading.current_thread())
    await asyncio.sleep(1)
    print('after %s' % threading.current_thread())

async def task1():
    while True:
        json_str = '{"id": 1, "name": "task1", "tags": ["Sport", "Music", "Read"]}'
        obj = json.loads(json_str)
        time.sleep(1)
        print(obj)


async def task2():
    json_str = '{"id": 2, "name": "task2", "tags": ["Sport", "Music", "Read"]}'
    obj = json.loads(json_str)
    print(obj)


def run():
    asyncio.run(task())

    # async def test():
    #     await asyncio.gather(task(), task())
    # asyncio.run(test())

    # async def test():
    #     tasks = [
    #         asyncio.create_task(task()),
    #         asyncio.create_task(task()),
    #     ]
    #     await asyncio.wait(tasks)
    # asyncio.run(test())

def run_forever():
    loop = asyncio.get_event_loop()

    try:
        loop.create_task(task1())
        loop.run_forever()
    finally:
        loop.close()

def run_until_complete():
    loop = asyncio.get_event_loop()
    interval = 1

    while True:
        loop.run_until_complete(task2())
        time.sleep(interval) 


"""
asyncio.run()
loop.run_forever()
loop.run_until_complete()
"""

if __name__ == "__main__":
    run()
    # run_forever()
    # run_until_complete()

