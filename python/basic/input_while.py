message = input("Please enter your name:\n")
print(message)

# python 无法将字符串和整数进行比较，需要讲字符串数字转成整数 int(str) 再比较
age = input("How old are you? ")
age = int(age)
print(age >= 18)

# while 循环
"""
关键词
break
continue
"""

unconfirmed_users = ['alice', 'brian', 'candace']
confirmed_users = []
while unconfirmed_users: # 不断运行，直到列表unconfirmed_users变成空的
    current_user = unconfirmed_users.pop()
    print(f"Verifying user: {current_user.title()}")
    confirmed_users.append(current_user)


# while 循环 input 读数据
responses = {}
# 设置一个标志，指出调查是否继续。 polling_active = True
while polling_active:
# 提示输入被调查者的名字和回答。
    name = input("\nWhat is your name? ")
    response = input("Which mountain would you like to climb someday? ")
    # 将回答存储在字典中。
    responses[name] = response
    # 看看是否还有人要参与调查。
    repeat = input("Would you like to let another person respond? (yes/ no) ")
    if repeat == 'no':
        polling_active = False
# 调查结束，显示结果。
print("\n--- Poll Results ---")
for name, response in responses.items():
    print(f"{name} would like to climb {response}.")