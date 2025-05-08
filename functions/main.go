package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func subt(a int, b int) (result int) { // result is alreay declraed here, and it will be returned
	result = a - b
	return
}

// error handling

func div(a float64, b float64) (ans float64, err error) {
	if b == 0 {
		return 0, fmt.Errorf("error handling")
	}

	return a / b, nil

}

func main() {
	ans := add(3, 2)
	fmt.Println(ans)
	fmt.Println(subt(5, 3))

	result, err := div(10, 0) // reuslt, _ := div(10,0) - to ignore the error the function is returning

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)

}
