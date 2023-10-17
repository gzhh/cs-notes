class Dog:
    """一次模拟小狗的简单尝试。"""

    def __init__(self, name, age):
        """初始化属性name和age。"""
        self.name = name
        self.age = age

    def sit(self):
        """模拟小狗收到命令时蹲下。"""
        print(f"{self.name} is now sitting.")

    def roll_over(self):
        """模拟小狗收到命令时打滚。"""
        print(f"{self.name} rolled over!")

"""
类编写风格：
1. 类名应采用 ，即将类名中的每个单词的首字母都大写，而不使用下划 线。实例名和模块名都采用小写格式，并在单词之间加上下划线。
2. 对于每个类，都应紧跟在类定义后面包含一个文档字符串。这种文档字符串简要地描述类的功能，并遵循编写函数的文档字符串时采用的格式约定。
    每个模块也都应包含一个文档字符串，对其中的类可用于做什么进行描述。
3. 可使用空行来组织代码，但不要滥用。在类中，可使用一个空行来分隔方法;而在模块中，可使用两个空行来分隔类。
4. 需要同时导入标准库中的模块和你编写的模块时，先编写导入标准库模块的import 语句，再添加一个空行，然后编写导入你自己编写的模块的import 语句。在包含 多条import 语句的程序中，这种做法让人更容易明白程序使用的各个模块都来自 何处。
"""

"""
__init__() 方法
1. 是一个特殊方 法，每当你根据类创建新实例时，Python都会自动运行它。
2. 形参self 必不可少，而且必须位于其他形参的前面。它是一个指 向实例本身的引用，让实例能够访问类中的属性和方法。
3. 通过实参向传递参数时， self 会自动传递，因此不需要传递它。

类和实例
"""

my_dog = Dog('Willie', 6)
print(f"My dog's name is {my_dog.name}.")
print(f"My dog is {my_dog.age} years old.")
my_dog.sit()
my_dog.roll_over()


"""
继承：一个类继承另一个类时，将自动获得另一个类的所有属性和方法。原有的类称为父类，而新类称为子类。
    子类继承了父类的所有属性和方法，同时 还可以定义自己的属性和方法。

1. 定义子类时，必须在圆括号内指定父类的名称。
2. 子类的方法 __init__
    在既有类的基础上编写新类时，通常要调用父类的方法__init__() 。
    这将初始化 在父类__init__() 方法中定义的所有属性，从而让子类包含这些属性。
3. super() 是一个特殊函数，让你能够调用父类的方法。
4. 子类可以添加新的方法
5. 子类可以重写父类的方法
6. 可以将实例作为类的属性
"""
class Car:
    """一次模拟汽车的简单尝试。"""

    def __init__(self, make, model, year):
        self.make = make
        self.model = model
        self.year = year
        self.odometer_reading = 0

    def get_descriptive_name(self):
        long_name = f"{self.year} {self.make} {self.model}"
        return long_name.title()

    def read_odometer(self):
        print(f"This car has {self.odometer_reading} miles on it.")

    def update_odometer(self, mileage):
        if mileage >= self.odometer_reading:
            self.odometer_reading = mileage
        else:
            print("You can't roll back an odometer!")

    def increment_odometer(self, miles):
        self.odometer_reading += miles

class ElectricCar(Car):
    """电动汽车的独特之处。"""

    def __init__(self, make, model, year):
        """初始化父类的属性。"""
        super().__init__(make, model, year)
    
    def describe_battery(self):
        """打印一条描述电瓶容量的消息。"""
        print(f"This car has a {self.battery_size}-kWh battery.")

my_tesla = ElectricCar('tesla', 'model s', 2019)
print(my_tesla.get_descriptive_name())