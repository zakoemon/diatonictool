package main

import "fmt"

type Hoge struct {
	Age     int
	Name    string
	Friends []string
}

/*
Entry point of the application
*/
func main() {
	// fmt.Println("hi")
	// types, _ := json.GetEntity()
	// for _, v := range types {
	// 	fmt.Println(v)
	// }

	hoge := Hoge{Age: 1, Name: "hello"}
	hoge.Friends = []string{}

	fmt.Println(&hoge)
	// test(&hoge)

}

// func test(hoge *Hoge) {
// 	fmt.Println(&hoge.Friends)
// }
