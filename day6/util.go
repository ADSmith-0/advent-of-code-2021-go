package day6

import "fmt"

const NUMBER_OF_DAYS = 256

func printArr(arr *[]int) {
	for i, value := range *arr {
		if i == len(*arr)-1 {
			fmt.Printf("%d\n", value)
			continue
		}
		fmt.Printf("%d,", value)
	}
}
