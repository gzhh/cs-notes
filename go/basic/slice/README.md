# Slice
- Go Slices: usage and internals https://go.dev/blog/slices-intro
- Arrays, slices (and strings): The mechanics of 'append' https://go.dev/blog/slices


## 使用
- Go Wiki: SliceTricks https://go.dev/wiki/SliceTricks
- https://pkg.go.dev/slices
- https://pkg.go.dev/golang.org/x/exp/slices

## 原理
源码
- https://go.dev/src/runtime/slice.go

深入理解
- 快速理解Go数组和切片的内部实现原理 https://i6448038.github.io/2018/08/11/array-and-slice-principle/
- 你真的了解go语言中的切片吗？https://mp.weixin.qq.com/s/uNajVcWr4mZpof1eNemfmQ
- 数组与切片的面试总结和性能提升点 https://mp.weixin.qq.com/s/wL9menRiuv7dZ1zIOqozYA


### 扩容机制
切片扩容 cap 增长示例
```
func main() {
	s := make([]int, 0, 10)
	carry := 1
	for i := 0; i < 12; i++ {
		fmt.Printf("len: %d, cap: %d\n", len(s), cap(s))
		tmpS := make([]int, 0, carry)
		for j := 0; j < carry; j++ {
			tmpS = append(tmpS, j)
		}
		carry = carry << 1
		s = append(s, tmpS...)
		time.Sleep(time.Second)
	}
	fmt.Printf("len: %d, cap: %d\n", len(s), cap(s))
}
```
