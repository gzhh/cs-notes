str1 = 'Hello, world'
str2 = "Hello, world"

# 字符串操作 
message = " Hello gzhh "
print(message.title())
print(message.upper())
print(message.lower())
# 删除空白字符（空格符、换行符、制表符）
print(message.strip())
print(message.lstrip())
print(message.rstrip())
# 字符串切分
print(message.split())

# 字符串格式化 f
first_name = "ada"
last_name = "lovelace"
full_name = f"{first_name} {last_name}"
print(full_name)
message = f"Hello, {full_name.title()}!"
print(message)


# 数字-科学计数法，下划线
universe_age = 14_000_000_000
print(universe_age)
universe_age = 14000_000000
print(universe_age)


# 多变量赋值
x, y, z = 0, 1, 2


# 常量类似于变量，但其值在程序的整个生命周期内保持不变。
# Python没有内置的常量类型，但Python程序员会使用全大写来指出应将某个变量视为常量，其值应始终不变
MAX_CONNECTIONS = 1000