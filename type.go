package ds

import "reflect"

//A function that checks the type of received data & returns its type
func Type(data interface{}) reflect.Type {
	return reflect.TypeOf(data)
}
