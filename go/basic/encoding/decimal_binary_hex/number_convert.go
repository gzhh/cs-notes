package main

import (
	"fmt"
	"strconv"
)

// 数字转换：
// 十六进制转十进制
// 二进制转十进制
func NumberConvert() {
	/*
	   a => 10 * 16^0 => 10
	   3 => 3 * 16^1 => 48
	   2 => 2 * 16^2 => 512
	*/
	hexNum := "23a"
	decimalNum, _ := strconv.ParseInt(hexNum, 16, 64)
	fmt.Println(decimalNum)

	/*
	   a => 1 * 2^0 => 1
	   3 => 0 * 2^1 => 0
	   2 => 1 * 2^2 => 4
	*/
	binaryNum := "101"
	decimalNum, _ = strconv.ParseInt(binaryNum, 2, 64)
	fmt.Println(decimalNum)
}
