import asyncio
import aiomysql

loop = asyncio.get_event_loop()


# conn pool
async def test_example_pool(loop):
    pool = await aiomysql.create_pool(host='127.0.0.1', port=3306,
                                      user='root', password='',
                                      db='test', loop=loop)
    async with pool.acquire() as conn:
        async with conn.cursor() as cur:
            await cur.execute("SELECT 42;")
            print(cur.description)
            (r,) = await cur.fetchone()
            assert r == 42
    pool.close()
    await pool.wait_closed()

# loop.run_until_complete(test_example_pool(loop))


# query (autocommit=False)
async def test_example(loop):
    conn = await aiomysql.connect(host='127.0.0.1', port=3306,
                                       user='root', password='', db='test',
                                       loop=loop)

    cur = await conn.cursor()
    await cur.execute("SELECT id,value FROM test")
    await conn.commit()

    print(cur.description)
    r = await cur.fetchall()
    print(r)
    await cur.close()
    conn.close()

loop.run_until_complete(test_example(loop))


# create and insert (autocommit=True)
async def test_example_execute(loop):
    conn = await aiomysql.connect(host='127.0.0.1', port=3306,
                                       user='root', password='', db='test',
                                       loop=loop, autocommit=True)

    cur = await conn.cursor()
    async with conn.cursor() as cur:
        await cur.execute("DROP TABLE IF EXISTS music_style;")
        await cur.execute("""CREATE TABLE music_style
                                  (id INT,
                                  name VARCHAR(255),
                                  PRIMARY KEY (id));""")

        # insert 3 rows one by one
        await cur.execute("INSERT INTO music_style VALUES(1,'heavy metal')")
        await cur.execute("INSERT INTO music_style VALUES(2,'death metal');")
        await cur.execute("INSERT INTO music_style VALUES(3,'power metal');")

    conn.close()

# loop.run_until_complete(test_example_execute(loop))
