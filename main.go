package main

import (
	"fmt"
	"github.com/elliotchance/orderedmap"
	"reflect"
)

func main(){
	m := orderedmap.NewOrderedMap()

	m.Set("name", "Nnamdi")
	m.Set("Price", 1.23)
	m.Set(123, true)

	m.Delete("Price")

	for _, key := range m.Keys() {
		value, _:= m.Get(key)
		fmt.Println(key, value)
	}

	// Iterate through all elements from oldest to newest:
	for el := m.Front(); el != nil; el = el.Next() {
		fmt.Println(el.Key, el.Value)
	}
	fmt.Printf("The type of m is: %v", reflect.TypeOf(m))
}