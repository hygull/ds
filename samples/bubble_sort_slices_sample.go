/*
	{
		"date_of_creation" => "23 Dec 2016, Friday",
		"aim_of_program"   => "To sort a slice of integers & floats using ds package",
		"coded_by"         => "Rishikesh Agrawani",
		"Go_version"	   => "1.7.1",
	}
*/
package main

import (
	"fmt"
	"github.com/hygull/ds"
)

func main() {
	//Creating a slice of integers...int8/int32/int64 are also allowed
	wholeNums := []int{16, 13, 34, 65, 78, 43, 1, 98, 0}
	fmt.Println("Before applying bubble sort : ", wholeNums)
	ds.BubbleSort(wholeNums)
	fmt.Println("After  applying bubble sort : ", wholeNums)
	fmt.Println()

	//Creating a slice of float32
	floatsList := []float32{16.5999, 13.21, -34.459, 65.99, 78.09, 23.32, 1.43, 56, 0.9}
	fmt.Println("Before applying bubble sort : ", floatsList)
	ds.BubbleSort(floatsList)
	fmt.Println("After  applying bubble sort : ", floatsList)
	fmt.Println()

	//Creating a slice of float64
	floatsList2 := []float64{16.5999, 0.0, 1653.21, 3454544.459, -0.8, 655675.99, 78.09, 23.32, 1.43, 98, -0.9}
	fmt.Println("Before applying bubble sort : ", floatsList2)
	ds.BubbleSort(floatsList2)
	fmt.Println("After  applying bubble sort : ", floatsList2)
	fmt.Println()

	//Creating an empty slice of float64
	floatsList3 := []float64{} //Enmpty slice
	fmt.Println("Before applying bubble sort : ", floatsList3)
	ds.BubbleSort(floatsList3)
	fmt.Println("After  applying bubble sort : ", floatsList3)
	fmt.Println()

	//Creating a slice of strings
	names := []string{"hygull", "rob", "robert", "ken"}
	fmt.Println("Before applying bubble sort : ", names)
	ds.BubbleSort(names)
	fmt.Println("After  applying bubble sort : ", names)
	fmt.Println()
}

/* OUTPUT :-

Before applying bubble sort :  [16 13 34 65 78 43 1 98 0]
After  applying bubble sort :  [0 1 13 16 34 43 65 78 98]

Before applying bubble sort :  [16.5999 13.21 -34.459 65.99 78.09 23.32 1.43 56 0.9]
After  applying bubble sort :  [-34.459 0.9 1.43 13.21 16.5999 23.32 56 65.99 78.09]

Before applying bubble sort :  [16.5999 0 1653.21 3.454544459e+06 -0.8 655675.99 78.09 23.32 1.43 98 -0.9]
After  applying bubble sort :  [-0.9 -0.8 0 1.43 16.5999 23.32 78.09 98 1653.21 655675.99 3.454544459e+06]

Before applying bubble sort :  []
After  applying bubble sort :  []

Before applying bubble sort :  [hygull rob robert ken]
[hygull/ds] Got  => [hygull rob robert ken]
[hygull/ds] Type => []string
[hygull/ds] You have to pass any one among these => []int, []int8,[]int16,[]int32,[]int64,[]float32,[]float64
panic: Got an unwanted slice to sort

goroutine 1 [running]:
panic(0x9bc40, 0xc420074420)
	/usr/local/go/src/runtime/panic.go:500 +0x1a1
ds.BubbleSort(0x93d00, 0xc42006e260)
	/Users/admin/projects/Go/GoWorkspace/src/ds/bubble_sort_slices.go:153 +0x2f1
main.main()
	/Users/admin/projects/Go/GoFiles/switch_type_check5.go:48 +0xd34
exit status 2

*/
