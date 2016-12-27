/*
	{
		"date_of_creation" => "27 Dec 2016, Tuesday(morning)",
		"aim_of_program"   => "Testing nodes display functions of singly linked list",
		"coded_by"         => "Rishikesh Agrawani",
		"Go_version"	   => "1.7.1",
	}
*/
package ds

import "testing"
import "fmt"

func TestShowNodes(t *testing.T) {

	node := CreateNode("Gopher", 24, 45.56, true, []int8{2, 5, 6, 7, 8, 2, 1}) //Complex data causes an error while marshalling

	node.Next = CreateNode([]string{"Bangalore", "Delhi", "Raipur"}, []float64{78.98, 65.43, 653.76}, []int8{2, 5, 6, 7, 8, 2, 1}) //Complex data causes an error while marshalling

	node.Next.Next = CreateNode([]rune{18, 9, 19, 8, 9}, "Tell me about Golang", []float32{45.3, 76.54, 34.21, 9.87}, 76.45, 67, 'A')
	ShowNodes(node)
}

func TestShowNodesHierarchy(t *testing.T) {

	node := CreateNode("Gopher", 24, 45.56, true, []int8{2, 5, 6, 7, 8, 2, 1}) //Complex data causes an error while marshalling

	node.Next = CreateNode([]string{"Bangalore", "Delhi", "Raipur"}, []float64{78.98, 65.43, 653.76}, []int8{2, 5, 6, 7, 8, 2, 1}) //Complex data causes an error while marshalling

	node.Next.Next = CreateNode([]rune{18, 9, 19, 8, 9}, "Tell me about Golang", []float32{45.3, 76.54, 34.21, 9.87}, 76.45, 67, 'A')
	ShowNodesHierarchy(node)
}

func TestGetNodesRecords(t *testing.T) {

	node := CreateNode("Gopher", 24, 45.56, true, []int8{2, 5, 6, 7, 8, 2, 1}) //Complex data causes an error while marshalling

	node.Next = CreateNode([]string{"Bangalore", "Delhi", "Raipur"}, []float64{78.98, 65.43, 653.76}, []int8{2, 5, 6, 7, 8, 2, 1}) //Complex data causes an error while marshalling

	node.Next.Next = CreateNode([]rune{18, 9, 19, 8, 9}, "Tell me about Golang", []float32{45.3, 76.54, 34.21, 9.87}, 76.45, 67, 'A')
	fmt.Println("\nList of maps, each showing the data part of a node of singly linked list")
	recordsList := GetNodesRecords(node)
	fmt.Println(recordsList)
	marshalledData, _ := Marshall(recordsList)
	fmt.Println(marshalledData)

	var start *Node //nil
	records := GetNodesRecords(start)
	if len(records) == 0 {
		fmt.Println("Singly linked list is empty")
	}
}
