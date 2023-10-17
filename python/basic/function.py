"""
实参和形参
函数的定义中，参数叫形参 parameters
调用函数的时候传递的参数值叫实参 arguments

实参传递
1. 位置实参，即按顺序传递参数
2. 关键词实参
describe_pet(animal_type='hamster', pet_name='harry')
3. 默认值
def describe_pet(pet_name, animal_type='dog'):
4. 三种方式混合使用

返回值

"""

# 传递列表
## 在函数中修改列表：在函数中对这个列表所做的任何修改都是永久性的，这让你能够高效地处理大量数据。
# 首先创建一个列表，其中包含一些要打印的设计。
unprinted_designs = ['phone case', 'robot pendant', 'dodecahedron']
completed_models = []
# 模拟打印每个设计，直到没有未打印的设计为止。
# 打印每个设计后，都将其移到列表completed_models中。
while unprinted_designs:
    current_design = unprinted_designs.pop()
    print(f"Printing model: {current_design}")
    completed_models.append(current_design)
# 显示打印好的所有模型。
print("\nThe following models have been printed:")
for completed_model in completed_models:
    print(completed_model)
print(unprinted_designs)

## 禁止修改列表：将列表的副本传递给函数，切片表示法[:] 创建列表的副本
## 虽然向函数传递列表的副本可保留原始列表的内容，但除非有充分的理由，否则还是应该将原始列表传递给函数。
## 这是因为让函数使用现成的列表可避免花时间和内存创建副本，从而提高效率，在处理大型列表时尤其如此。
# function_name(list_name_[:])


# 传递任意数量的实参
## 形参名*toppings 中的星号让Python创建一个名为toppings 的空元组，并将收 到的所有值都封装到这个元组中。
def make_pizza(*toppings):
    """打印顾客点的所有配料。"""
    print(toppings)
make_pizza('pepperoni')
make_pizza('mushrooms', 'green peppers', 'extra cheese')
## 如果函数接受不同类型的实参，必须在函数定义中将接纳任意数量实参的形参放在最后
## def build_profile(first, last, **user_info):


# 将函数存储在模块中
"""
使用函数的优点之一是可将代码块与主程序分离。通过给函数指定描述性名称，可 让主程序容易理解得多。你还可以更进一步，将函数存储在称为 的独立文件 中，再将模块 到主程序中。import 语句允许在当前运行的程序文件中使用模 块中的代码。

通过将函数存储在独立的文件中，可隐藏程序代码的细节，将重点放在程序的高层
逻辑上。这还能让你在众多不同的程序中重用函数。将函数存储在独立文件中后，
可与其他程序员共享这些文件而不是整个程序。知道如何导入函数还能让你使用其
他程序员编写的函数库。
"""


# 函数编写规范
"""
命名：应给函数指定描述性名称，且只在其中使用小写字母和下划线。给模块命名时同理。
注释：注释应紧跟在函数定义后面，并采用文档字符串格式。
参数：给形参指定默认值时，等号两边不要有空格；对于函数调用中的关键字实参，也应遵循这种约定；
"""