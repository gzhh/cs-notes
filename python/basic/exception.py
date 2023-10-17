"""
异常：
Python使用称为的特殊对象来管理程序执行期间发生的错误。每当发生让 Python不知所措的错误时，它都会创建一个异常对象。
如果你编写了处理该异常的 代码，程序将继续运行;如果未对异常进行处理，程序将停止并显示traceback，其中包含有关异常的报告。


"""

# 使用 try catch 处理异常
try:
    print(5/0)
except ZeroDivisionError:
    print("You can't divide by zero!")
except Exception:
    print("Got exception")


# 使用异常避免崩溃
"""
print("Give me two numbers, and I'll divide them.")
print("Enter 'q' to quit.")
while True:
    first_number = input("\nFirst number: ")
    if first_number == 'q':
        break
    second_number = input("Second number: ")
    if second_number == 'q':
        break

    try:
        answer = int(first_number) / int(second_number)
    except ZeroDivisionError:
        print("You can't divide by 0!")
    else:
        print(answer)
"""


# 处理文件 not found 异常
filename = 'alice.txt'
try:
    with open(filename, encoding='utf-8') as f:
        contents = f.read()
except FileNotFoundError:
    print(f"Sorry, the file {filename} does not exist.")


# demo: count words
def count_words(filename):
    """计算一个文件大致包含多少个单词。"""
    try:
        with open(filename, encoding='utf-8') as f:
            contents = f.read()
    except FileNotFoundError:
        print(f"Sorry, the file {filename} does not exist.")
    else:
        words = contents.split()
        num_words = len(words)
        print(f"The file {filename} has about {num_words} words.")

filename = 'programming.txt'
count_words(filename)


# 静默失败
"""
要让程序静默失败，可像通常那样编写try 代码块，但在except 代码块 中明确地告诉Python什么都不要做。
Python有一个pass 语句，可用于让Python在 代码块中什么都不要做:

"""
def count_words(filename):
    """计算一个文件大致包含多少个单词。"""
    try:
        with open("not_found.txt") as f:
            f.read()
    except FileNotFoundError:
        pass # 不做任何处理
    else:
        print("else branch")
filenames = ['alice.txt', 'siddhartha.txt', 'moby_dick.txt', 'little_women.txt']
for filename in filenames:
    count_words(filename)