package main

import "fmt"

func main() {
	//fmt.Println("Program is running...")

	// var a interface{} = 5
	// var b interface{} = 6.5
	// var c interface{} = "AA"
	// var d interface{} = [3]int{1, 2, 3}

	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)

	var e interface{} = map[string]interface{}{
		"first":     10,
		"second":    20.7,
		"something": "Nahid",
		"other": map[string]interface{}{
			"a": 4,
			"b": 8,
		},
	}
	var eNew = e.(map[string]interface{})["other"].(map[string]int)["a"]
	fmt.Println(eNew)

	// var e1 = e.(map[string]interface{})
	// var e2 = e1["other"]
	// var n1 = e2.(map[string]int)["a"]

	// fmt.Println(n1)

	// var eNew = e.(map[string]interface{})
	// fmt.Println(eNew["something"])

	// mp := e.(map[string]int)["first"]
	// fmt.Println(mp)

	// f := map[string]int{
	// 	"first":  10,
	// 	"second": 20,
	// }
	// fmt.Println(f["first"])
}
