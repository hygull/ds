/*
	{
		"date_of_creation" => "26 Dec 2016, Monday(morning)",
		"aim_of_program"   => "To test node creation function of singly linked list",
		"coded_by"         => "Rishikesh Agrawani",
		"Go_version"	   => "1.7.1",
	}
*/
package ds

import "testing"
import "reflect"
import "fmt"

func TestCreateNode(t *testing.T) {

	/* Case1 */
	node := CreateNode("Rishikesh", 24, 45.56, true, []int8{2, 5, 6, 7, 8, 2, 1}) //Complex data causes an error while marshalling
	fmt.Println((*node).Record["Data0"])

	d := reflect.ValueOf(node).Elem() // user User{"Rishikesh"}; reflect.ValueOf(&user).Elem()

	for i := 0; i < d.NumField(); i++ {
		fmt.Println(d.Type().Field(i).Name, " ==> ", d.Field(i).Type(), " ==> ", d.Field(i).Interface())
	}
	fmt.Println()

	/* Case2 */
	node2 := CreateNode([]string{"Bangalore", "Delhi", "Raipur"}, []float64{78.98, 65.43, 653.76}, []int8{2, 5, 6, 7, 8, 2, 1}) //Complex data causes an error while marshalling
	fmt.Println((*node2).Record["Data1"])

	d1 := reflect.ValueOf(node2).Elem() // user User{"Rishikesh"}; reflect.ValueOf(&user).Elem()

	for i := 0; i < d1.NumField(); i++ {
		fmt.Println(d1.Type().Field(i).Name, " ==> ", d1.Field(i).Type(), " ==> ", d1.Field(i).Interface())
	}
}
