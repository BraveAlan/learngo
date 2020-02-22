package main

import "fmt"

// 数组是值类型，这意味着修改数组的值，外面是不会有变化的
// [10]int和[20]int 是不同类型
func printArray(arr [5]int) {
	//for i:=0; i < len(arr); i++ {
	//	fmt.Println(arr[i])
	//}
	//
	//for i := range arr {
	//	fmt.Println(arr[i])
	//}

	for i, v := range arr {
		fmt.Println(i, v)
	}

	//for _, v := range arr {
	//	fmt.Println(v)
	//}
}

func printArrayPointer(arr *[5]int) {
	arr[0] = 100
	(*arr)[1] = 101
	for i, v := range arr {
		fmt.Println(i, v)
	}

}

func main() {

	// 数组定义
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2,4,6,8,10} // 让编译器帮我们数有多少个int

	var grid [4][5]int
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	// 遍历数组
	printArray(arr1)
	printArray(arr3)
	//printArray(arr2) //报错
	printArrayPointer(&arr1)

}
