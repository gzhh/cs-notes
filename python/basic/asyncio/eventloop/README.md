### Event Loop Methods

**lower-level**

The event loop is the core of every asyncio application. Event loops run asynchronous tasks and callbacks, perform network IO operations, and run subprocesses.

Obtaining the Event Loop
- asyncio.get_running_loop()
  - Return the running event loop in the current OS thread.
- asyncio.get_event_loop()
  - Get the current event loop.
- asyncio.set_event_loop(loop)
  - Set loop as the current event loop for the current OS thread.
- asyncio.new_event_loop()
  - Create and return a new event loop object.


Running and stopping the loop
- loop.run_until_complete(future)
- loop.run_forever()
- loop.stop()
- loop.is_running()
- loop.is_close()
- loop.close()

Scheduling callbacks
- loop.call_soon


Scheduling delayed callbacks
- loop.call_later(delay, callback, *args, context=None)
- loop.call_at(when, callback, *args, context=None)
- loop.time()

Creating Futures and Tasks
- loop.create_future()
- loop.create_task(coro, *, name=None, context=None)
- loop.set_task_factory(factory)
- loop.get_task_factory()

Opening network connections
- loop.create_connection(protocol_factory, host=None, port=None, *,
- loop.create_datagram_endpoint(protocol_factory, local_addr=None, remote_addr=None, *,
- loop.create_unix_connection(protocol_factory, path=None, *,

Creating network servers
- loop.create_server(protocol_factory, host=None, port=None, *,
- loop.create_unix_server(protocol_factory, path=None, *,
- loop.connect_accepted_socket(protocol_factory, sock, *,

Transferring files
- loop.sendfile(transport, file, offset=0, count=None, *, fallback=True)

TLS Upgrade
- loop.start_tls(transport, protocol, sslcontext, *,

DNS
- loop.getaddrinfo(host, port, *, family=0, type=0, proto=0, flags=0)
- loop.getnameinfo(sockaddr, flags=0)

Executing code in thread or process pools
- awaitable loop.run_in_executor(executor, func, *args)

Enabling debug mode
- loop.get_debug()
- loop.set_debug(enabled: bool)
