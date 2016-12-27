/*
	#created_on : 27 Dec 2016,Tuesday(afternoon)
	#coded_by   : hygull
*/
package ds

import (
	"fmt"
	"testing"
)

func TestInsertNodeAtBeginningAndEnd(t *testing.T) {
	fmt.Println("\n********* Testing (Inserting nodes at beginning & end of the linked list) **********\n")
	node := CreateNode("Hygull", 24)
	start := InsertSllNodeAtBeginning(node) //We can also use the global START variable
	/*
				+---------------+----+
		start-->| {"Hygull",24} | nil|
				+--------------------+
	*/
	node2 := CreateNode("Hemkesh", 22)
	start = InsertSllNodeAtBeginning(node2)
	/*
					+---------------+----+    +----------------+----+
		   start===>| {"Hemkesh",22}|  ======>| {"Hygull", 24} | nil|
					+---------------+----+    +----------------+----+
	*/
	node3 := CreateNode("Darshan", 23)
	start = InsertSllNodeAtEnd(node3)
	/*
					+---------------+----+    +----------------+----+     +----------------+----+
		   start===>| {"Hemkesh",22}|  ======>| {"Hygull", 24} |  =======>| {"Darshan",23} | nil|
					+---------------+----+    +----------------+----+     +----------------+----+
	*/
	ShowNodes(start)
	fmt.Println()

	ShowNodesHierarchy(START) //using global START variable
	fmt.Println()

	recordsList := GetNodesRecords(start)
	if len(recordsList) == 0 {
		fmt.Println("Singly linked list is empty")
	}
	fmt.Println()

	marshalledRecords, _ := Marshall(recordsList)
	fmt.Println(marshalledRecords)

	start = nil
	err := ShowNodes(start)
	if err != nil {
		fmt.Println(err) //singly linked list is empty
	}

}
