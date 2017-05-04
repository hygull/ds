/*
	@auhthor : Rishikesh Agrawani
	@coded_on: 25/12/2016, Sunday
*/
package ds

import "fmt"

//A function that that reverses the order of elements in int/float/string slice & also reverses the string
func Reverse(data interface{}) error {
	switch t := data.(type) {
	case []int:
		a := data.([]int)
		for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
			a[i], a[j] = a[j], a[i]
		}
	case []int8:
		a := data.([]int8)
		for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
			a[i], a[j] = a[j], a[i]
		}
	case []int16:
		a := data.([]int16)
		for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
			a[i], a[j] = a[j], a[i]
		}
	case []int32:
		a := data.([]int32)
		for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
			a[i], a[j] = a[j], a[i]
		}
	case []int64:
		a := data.([]int64)
		for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
			a[i], a[j] = a[j], a[i]
		}
	case []uint:
		a := data.([]uint)
		for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
			a[i], a[j] = a[j], a[i]
		}
	case []uint8:
		a := data.([]uint8)
		for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
			a[i], a[j] = a[j], a[i]
		}
	case []uint16:
		a := data.([]uint16)
		for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
			a[i], a[j] = a[j], a[i]
		}
	case []uint32:
		a := data.([]uint32)
		for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
			a[i], a[j] = a[j], a[i]
		}
	case []uintptr:
		a := data.([]uintptr)
		for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
			a[i], a[j] = a[j], a[i]
		}
	case []float32:
		a := data.([]float32)
		for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
			a[i], a[j] = a[j], a[i]
		}
	case []float64:
		a := data.([]float64)
		for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
			a[i], a[j] = a[j], a[i]
		}

	case []string:
		a := data.([]string)
		for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
			a[i], a[j] = a[j], a[i]
		}
	case []bool:
		a := data.([]bool)
		for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
			a[i], a[j] = a[j], a[i]
		}
	case []complex64:
		a := data.([]complex64)
		for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
			a[i], a[j] = a[j], a[i]
		}
	case []complex128:
		a := data.([]complex128)
		for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
			a[i], a[j] = a[j], a[i]
		}
	default:
		return fmt.Errorf("%v %v %v", "Got ", t, " an invalid slice/string to reverse")
	}
	return nil
}

/* a := data.([]bool)  ==>  a := data.([]string).......It will show the error like below
--- FAIL: TestReverse (0.00s)
panic: interface conversion: interface is []bool, not []string [recovered]
	panic: interface conversion: interface is []bool, not []string
*/
