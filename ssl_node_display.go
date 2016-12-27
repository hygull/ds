package ds

import "errors"
import "fmt"

//A fuunction that displays all the nodes hirarchy by marshalling the complete list, if list is empty,
//or if there's a any problem in marshalling, then throws an error
func ShowNodesHierarchy(__startNodePtr *Node) error {
	if __startNodePtr == nil {
		return errors.New("singly linked list is empty")
	}

	listAsSlice, err := Marshall(*__startNodePtr)
	if err != nil {
		return err
	}
	fmt.Println()
	fmt.Println(string(listAsSlice))
	return nil
}

//A fuunction that displays all the nodes by marshalling it, if list is empty,
//or if there's a any problem in marshalling, then throws an error
func ShowNodes(__startNodePtr *Node) error {
	if __startNodePtr == nil {
		return errors.New("singly linked list is empty")
	}
	fmt.Println()
	i := 1
	for __startNodePtr != nil {
		nodeRecordAsSlice, err := Marshall((*__startNodePtr).Record)
		if err != nil {
			return err
		}
		fmt.Println("--------------------- Record at Node" + fmt.Sprint(i) + " -----------------------------")
		fmt.Println(string(nodeRecordAsSlice))
		fmt.Println()
		__startNodePtr = __startNodePtr.Next
		i++
	}
	return nil
}
