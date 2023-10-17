"""
装饰器：
https://www.liaoxuefeng.com/wiki/1016959663602400/1017451662295584

"""

"""
def log(func):
    def wrapper(*args, **kw):
        print('call %s():' % func.__name__)
        return func(*args, **kw)
    return wrapper

@log # 把@log放到now()函数的定义处，相当于执行了语句：now = log(now)
def now():
    print('2015-3-25')

if __name__ == '__main__':
    now()
"""

class Thing1:
    def __init__(self, my_word):
        self._word = my_word 

    def word(self):
        return self._word
    
class Thing2:
    def __init__(self, my_word):
        self._word = my_word 

    @property
    def word(self):
        return self._word

if __name__ == '__main__':
    thing1 = Thing1('ok')
    print(thing1.word())

    thing2 = Thing2('ok')
    print(thing2.word)