# Slice

## 原理
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
