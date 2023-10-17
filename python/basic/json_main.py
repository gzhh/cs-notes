import json

"""
JSON 数据操作
"""

# json.dump()
numbers = [2, 3, 5, 7, 11, 13]
filename = 'json_numbers.json'
with open(filename, 'w') as f:
    json.dump(numbers, f)

# json.load()
filename = 'json_numbers.json'
with open(filename) as f:
    numbers = json.load(f)
print(numbers)

