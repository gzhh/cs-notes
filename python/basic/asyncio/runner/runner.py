"""
Runner context manager
A context manager that simplifies multiple async function calls in the same context.
"""

import asyncio

async def main():
    await asyncio.sleep(1)
    print('hello')

with asyncio.Runner() as runner:
    runner.run(main())