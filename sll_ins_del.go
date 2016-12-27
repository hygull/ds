/*
	{
		"date_of_creation" => "27 Dec 2016, Tuesday(afternoon)",
		"aim_of_program"   => "Performing operations on the linked list",
		"coded_by"         => "hygull",
		"Go_version"	   => "1.7.1",
	}
*/
package ds

//To reduce the number of parameters(These will be useful in other functions)
var START, END *Node

//A function that inserts a node at the beginning of the singly linked list
func InsertSllNodeAtBeginning(node *Node) *Node {
	if node == nil { //If node is not pointing to any Node(just nil)
		return nil
	}

	if START == nil {
		START = node
		END = START
	} else {
		node.Next = START
		START = node
	}
	return START //Not required (It is useful, where we wanna use this current function as an expression)
}

//A function that inserts a node at the end of the singly linked list
func InsertSllNodeAtEnd(node *Node) *Node {
	if node == nil { //If node is not pointing to any Node(just nil)
		return nil
	}

	if START == nil {
		START = node
		END = START
	} else {
		END.Next = node
	}
	return START //Not required (It is useful, where we wanna use this current function as an expression)
}
