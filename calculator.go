package main;

import (
	"fmt"
)

func Sum(i ...int) int {
	var total = 0;
	for j := range i {
		total += i[j];
	}
	return total;
}

func Subtract(i ...int) int {
	var total = 0;
	for j := range i {
		total -= i[j];
	}
	return total;
}

func Multiply(i ...int) int {
	var total = 1;
	for j := range i {
		total *= i[j];
	}
	return total;
}

func Divide(i ...int) float64 {
	var total = float64(i[0]);
	for j := range i {
		if j == 0 {
			continue;
		}
		total /= float64(i[j]);
	}
	return total;
}

func main() {

	var totalSum = Sum(1, 2, 3);
	var totalSub = Subtract(1, 2, 3);
	var totalMult = Multiply(1, 2, 3, 4);
	var totalDiv = Divide(10, 2, 2, 2, 2, 2);

	fmt.Println("Soma total:", totalSum);
	fmt.Println("Sub total:", totalSub);
	fmt.Println("Mult total:", totalMult);
	fmt.Println("Div total:", totalDiv);
}