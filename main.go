package main

import (
	"fmt"
	"github.com/elliotchance/orderedmap"
)

func main(){
	m := orderedmap.NewOrderedMap()
	m.Set("name", "Nnamdi")
	m.Set("Price", 1.23)
	m.Set(123, true)

	m.Delete("Price")

	myMap := orderedmap.NewOrderedMap()
	mySlice := []int{1, 2, 3, 1, 4}
	for _, v := range mySlice {
		myMap.Set(v,1)
	}
	//fmt.Println(myMap)

	//for i, key := range m.Keys() {
	//	value, _:= m.Get(key)
	//	fmt.Println(key, value)
	//	fmt.Println(i,  "here")
	//}

	// Iterate through all elements from oldest to newest:
	for el := myMap.Front(); el != nil; el = el.Next() {
		fmt.Println(el.Key, el.Value)
	}
	//fmt.Printf("The type is: %T \n", myMap)
}