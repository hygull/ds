/*
	{
		"date_of_creation" => "27 Dec 2016, Tuesday(afternoon)",
		"aim_of_program"   => "Performing operations on the linked list",
		"coded_by"         => "hygull",
		"Go_version"	   => "1.7.1",
	}
*/
package ds

import "errors"

//A function that inserts a node at the beginning of the singly linked list
func InsertSllNodeAtBeginning(start *Node, node *Node) *Node {

	if node == nil { //If node is not pointing to any Node(just nil)
		return nil
	}
	if start == nil { //If list is empty
		start = node
	} else {
		node.Next = start
		start = node
	}

	return start
}

//A function that inserts a node at the end of the singly linked list
func InsertSllNodeAtEnd(start *Node, node *Node) *Node {
	if node == nil { //If node is not pointing to any Node(just nil)
		return nil
	}
	if start == nil {
		start = node //If list is empty
	} else {
		temp := start
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = node
	}

	return start
}

//A function that deletes a node from beginning
func DeleteSllNodeFromBeginning(start *Node) (*Node, error) {
	if start != nil { //If tree is empty
		var node *Node
		node = start
		start = start.Next
		node.Next = nil
	} else {
		return start, errors.New("singly linked is empty, there is no more to be deleted\n")
	}
	return start, nil
}

//A function that deletes a node from end
func DeleteSllNodeFromEnd(start *Node) error { //Only return error(As node is being inserted at end, so start wo'nt change)
	if start != nil { //If tree is empty
		if start.Next == nil { //Only 1 node is there
			start = nil
		} else { //More than 2 nodes are there
			var node *Node = start
			for node.Next.Next != nil {
				node = node.Next
			}
			node.Next = nil
		}
	} else {
		return errors.New("singly linked is empty, there is no more to be deleted\n")
	}
	return nil
}
