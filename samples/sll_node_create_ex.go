/*
	{
		"date_of_creation" => "26 Dec 2016, Monday",
		"aim_of_program"   => "To create a node for singly linked list with variable number of information",
		"coded_by"         => "hygull",
		"Go_version"	   => "1.7.1",
	}
*/
package main

import "fmt"
import "ds"
import "encoding/json"

type G struct {
	Name string
	Age  int
}

func main() {
	node := ds.CreateNode("Rishikesh", 24, 45.56, true, []int8{2, 5, 6, 7, 8, 2, 1}) //Complex data causes an error while marshalling
	fmt.Println((*node).Record["Data0"])
	fmt.Printf("%v %T\n", *node, *node)
	data, err := json.MarshalIndent(*node, "....", "\t")
	if err != nil {
		fmt.Println("Error...")
	}
	fmt.Println("data: ", string(data))

	node.Next = ds.CreateNode("Hemkesh", 22)
	fmt.Print(node.Next)

	/* Extra */
	str := ds.Unmarshall(G{"Hemkesh", 22})

	fmt.Print("\n", str)

	fmt.Println(ds.Unmarshall(1))
}
