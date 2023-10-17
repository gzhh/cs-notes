"""
函数
import module_name #                                          1.导入整个模块
from module_name import function_0, function_1, function_2 #  2.导入特定的函数
from module_name import function_name as fn #                 3.使用 as 给函数指定别名
import module_name as mn                                      4.使用 as 给模块指定别名
from module_name import * # 4.使                              5.导入模块中的所有函数

类
1. 导入单个类
    from module_name import Class
2. 在一个模块中存储多个类
3. 从一个模块中导入多个类
    from module_name import Class1, Class2, Class3
4. 导入整个模块（推荐）
    import module_name
5. 导入模块中所有的类（不推荐）
    from module_name import *
6. 在一个模块中导入另一个模块
7. 使用别名 as
8. 自定义工作流程
    一开始应让代码结构尽可能简单。先尽可能在一个文件中完成所有的工作，确定一
    切都能正确运行后，再将类移到独立的模块中。如果你喜欢模块和文件的交互方
    式，可在项目开始时就尝试将类存储到模块中。先找出让你能够编写出可行代码的
    方式，再尝试改进代码。
"""


"""
Python 标准库
是一组 Python 安装后自带的系统模块
>>> from random import randint
>>> randint(1, 6)
3 

Python 外部模块
还可以从其他地方下载外部模块。
"""