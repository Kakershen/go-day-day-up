package main

import "fmt"

/*
append函数是用来在slice末尾追加一个或者多个元素。
当追加元素时，发现slice的len>cap时，会重新开辟一个2*cap的内存空间去存储追加过元素的slice。
如果追加元素后slice的len<=cap,则append返回的新生成的slice的内存地址依旧是传入的slice参数的内存地址。
*/

/*
sliceApend函数输出：
s1 = [], s1内存地址 = 0xc000098060, len = 0, cap = 0
s1 = [1], s1内存地址 = 0xc000098060, len = 1, cap = 1
s1 = [1 1], s1内存地址 = 0xc000098060, len = 2, cap = 2
s1 = [1 1 1], s1内存地址 = 0xc000098060, len = 3, cap = 4
*/

func sliceAppend() {
	s1 := []int{}
	fmt.Printf("s1 = %v, s1内存地址 = %p, len = %d, cap = %d \n", s1, &s1, len(s1), cap(s1))

	s1 = append(s1, 1)
	fmt.Printf("s1 = %v, s1内存地址 = %p, len = %d, cap = %d \n", s1, &s1, len(s1), cap(s1))

	s1 = append(s1, 1)
	fmt.Printf("s1 = %v, s1内存地址 = %p, len = %d, cap = %d \n", s1, &s1, len(s1), cap(s1))

	s1 = append(s1, 1)
	fmt.Printf("s1 = %v, s1内存地址 = %p, len = %d, cap = %d \n", s1, &s1, len(s1), cap(s1))
}

func slice_output() {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for key, value := range slice {
		fmt.Printf("key=%v, value=%v \n", key, value)
		m[key] = &value
	}

	fmt.Println(*m[0])
	fmt.Println(*m[1])
	fmt.Println(*m[2])
	fmt.Println(*m[3])

	scene := make(map[string]int)
	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960
	for k, v := range scene {
		fmt.Println(k, v)
		fmt.Printf("v的地址 = %v \n", &v)
	}
}
