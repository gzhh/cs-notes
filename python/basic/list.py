bicycles = ['trek', 'cannondale', 'redline', 'specialized']
print(bicycles)
print(len(bicycles))

print(bicycles[0]) # 第一个元素
print(bicycles[-1]) # 最后一个元素，仅当列表为空时，这种访问最后一个元素的方式才会导致错误

# 添加元素
bicycles.append("gaint")
print(bicycles)

bicycles.insert(0, "merida")
print(bicycles)

# 删除元素
popedBicycle = bicycles.pop()
print(popedBicycle)
popedBicycle = bicycles.pop(2)
print(popedBicycle)
print(bicycles)

bicycles.remove("trek") # 方法remove() 只删除第一个指定的值。如果要删除的值可能在列表中 出现多次，就需要使用循环来确保将每个值都删除。
print(bicycles)

del(bicycles[0])
print(bicycles)


# 组织列表
cars = ['bmw', 'audi', 'toyota', 'subaru']
cars.sort()
print(cars)
cars.sort(reverse=True)
print(cars)
print("\nHere is the sorted list:")
print(sorted(cars)) # 临时排序

cars.reverse()
print(cars)


# 操作列表
magicians = ['alice', 'david', 'carolina']
for magician in magicians:
    print(magician)

magicians = ['alice', 'david', 'carolina']
for magician in magicians:
    print(f"{magician.title()}, that was a great trick!")
# 逻辑错误。但仍可以读到 magician 变量
print(f"I can't wait to see your next trick, {magician.title()}.\n")

# 数字列表
for value in range(1, 5): # range 左闭右开
    print(value)
numbers = list(range(1, 6))
print(numbers)
even_numbers = list(range(2, 11, 2)) # 步长
print(even_numbers)
digits = [1, 2, 3, 4, 5, 6, 7, 8, 9, 0]
min(digits)
max(digits)
sum(digits)

# 列表解析：可以精简代码为一行
squares = [value**2 for value in range(1, 11)]
print(squares)


# 切片：列表的部分元素
players = ['charles', 'martina', 'michael', 'florence', 'eli']
print(players[0:3]) # 左闭右开
print(players[:4]) # 左闭右开
print(players[2:]) # 左闭右开
print(players[-3:]) # 左闭右开
print(players[1:5:2]) # 左闭右开
## 遍历
for player in players[:3]:
    print(player.title())
## 复制列表（深拷贝）
my_foods = ['pizza', 'falafel', 'carrot cake']
friend_foods = my_foods[:]
## 复制列表（浅拷贝）
friend_foods = my_foods


# 元组 tuple：一系列不可修改的元素
dimensions = (200, 50)
print(dimensions[0])
print(dimensions[1])
## 一个元素的元组
my_t = (3,)
## 遍历
dimensions = (200, 50)
for dimension in dimensions:
    print(dimension)
## 元组可以重新赋值
dimensions = (200, 50)
dimensions = (400, 100)