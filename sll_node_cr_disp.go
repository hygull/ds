package ds

import "fmt"

//Defining the singly linked list Node
type Node struct {
	Record map[string]interface{}
	Next   *Node
}

//A function that creates a node for singly lint based on the information provided (as list of data)
func CreateNode(compulsoryData interface{}, data ...interface{}) *Node {
	fmt.Print(compulsoryData, data)
	tempMap := make(map[string]interface{})
	node := Node{}
	dataName := "Data"

	data = append(data, compulsoryData)

	for index, item := range data {
		switch t := item.(type) {
		case int:
			tempMap[dataName+fmt.Sprint(index)] = item.(int)
		case []int:
			tempMap[dataName+fmt.Sprint(index)] = item.([]int)
		case int8:
			tempMap[dataName+fmt.Sprint(index)] = item.(int8)
		case []int8:
			tempMap[dataName+fmt.Sprint(index)] = item.([]int8)
		case int16:
			tempMap[dataName+fmt.Sprint(index)] = item.(int16)
		case []int16:
			tempMap[dataName+fmt.Sprint(index)] = item.([]int16)
		case int32:
			tempMap[dataName+fmt.Sprint(index)] = item.(int32)
		case []int32:
			tempMap[dataName+fmt.Sprint(index)] = item.([]int32)
		case int64:
			tempMap[dataName+fmt.Sprint(index)] = item.(int64)
		case []int64:
			tempMap[dataName+fmt.Sprint(index)] = item.([]int64)
		case uint:
			tempMap[dataName+fmt.Sprint(index)] = item.(uint)
		case uint8:
			tempMap[dataName+fmt.Sprint(index)] = item.(uint8)
		case uint16:
			tempMap[dataName+fmt.Sprint(index)] = item.(uint16)
		case uint32:
			tempMap[dataName+fmt.Sprint(index)] = item.(uint32)
		case uint64:
			tempMap[dataName+fmt.Sprint(index)] = item.(uint64)
		case []uint:
			tempMap[dataName+fmt.Sprint(index)] = item.([]uint)
		case []uint8:
			tempMap[dataName+fmt.Sprint(index)] = item.([]uint8)
		case []uint16:
			tempMap[dataName+fmt.Sprint(index)] = item.([]uint16)
		case []uint32:
			tempMap[dataName+fmt.Sprint(index)] = item.([]uint32)
		case []uint64:
			tempMap[dataName+fmt.Sprint(index)] = item.([]uint64)
		case bool:
			tempMap[dataName+fmt.Sprint(index)] = item.(bool)
		case []bool:
			tempMap[dataName+fmt.Sprint(index)] = item.([]bool)
		case uintptr:
			tempMap[dataName+fmt.Sprint(index)] = item.(uintptr)
		case []uintptr:
			tempMap[dataName+fmt.Sprint(index)] = item.([]uintptr)
		case float32:
			tempMap[dataName+fmt.Sprint(index)] = item.(float32)
		case []float32:
			fmt.Print("gh")
			tempMap[dataName+fmt.Sprint(index)] = item.([]float32)
		case float64:
			fmt.Print("gh64")
			tempMap[dataName+fmt.Sprint(index)] = item.(float64)
		case []float64:
			tempMap[dataName+fmt.Sprint(index)] = item.([]float64)
		case string:
			tempMap[dataName+fmt.Sprint(index)] = item.(string)
		case []string:
			tempMap[dataName+fmt.Sprint(index)] = item.([]string)

		/* Marshalling problem occurs, So the follwing is commented */

		// case complex64:
		// 	tempMap[dataName+fmt.Sprint(index)] = item.(complex64)
		// case complex128:
		// 	tempMap[dataName+fmt.Sprint(index)] = item.(complex128)
		// case []complex64:
		// 	tempMap[dataName+fmt.Sprint(index)] = item.([]complex64)
		// case []complex128:
		// 	tempMap[dataName+fmt.Sprint(index)] = item.([]complex128)

		default:
			fmt.Println("Got ", t)
			panic("An illegal data value")
		}
	}
	node.Record = tempMap
	/* node.next = nil */
	return &node
}
