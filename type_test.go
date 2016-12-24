/*
	{
		"date_of_creation" => "24 Dec 2016, Friday",
		"aim_of_program"   => "To check the type of data",
		"coded_by"         => "Rishikesh Agrawani",
		"Go_version"	   => "1.7.1",
	}
*/
package ds

import "testing"
import "fmt"

func TestType(t *testing.T) {
	type Person struct {
		Name  string
		Age   int8
		Hobby string
	}
	person := Person{"Rishikesh", 24, "Programming"}

	fmt.Println("TypeOf(89)", Type(89))
	fmt.Println("TypeOf(int32(13))", Type(int32(13)))
	fmt.Println("TypeOf(int64(239))", Type(int64(239)))
	fmt.Println("TypeOf(float32(5.6))", Type(float32(5.6)))
	fmt.Println("TypeOf(34.56)", Type(34.56))
	fmt.Println(`TypeOf(Person{"Rishikesh",24,"Programming"})`, Type(person))
}

/*

TypeOf(89) int
TypeOf(int32(13)) int32
TypeOf(int64(239)) int64
TypeOf(float32(5.6)) float32
TypeOf(34.56) float64
TypeOf(Person{"Rishikesh",24,"Programming"}) ds.Person

*/
