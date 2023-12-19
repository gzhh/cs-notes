import asyncio

async def long_running_task():
    print("long running task start")
    await asyncio.sleep(15)
    print("long running task end")

async def main():
    loop = asyncio.get_running_loop()
    deadline = loop.time() + 10
    try:
        async with asyncio.timeout_at(deadline):
            await long_running_task()
    except TimeoutError:
        print("The long operation timed out, but we've handled it.")

    print("This statement will run regardless.")

asyncio.run(main())
