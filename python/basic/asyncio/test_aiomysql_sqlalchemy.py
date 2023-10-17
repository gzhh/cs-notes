import asyncio
import sqlalchemy as sa

from aiomysql.sa import create_engine


"""
使用 sqlalchemy 拼接 sql 语句，然后使用 aiomysql execute 执行语句。
不使用 sqlalchemy 的 orm 功能。 
"""

metadata = sa.MetaData()

tbl = sa.Table('test', metadata,
               sa.Column('id', sa.Integer, primary_key=True),
               sa.Column('val', sa.String(255)))


async def query(self, statement: str) -> list:
    await self.init_engine()
    async with self._engine.acquire() as conn:
        result = await conn.execute(statement)
        ret = await result.fetchall()
        return ret


async def go(loop):
    engine = await create_engine(user='root', db='test',
                                 host='127.0.0.1', password='', loop=loop, autocommit=True)
    async with engine.acquire() as conn:
        # insert
        # await conn.execute(tbl.insert().values(val='abc'))
        # await conn.execute(tbl.insert().values(val='xyz'))

        # select demo1
        # async for row in await conn.execute(tbl.select()):
        #     print(row.id, row.val)

        # select demo2
        # res = await conn.execute(tbl.select())
        # async for row in res:
        #     print(row.id, row.val)

        # select fetchall(list type)
        stmt = 'select * from test'
        print("---stmt---\n", stmt)
        result = await conn.execute(stmt)
        rows = await result.fetchall()
        print(rows)
        for row in rows:
            print(type(row))
            print(row)
            print(row[0], row[1])
            print(row.id, row.val)

        # stmt = tbl.select()
        # print("---stmt---\n", stmt)

        # stmt = tbl.select().where(tbl.c.id == 1)
        # print("---stmt---\n", stmt)

        # from sqlalchemy import select
        # stmt = select([tbl]).where(tbl.c.val == 'val1')
        # print("---stmt---\n", stmt)

        # # specific fields and tables
        # stmt = select([tbl.c.id]).select_from(tbl) \
        #     .where(tbl.c.val == 'val1')
        # print("---stmt---\n", stmt)

        # stmt = select([tbl.c.id]).select_from(tbl \
        #         .join(tbl, tbl.c.id == tbl.c.id)
        #     ) \
        #     .where(tbl.c.val == 'val1')
        # print("---stmt---\n", stmt)


    engine.close()
    await engine.wait_closed()


loop = asyncio.get_event_loop()
loop.run_until_complete(go(loop))
