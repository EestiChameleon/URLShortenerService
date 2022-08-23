package main

import (
	"fmt"
	"strings"
)

//func main() {
//	//wg := sync.WaitGroup{}
//	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9} //(или до 1000 тут можно накидать генератором)
//	for _, i := range arr {
//		//wg.Add(1)
//		go func(x int) {
//			//defer wg.Done()
//			time.Sleep(time.Millisecond * time.Duration(x))
//			println(x)
//		}(i)
//	}
//	//wg.Wait()
//	//fmt.Println(someFunc(1,2))
//}

func main() {
	fmt.Println(someFunc(1, 1))
	s := "http://local/host"
	ss := strings.Split(s, "/")
	fmt.Println(ss[len(ss)-1])
}

func someFunc(a, b int) (c, d int) {
	defer func() { c = 5 }()
	defer func() { d = 7 }()

	c = a * 2
	d = b * 3
	return c, d
}
