package main

import "fmt"

var product = make(map[string]float64)

func main() {
	fmt.Println("product =", product)

	//add
	product["Macbook"] = 40000
	product["iPhone"] = 30000
	product["iPad"] = 18000
	fmt.Println("product =", product)

	//delete
	delete(product, "iPad")
	fmt.Println("product =", product)

	//update
	product["iPhone"] = 25000
	fmt.Println("product =", product)

	//value
	value1 := product["iPhone"]
	fmt.Println("value1 =", value1)

	//map
	courseName := map[string]string{"101": "Java", "102": "Python"}
	fmt.Println("courseName =", courseName)
}
