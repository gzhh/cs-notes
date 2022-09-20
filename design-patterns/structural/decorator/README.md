## Decorator Pattern

Decorator structural pattern allows extending the function of an existing object dynamically without altering its internals.

Decorators provide a flexible method to extend functionality of objects.

装饰模式使用对象组合的方式动态改变或增加对象行为。

Go语言借助于匿名组合和非入侵式接口可以很方便实现装饰模式。

使用匿名组合，在装饰器中不必显式定义转调原对象方法。

### Rules of Thumb
- Unlike Adapter pattern, the object to be decorated is obtained by injection.
- Decorators should not alter the interface of an object.
