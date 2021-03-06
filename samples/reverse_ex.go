/*
	{
		"date_of_creation" => "25 Dec 2016, Sunday",
		"aim_of_program"   => "To use Reverse() of ds package",
		"coded_by"         => "hygull",
		"Go_version"	   => "1.7.1",
	}
*/
package main

import "fmt"

import "github.com/hygull/ds"

func main() {
	intSlice := []int{34, 89, 21, 11, 98, 43, 24, 1, -43, 67, -2, -23}
	fmt.Println(intSlice)
	ds.Reverse(intSlice)
	fmt.Println(intSlice)
	fmt.Print("\n")

	int8Slice := []int{14, 11, 98, -43, 24, 91, -43, 679, -25, -23}
	fmt.Println(int8Slice)
	ds.Reverse(int8Slice)
	fmt.Println(int8Slice)
	fmt.Print("\n")

	int16Slice := []int{0, 34, 89, 21, 67, 11, 98, -2, -23}
	fmt.Println(int16Slice)
	ds.Reverse(int16Slice)
	fmt.Println(int16Slice)
	fmt.Print("\n")

	int32Slice := []int{354, 89, 21, 111, 98, 43, 243, 1, -463, 67, -2, -243}
	fmt.Println(int32Slice)
	ds.Reverse(int32Slice)
	fmt.Println(int32Slice)
	fmt.Print("\n")

	int64Slice := []int{0, 34, 849, 219, 11, 98, 43, 24, 1, -43, 67, -2, -283}
	fmt.Println(int64Slice)
	ds.Reverse(int64Slice)
	fmt.Println(int64Slice)
	fmt.Print("\n")

	strSlice := []string{"Delhi", "Raipur", "Bangalore", "Jaipur", "Kondagaon", "Newyork", "America", "China"}
	fmt.Println(strSlice)
	ds.Reverse(strSlice)
	fmt.Println(strSlice)
	fmt.Print("\n")

	float64Slice := []float64{0, 34.2, 8.3, 219, 11.9, 98.4, 4.3, 24, 1, -43.8, 6.7, -2, -28.3}
	fmt.Println(float64Slice)
	ds.Reverse(float64Slice)
	fmt.Println(float64Slice)
	fmt.Print("\n")

	intSlice2 := []int{}
	fmt.Println(intSlice2)
	ds.Reverse(intSlice2)
	fmt.Println(intSlice2)
	fmt.Print("\n")

	strBytes := []byte("Rishikesh Agrawani")
	fmt.Println(string(strBytes))
	ds.Reverse(strBytes)
	fmt.Println(string(strBytes))
}
