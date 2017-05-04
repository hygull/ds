/*
	#created_on : 27 Dec 2016,Tuesday(afternoon)
	#coded_by   : hygull
*/
package main

import "fmt"
import "ds"

func main() {
	/* CASE 1*/
	var start2 *ds.Node = nil
	err := ds.ShowNodes(start2)
	if err != nil {
		fmt.Println(err) //singly linked list is empty
	}

	/* CASE 2 */
	var start *ds.Node
	fmt.Println("\n********* Testing (Inserting nodes at beginning & end of the linked list) **********\n")
	node := ds.CreateNode("Hygull", 24)
	start = ds.InsertSllNodeAtBeginning(start, node) //We can also use the global START variable
	/*
				+---------------+----+
		        start-->| {"Hygull",24} | nil|
				+--------------------+
	*/
	node2 := ds.CreateNode("Hemkesh", 22)
	start = ds.InsertSllNodeAtBeginning(start, node2)
	/*
				 +---------------+----+    +----------------+----+
		        start===>| {"Hemkesh",22}|  ======>| {"Hygull", 24} | nil|
				 +---------------+----+    +----------------+----+
	*/
	node3 := ds.CreateNode("Darshan", 23)
	start = ds.InsertSllNodeAtEnd(start, node3)
	/*
				 +---------------+----+    +----------------+----+     +----------------+----+
		        start===>| {"Hemkesh",22}|  ======>| {"Hygull", 24} |  =======>| {"Darshan",23} | nil|
				 +---------------+----+    +----------------+----+     +----------------+----+
	*/
	ds.ShowNodes(start)
	fmt.Println()

	ds.ShowNodesHierarchy(start)
	fmt.Println()

	recordsList := ds.GetNodesRecords(start)
	if len(recordsList) == 0 {
		fmt.Println("Singly linked list is empty")
	}
	fmt.Println()

	marshalledRecords, _ := ds.Marshall(recordsList)
	fmt.Println(marshalledRecords)

	/* Deleting nodes from the beginning of the linked list */
	start, _ = ds.DeleteSllNodeFromBeginning(start)
	fmt.Println("After deleting 1 node from beginning")
	fmt.Println(ds.GetNodesRecords(start))

	/* Inserting 1 node at end*/
	node4 := ds.CreateNode("Surendra", 24)
	start = ds.InsertSllNodeAtEnd(start, node4)
	fmt.Println("After adding 1 node at end")
	fmt.Println(ds.GetNodesRecords(start))

	/* Deleting 1 node from the end of linked list */
	err = ds.DeleteSllNodeFromEnd(start)
	if err != nil { //List is empty
		fmt.Println(err)
	}
	fmt.Println("After deleting 1 node from end")
	fmt.Println(ds.GetNodesRecords(start))

}
