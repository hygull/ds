package ds

import (
	"errors"
	"fmt"
	"log"
)

func (iSlice IntSlice) GetSum() interface{} {
	sum := 0
	for _, num := range iSlice {
		sum += num
	}
	return sum
}
func (i8Slice Int8Slice) GetSum() interface{} {
	sum := int8(0)
	for _, num := range i8Slice {
		sum += num
	}
	return sum
}
func (i16Slice Int16Slice) GetSum() interface{} {
	sum := int16(0)
	for _, num := range i16Slice {
		sum += num
	}
	return sum
}
func (i32Slice Int32Slice) GetSum() interface{} {
	sum := int32(0)
	for _, num := range i32Slice {
		sum += num
	}
	return sum
}
func (i64Slice Int64Slice) GetSum() interface{} {
	sum := int64(0)
	for _, num := range i64Slice {
		sum += num
	}
	return sum
}
func (uiSlice UintSlice) GetSum() interface{} {
	sum := uint(0)
	for _, num := range uiSlice {
		sum += num
	}
	return sum
}
func (ui8Slice Uint8Slice) GetSum() interface{} {
	sum := uint8(0)
	for _, num := range ui8Slice {
		sum += num
	}
	return sum
}
func (ui16Slice Uint16Slice) GetSum() interface{} {
	sum := uint16(0)
	for _, num := range ui16Slice {
		sum += num
	}
	return sum
}
func (ui32Slice Uint32Slice) GetSum() interface{} {
	sum := uint32(0)
	for _, num := range ui32Slice {
		sum += num
	}
	return sum
}
func (ui64Slice Uint64Slice) GetSum() interface{} {
	sum := uint64(0)
	for _, num := range ui64Slice {
		sum += num
	}
	return sum
}

func (f32Slice Float32Slice) GetSum() interface{} {
	sum := float32(0.0)
	for _, num := range f32Slice {
		sum += num
	}
	return sum
}
func (f64Slice Float64Slice) GetSum() interface{} {
	sum := 0.0
	for _, num := range f64Slice {
		sum += num
	}
	return sum
}
func Sum(intsSlice interface{}) interface{} {

	switch t := intsSlice.(type) {
	case []int:
		return IntSlice(intsSlice.([]int)).GetSum()
	case []int8:
		return Int8Slice(intsSlice.([]int8)).GetSum()
	case []int16:
		return Int16Slice(intsSlice.([]int16)).GetSum()
	case []int32:
		return Int32Slice(intsSlice.([]int32)).GetSum()
	case []int64:
		return Int64Slice(intsSlice.([]int64)).GetSum()
	case []uint:
		return UintSlice(intsSlice.([]uint)).GetSum()
	case []uint8:
		return Uint8Slice(intsSlice.([]uint8)).GetSum()
	case []uint16:
		return Uint16Slice(intsSlice.([]uint16)).GetSum()
	case []uint32:
		return Uint32Slice(intsSlice.([]uint32)).GetSum()
	case []uint64:
		return Uint64Slice(intsSlice.([]uint64)).GetSum()
	case []float64:
		return Float64Slice(intsSlice.([]float64)).GetSum()
	case []float32:
		return Float32Slice(intsSlice.([]float32)).GetSum()
	default:
		fmt.Println(t)
		log.Fatal(errors.New("Got an invalid slice, required a slice of integers/floats/unsigned integers"))
	}
	return 0
}
