# 字典：一系列键值对
alien_0 = {}
alien_0['color'] = 'green'
alien_0['points'] = 5
print(alien_0)

alien_0 = {'color': 'green', 'points': 5}
print(alien_0['color'])
print(alien_0['points'])

## 增
alien_0['x_position'] = 0
alien_0['y_position'] = 25
print(alien_0)
## 改
alien_0['x_position'] = 10
## 删
del alien_0['points']
print(alien_0)

## 多行赋值
favorite_languages = {
    'jen': 'python',
    'sarah': 'c',
    'edward': 'ruby',
    'phil': 'python',
    }

## 元素访问
alien_0 = {'color': 'green', 'speed': 'slow'}
# print(alien_0['points']) # error, key not found
# 调用get() 时，如果没有指定第二个参数且指定的键不存在，Python 将返回值None
point_value = alien_0.get('points', 'No point value assigned.')
print(point_value)

## 遍历所有key-value
user_0 = {
      'username': 'efermi',
      'first': 'enrico',
      'last': 'fermi',
      }
for key, value in user_0.items():
    print(f"\nKey: {key}")
    print(f"Value: {value}")
## 遍历所有键
for name in favorite_languages.keys():
    print(name.title())
## 遍历所有值
for language in favorite_languages.values():
    print(language.title())
## 按排序遍历
for name in sorted(favorite_languages.keys()):
    print(f"{name.title()}, thank you for taking the poll.")


# 字典列表

# 字典嵌套
## 在字典中存储列表
pizza = {
    'crust': 'thick',
    'toppings': ['mushrooms', 'extra cheese'],
    }
## 在字典中存储字典
users = {
      'aeinstein': {
          'first': 'albert',
          'last': 'einstein',
          'location': 'princeton',
          },
      'mcurie': {
          'first': 'marie',
          'last': 'curie',
          'location': 'paris',
          },
}