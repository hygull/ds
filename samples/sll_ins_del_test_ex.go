/*
	#created_on : 27 Dec 2016,Tuesday(afternoon)
	#coded_by   : hygull
*/
package main

import "fmt"
import "ds"

func main() {
	fmt.Println("\n********* Testing (Inserting nodes at beginning & end of the linked list) **********\n")
	node := ds.CreateNode("Hygull", 24)
	start := ds.InsertSllNodeAtBeginning(node) //We can also use the global START variable
	/*
				+---------------+----+
		start-->| {"Hygull",24} | nil|
				+--------------------+
	*/
	node2 := ds.CreateNode("Hemkesh", 22)
	start = ds.InsertSllNodeAtBeginning(node2)
	/*
					+---------------+----+    +----------------+----+
		   start===>| {"Hemkesh",22}|  ======>| {"Hygull", 24} | nil|
					+---------------+----+    +----------------+----+
	*/
	node3 := ds.CreateNode("Darshan", 23)
	start = ds.InsertSllNodeAtEnd(node3)
	/*
					+---------------+----+    +----------------+----+     +----------------+----+
		   start===>| {"Hemkesh",22}|  ======>| {"Hygull", 24} |  =======>| {"Darshan",23} | nil|
					+---------------+----+    +----------------+----+     +----------------+----+
	*/
	ds.ShowNodes(start)
	fmt.Println()

	ds.ShowNodesHierarchy(ds.START) //using global START variable
	fmt.Println()

	recordsList := ds.GetNodesRecords(start)
	if len(recordsList) == 0 {
		fmt.Println("Singly linked list is empty")
	}
	fmt.Println()

	marshalledRecords, _ := ds.Marshall(recordsList)
	fmt.Println(marshalledRecords)

	start = nil
	err := ds.ShowNodes(start)
	if err != nil {
		fmt.Println(err) //singly linked list is empty
	}

}
