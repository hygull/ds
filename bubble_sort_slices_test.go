/*
	@auhthor : Rishikesh Agrawani
	@coded_on: 23/12/2016, Friday
*/
package ds

import "testing"
import "reflect"
import "github.com/fatih/color"

//A function to test the BubbleSort() function of ds package
func TestBubbleSort(t *testing.T) {
	type Slices struct {
		IntSlice     []int
		Int8Slice    []int8
		Int16Slice   []int16
		Int32Slice   []int32
		Int64Slice   []int64
		Float32Slice []float32
		Float64Slice []float64
		//StringSlice   []string
		Float64Slice2 []float32
		Int64Slice2   []int64
	}
	slices := Slices{
		[]int{12, -3, 34, 0, 56},
		[]int8{111, 12, 67, 98, -15},
		[]int16{34, -43, 12, 454, 0, 56, 23},
		[]int32{29, 23, 0, 12, 50, -12, 100},
		[]int64{12, -32, 122, 10, -2, 1},
		[]float32{12.07, -34.87, 87.234, 123.99, 32},
		[]float64{1, 3.3, 6, 7, 9.5},
		//[]string{"hygull", "rob", "robert", "ken"},
		[]float32{},
		[]int64{},
	}
	e := reflect.ValueOf(&slices).Elem()
	numOfFields := e.NumField()
	for i := 0; i < numOfFields; i++ {

		color.Green("Before sorting : %v", e.Field(i).Interface()) //In case of success, this message will not be displayed

		BubbleSort(e.Field(i).Interface())

		color.Blue("After sorting  : %v", e.Field(i).Interface()) //In case of success, this message will not be displayed

	}
}

/*
Before sorting : [12 -3 34 0 56]
After sorting  : [-3 0 12 34 56]
Before sorting : [111 12 67 98 -15]
After sorting  : [-15 12 67 98 111]
Before sorting : [34 -43 12 454 0 56 23]
After sorting  : [-43 0 12 23 34 56 454]
Before sorting : [29 23 0 12 50 -12 100]
After sorting  : [-12 0 12 23 29 50 100]
Before sorting : [12 -32 122 10 -2 1]
After sorting  : [-32 -2 1 10 12 122]
Before sorting : [12.07 -34.87 87.234 123.99 32]
After sorting  : [-34.87 12.07 32 87.234 123.99]
Before sorting : [1 3.3 6 7 9.5]
After sorting  : [1 3.3 6 7 9.5]
Before sorting : []
After sorting  : []
Before sorting : []
After sorting  : []

*/
