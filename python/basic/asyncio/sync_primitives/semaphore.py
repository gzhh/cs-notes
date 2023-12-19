"""
sem = asyncio.Semaphore(10)

# ... later
async with sem:
    # work with shared resource


which is equivalent to:


sem = asyncio.Semaphore(10)

# ... later
await sem.acquire()
try:
    # work with shared resource
finally:
    sem.release()
"""
