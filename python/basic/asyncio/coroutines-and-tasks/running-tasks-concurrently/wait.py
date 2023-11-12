"""
asyncio.gather vs asyncio.wait
- https://stackoverflow.com/questions/42231161/asyncio-gather-vs-asyncio-wait-vs-asyncio-taskgroup
- https://www.dongwm.com/post/understand-asyncio-2/

asyncio.gather()
- high level

asyncio.wait()
- lower leve
"""

import asyncio

async def a():
    print('Suspending a')
    await asyncio.sleep(3)
    print('Resuming a')
    return 'A'


async def b():
    print('Suspending b')
    await asyncio.sleep(1)
    print('Resuming b')
    return 'B'

async def main():
    task_a = asyncio.create_task(a())
    task_b = asyncio.create_task(b())
    done, pending = await asyncio.wait([task_a, task_b])
    print('---done ', done)
    print('---pending ', pending)

asyncio.run(main())
