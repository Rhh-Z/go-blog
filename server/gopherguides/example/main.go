package main

import (
	"fmt"

	greet "github.com/mymodule/gopherguides"
	"github.com/spf13/cobra"
)

// const (
// 	x = iota
// 	_
// 	y
// 	z = "zz"
// 	k
// 	p = iota
// )
// const (
// 	name = "name"
// 	c    = iota
// 	d    = iota
// )

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (u *User) ShowUserName() {
	fmt.Println(u.Name)
}

// type MyGiftsData struct {
// 	Num    int  `json:"num"`
// 	HasGet bool `json:"has_get"`
// }

// type MyGiftsList []MyGiftsData

// func (m MyGiftsList) Len() int {
// 	return len(m)
// }

// func (m MyGiftsList) Swap(i, j int) {
// 	m[i], m[j] = m[j], m[i]
// }

// func (m MyGiftsList) Less(i, j int) bool {
// 	return m[i].Num > m[j].Num
// }

func main() {
	// xin := &User{
	// 	Name: "xin",
	// 	Age:  23,
	// }

	a := "123"
	b := "abc"
	c := a + b
	fmt.Println(c)
	// xin.ShowUserName()

	cmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello, Modules!")

			greet.Hello()
		},
	}

	fmt.Println("Calling cmd.Execute()!")
	cmd.Execute()

	// fmt.Println(x, y, z, k, p)
	// fmt.Println(c, d)

	// name, test := 123
	// greet.DemoCodeCaptchaCreate()
	// fmt.Println(greet.Shark)

	// oct := greet.Octopus{
	// 	Name:  "andy",
	// 	Color: "red",
	// }
	// fmt.Println(oct.String())

	// var m1 = map[string]string{"zhangsan": "10220308"}
	// fmt.Println(m1)

	// map2 := make([]bool, 4)
	// fmt.Println(map2)

	// var st = "13235"
	// cast.ToInt(st)
	// fmt.Println(strings.Split(st, "3"))
	// fmt.Println(strings.Replace(st, "3", "2", 1))
	// var obj = map[string]int{
	// 	"a": 1,
	// 	"b": 2,
	// }
	// x := obj["a"]
	// fmt.Println(obj)
	// delete(obj, "b")
	// fmt.Println(obj)
	// fmt.Println(x)

	// value, ok := map[key]
	// fmt.Println(value, ok)
	// slice := []int{123, 1321, 1113, 4, 5}

	// for value := range slice {
	// 	fmt.Println(value)
	// }
	// sayDefer()
	// getScore(91)

	// fmt.Println(runtime.GOARCH)
}

// func sayDefer() {
// 	defer sayEnd()
// 	sayStart()
// }

// func getScore(value int) {
// 	switch score := value; {
// 	case score >= 90:
// 		fmt.Println("优秀")
// 	case score >= 60:
// 		fmt.Println("及格")
// 	default:
// 		fmt.Println("不及格")
// 	}
// }

// func sayStart() {
// 	fmt.Println("start")
// }
