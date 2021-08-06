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
	for i, v := range mySlice {
		myMap.Set(v,i*10)
	}
	//fmt.Println(myMap)

	//for i, key := range m.Keys() {
	//	value, _:= m.Get(key)
	//	fmt.Println(key, value)
	//	fmt.Println(i,  "here")
	//}

	//myMap.Get(mySlice[0])

	for _, k := range mySlice {
		val, ok := myMap.Get(k)
		if !ok {
			fmt.Println("this key does not exist", k)
		}
		fmt.Println(val)
	}

	// Iterate through all elements from oldest to newest:
	//for el := myMap.Front(); el != nil; el = el.Next() {
	//	fmt.Println(el.Key, el.Value)
	//}
	//fmt.Printf("The type is: %T \n", myMap)
}