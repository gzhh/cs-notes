import unittest

"""
Python 单元测试：
对于大型项目，要进行全覆盖测试可能很难。通常，最初只要针对代 码的重要行为编写测试即可，等项目被广泛使用时再考虑全覆盖。

1. 先导入模块unittest 和 要测试的函数，再创建一个继承unittest.TestCase 的类
2. 并编写一系列方法对 函数行为的不同方面进行测试。
3. 断言方法核实得到的 结果是否与期望的结果一致。

测试未通过时，不要修改测试，而应修复导致测试不能通过的代码:检查刚刚对函数所做的修改，找出导致函数行为不符合预期的修改。
"""

def get_formatted_name(first, last, middle=''):
    """生成整洁的姓名。"""
    if middle:
        full_name = f"{first} {middle} {last}"
    else:
        full_name = f"{first} {last}"
    return full_name.title()

class NamesTestCase(unittest.TestCase):
    """测试name_function.py。"""

    def test_first_last_name(self):
        """能够正确地处理像Janis Joplin这样的姓名吗?"""
        # 能
        formatted_name = get_formatted_name('janis', 'joplin')
        self.assertEqual(formatted_name, 'Janis Joplin')

        # 不能
        # formatted_name = get_formatted_name('janis', 'joplin')
        # self.assertEqual(formatted_name, 'JanisJoplin')
    

    def test_first_last_middle_name(self):
        """能够正确地处理像Wolfgang Amadeus Mozart这样的姓名吗?"""
        formatted_name = get_formatted_name('wolfgang', 'mozart', 'amadeus')
        self.assertEqual(formatted_name, 'Wolfgang Amadeus Mozart')

# 如果这个文件作为主程序执行，变 量__name__ 将被设置为'__main__' 。
if __name__ == '__main__':
    unittest.main()