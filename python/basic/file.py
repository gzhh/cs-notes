"""
文件：

关键字with 在不再需要访问文件后将其关闭。在这个程序中，注意到我们调用了 open() ，比显示的调用close()更好 。
"""

# 读取整个文件
with open('file.txt') as file_object:
    contents = file_object.read()
print(contents)


# 逐行读取
with open("file.txt") as file_object:
    for line in file_object:
        print(line)

# 读取文件行列表
with open("file.txt") as file_object:
    lines = file_object.readlines()

pi_string = ''
for line in lines:
    pi_string += line.strip()

print(pi_string)
print(len(pi_string))


# million pi demo 
with open("pi_million_digits.txt") as file_object:
    lines = file_object.readlines()

for line in lines:
    pi_string += line.strip()

birthday = input("Enter your birthday, in the form mmddyy: ")
if birthday in pi_string:
    print("Your birthday appears in the first million digits of pi!")
else:
    print("Your birthday does not appear in the first million digits of pi.")


# 写入文件
"""
写模式：
r 读
w 写
a 附加
r+ 读写

默认为只读，如果要写入的文件不存在，函数open() 将自动创建它；
Python只能将字符串写入文本文件。要将数值数据存储到文本文件中， 必须先使用函数str() 将其转换为字符串格式。
"""
filename = 'programming.txt'
with open(filename, 'w') as file_object:
    file_object.write("I love programming.")