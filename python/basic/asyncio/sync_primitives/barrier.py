"""
A barrier is a simple synchronization primitive that allows to block until parties number of tasks are waiting on it.
Tasks can wait on the wait() method and would be blocked until the specified number of tasks end up waiting on wait().
At that point all of the waiting tasks would unblock simultaneously.
"""

import asyncio

async def example_barrier():
   # barrier with 3 parties
   b = asyncio.Barrier(3)

   # create 2 new waiting tasks
   asyncio.create_task(b.wait())
   asyncio.create_task(b.wait())

   await asyncio.sleep(0)
   print(b)

   # The third .wait() call passes the barrier
   await b.wait()
   print(b)
   print("barrier passed")

   await asyncio.sleep(0)
   print(b)

asyncio.run(example_barrier())
