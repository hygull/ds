/*
	{
		"date_of_creation" => "25 Dec 2016, Sunday(morning)",
		"aim_of_program"   => "To test  reversing the slices and string(by converting it into bytes )",
		"coded_by"         => "hygull",
		"Go_version"	   => "1.7.1",
	}
*/
package ds

import "fmt"
import "testing"

func TestReverse(t *testing.T) {
	intSlice := []int{34, 89, 21, 11, 98, 43, 24, 1, -43, 67, -2, -23}
	fmt.Println(intSlice)
	err := Reverse(intSlice)
	if err != nil {
		fmt.Println("Error => ", err)
	}
	fmt.Println(intSlice)
	fmt.Print("\n")

	int8Slice := []int{14, 11, 98, -43, 24, 91, -43, 679, -25, -23}
	fmt.Println(int8Slice)
	Reverse(int8Slice)
	fmt.Println(int8Slice)
	fmt.Print("\n")

	int16Slice := []int{0, 34, 89, 21, 67, 11, 98, -2, -23}
	fmt.Println(int16Slice)
	Reverse(int16Slice)
	fmt.Println(int16Slice)
	fmt.Print("\n")

	int32Slice := []int{354, 89, 21, 111, 98, 43, 243, 1, -463, 67, -2, -243}
	fmt.Println(int32Slice)
	Reverse(int32Slice)
	fmt.Println(int32Slice)
	fmt.Print("\n")

	int64Slice := []int{0, 34, 849, 219, 11, 98, 43, 24, 1, -43, 67, -2, -283}
	fmt.Println(int64Slice)
	Reverse(int64Slice)
	fmt.Println(int64Slice)
	fmt.Print("\n")

	strSlice := []string{"Delhi", "Raipur", "Bangalore", "Jaipur", "Kondagaon", "Newyork", "America", "China"}
	fmt.Println(strSlice)
	Reverse(strSlice)
	fmt.Println(strSlice)
	fmt.Print("\n")

	intSlice2 := []int{}
	fmt.Println(intSlice2)
	Reverse(intSlice2)
	fmt.Println(intSlice2)
	fmt.Print("\n")

	float64Slice := []float64{0, 34.2, 8.3, 219, 11.9, 98.4, 4.3, 24, 1, -43.8, 6.7, -2, -28.3}
	fmt.Println(float64Slice)
	Reverse(float64Slice)
	fmt.Println(float64Slice)
	fmt.Print("\n")

	strBytes := []byte("Rishikesh Agrawani")
	fmt.Println(string(strBytes))
	Reverse(strBytes)
	fmt.Println(string(strBytes))

	uintSlice := []uint{34, 89, 21, 11, 98, 43, 24, 1}
	fmt.Println(uintSlice)
	err = Reverse(uintSlice)
	if err != nil {
		fmt.Println("Error => ", err)
	}
	fmt.Println(uintSlice)
	fmt.Print("\n")

	uint8Slice := []uint8{14, 11, 98, 67, 24, 91}
	fmt.Println(uint8Slice)
	Reverse(uint8Slice)
	fmt.Println(uint8Slice)
	fmt.Print("\n")

	uint16Slice := []uint16{0, 34, 89, 21, 67, 11, 98, 45}
	fmt.Println(uint16Slice)
	Reverse(uint16Slice)
	fmt.Println(uint16Slice)
	fmt.Print("\n")

	uint32Slice := []uint32{354, 89, 21, 111, 98, 43, 243, 1}
	fmt.Println(uint32Slice)
	Reverse(uint32Slice)
	fmt.Println(uint32Slice)
	fmt.Print("\n")

	uint64Slice := []uint64{0, 34, 849, 219, 11, 98, 43, 24, 1, 3}
	fmt.Println(uint64Slice)
	Reverse(uint64Slice)
	fmt.Println(uint64Slice)
	fmt.Print("\n")

	boolSlice := []bool{true, 10 < 9, 23 == 23, false, 45.0 == float64(45), 45.0 == float32(45), float32(46.78) == float32(46.78)}
	fmt.Println(boolSlice)
	Reverse(boolSlice)
	fmt.Println(boolSlice)
	fmt.Print("\n")

	complex64Slice := []complex64{4 + 6i, -3 + 8i, -6 - 4i, 12 - 45i, 12, 8i, -23, -34i, 78.9 + 6i, 23 + 7.9i, -23.7i}
	fmt.Println(complex64Slice)
	Reverse(complex64Slice)
	fmt.Println(complex64Slice)
	fmt.Print("\n")

	//Reverse(34)     /*Illegal request*/

}
