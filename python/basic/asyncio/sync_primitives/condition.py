"""
cond = asyncio.Condition()

# ... later
async with cond:
    await cond.wait()


which is equivalent to:


cond = asyncio.Condition()

# ... later
await cond.acquire()
try:
    await cond.wait()
finally:
    cond.release()
"""
