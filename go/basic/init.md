# init 函数
init 函数在程序中的执行顺序
1. 全局变量
2. init()
3. main()

多个 init 函数执行顺序
1. 不同 package 中的多个 init()，按依赖顺序依次执行，被依赖层级最深的 package 的 init() 最先被执行
2. 同一个 package 多个文件中的 init() 按文件字典序从小到大的顺序执行
3. 同文件中的多个 init() 按从上到下的顺序执行

注意点
1. 导入 package 不用时，例如 import _ foo，其中 foo 包中的 init() 也会被执行
2. init() 函数不能显示直接调用

init 可能导致的一些问题
1. 限制错误管理，只能在 init() 中进行 panic
2. 有时单元测试的时候不需要执行 init() 中的依赖逻辑
3. 在 init() 中设置一些初始化状态的时候，必须要通过全局变量，但是通过普通函数返回具体实例的方式会更好

init 适用的场景
1. 做一些无依赖的静态化配置
