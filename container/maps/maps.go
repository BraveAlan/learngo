package main

import "fmt"

func main() {
	m := map[string]string {
		"name": "jojo",
		"course": "golang",
		"site": "imooc",
		"quality": "notbad",
	}
	fmt.Println(m)

	m2 := make(map[string]int) // m2 == empty map

	var m3 map[string]int // m3 == nil，即使是个nil的值也是可以使用的

	fmt.Println(m, m2, m3)

	fmt.Println("Traversing map")
	for k, v := range m {
		fmt.Println(k, v) // 不保证顺序，这是hashmap
	}

	fmt.Println("Getting Values")
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)
	causeName, ok := m["cause"] // map中不存在的值
	fmt.Println(causeName, ok)

	fmt.Println("Deleting values")
	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")
	name, ok = m["name"]
	fmt.Printf("%q %v\n", name, ok)


}
