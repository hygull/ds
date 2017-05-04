/*
	#created_on : 27 Dec 2016,Tuesday(afternoon)
	#coded_by   : hygull
*/
package ds

import (
	"fmt"
	"testing"
)

func TestInsertAndDeleteNodeAtBeginningAndEnd(t *testing.T) {
	/* CASE 1*/
	var start2 *Node = nil
	err := ShowNodes(start2)
	if err != nil {
		fmt.Println(err) //singly linked list is empty
	}

	/* CASE 2 */
	var start *Node
	fmt.Println("\n********* Testing (Inserting nodes at beginning & end of the linked list) **********\n")
	node := CreateNode("Hygull", 24)
	start = InsertSllNodeAtBeginning(start, node) //We can also use the global START variable
	/*
				+---------------+----+
		        start-->| {"Hygull",24} | nil|
				+--------------------+
	*/
	node2 := CreateNode("Hemkesh", 22)
	start = InsertSllNodeAtBeginning(start, node2)
	/*
				 +---------------+----+    +----------------+----+
		        start===>| {"Hemkesh",22}|  ======>| {"Hygull", 24} | nil|
				 +---------------+----+    +----------------+----+
	*/
	node3 := CreateNode("Darshan", 23)
	start = InsertSllNodeAtEnd(start, node3)
	/*
				 +---------------+----+    +----------------+----+     +----------------+----+
		        start===>| {"Hemkesh",22}|  ======>| {"Hygull", 24} |  =======>| {"Darshan",23} | nil|
				 +---------------+----+    +----------------+----+     +----------------+----+
	*/
	ShowNodes(start)
	fmt.Println()

	ShowNodesHierarchy(start)
	fmt.Println()

	recordsList := GetNodesRecords(start)
	if len(recordsList) == 0 {
		fmt.Println("Singly linked list is empty")
	}
	fmt.Println()

	marshalledRecords, _ := Marshall(recordsList)
	fmt.Println(marshalledRecords)

	/* Deleting 1 node from the beginning of linked list */
	start, _ = DeleteSllNodeFromBeginning(start)
	fmt.Println("After deleting 1 node from beginning")
	fmt.Println(GetNodesRecords(start))

	/* Inserting 1 node at end */
	node4 := CreateNode("Surendra", 24)
	start = InsertSllNodeAtEnd(start, node4)
	fmt.Println("After adding 1 node at end")
	fmt.Println(GetNodesRecords(start))

	/* Deleting 1 node from the end of linked list */
	err = DeleteSllNodeFromEnd(start)
	if err != nil { //List is empty
		fmt.Println(err)
	}
	fmt.Println("After deleting 1 node from end")
	fmt.Println(GetNodesRecords(start))

}
