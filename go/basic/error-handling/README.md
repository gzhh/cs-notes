In Go, if an application panics, there is a bug — not an error.
Many developers new to Go confuse a panic with a runtime error in other languages. Go has errors for that. Panics should not be used to check if something works or not. Always use errors for that.
Catching panics can make sense to prevent unknown bugs from crashing a production server. Such a bug needs to be fixed though and using recover to catch it should only be a measure to prevent disaster in high-stake scenarios — never to be used as a way to do “alternative logic” as you put it.

注意点：
- 检查错误的类型
- 检查错误的值
- 不要处理错误两次
- 不要忘了在 defer 中处理错误
