package basics

import (
	"fmt"
	"maps"
)

func main() {
	//  var mapVariable map[keyType]ValueType

	// mapVariable = make(map[string]int)

	// using a Map literal
	// mapVariable = map[keyType]valueType{
	//     "key1": value1,
	//     "key2": value2,
	// }

	myMap := make(map[string]int)

	// Only the comparable type can be used as a key type in a map. The comparable types are boolean, numeric, string, pointer, channel, and interface types. The array and struct types are comparable if all their fields are comparable. The slice, map, and function types are not comparable and cannot be used as a key type in a map.

	fmt.Println(myMap)

	myMap["key1"] = 9
	myMap["code"] = 18

	fmt.Println(myMap)
	fmt.Println(myMap["key1"])

	fmt.Println(myMap["nonexistentKey"]) // This will return the zero value for the value type, which is 0 for int.

	delete(myMap, "key1")
	fmt.Println(myMap)
	myMap["key1"] = 9
	myMap["key2"] = 18
	myMap["key3"] = 27
	// clear(myMap)
	// fmt.Println(myMap)

	value, ok := myMap["key1"]
	fmt.Println(value)
	fmt.Println(ok)

	_, ok2 := myMap["key1"]
	fmt.Println("ok2:", ok2)

	// The second return value of a map lookup is a boolean that indicates whether the key was found in the map. Sometimes we just want to check if the key exists. If the key was found, the boolean will be true; otherwise, it will be false. But somehow when print by getter myMap["key1"], it will return only the value, and the second return value is ignored by compiler.

	myMap2 := map[string]int{"a": 1, "b": 2, "c": 3}
	myMap3 := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println("myMap2:", myMap2)
	fmt.Println("myMap3:", myMap3)

	if maps.Equal(myMap3, myMap2) { // Maps are reference types
		fmt.Println("myMap3 and myMap2 are equal")
	} else {
		fmt.Println("myMap3 and myMap2 are not equal")
	}

	for k, v := range myMap2 {
		fmt.Println("Key:", k, "Value:", v)
	}

	var myMap4 map[string]int // nil map, non existence, but the type is defined, not yet allocated
	fmt.Println("myMap4:", myMap4)
	if myMap4 == nil {
		fmt.Println("myMap4 is nil")
		fmt.Println("myMap4 before make:", len(myMap4))
	}
	// myMap4["key"] = 1 // This will cause a runtime panic because you cannot assign a value to a nil map. Once again, because the maps still not exist, you cannot assign a value.
	myMap4 = make(map[string]int) // Now the map is exist. Already allocated.

	fmt.Println("myMap4 after make:", len(myMap4))

	myMap5 := make(map[string]map[string]int) // Nested map, map of map
	myMap5["map1"] = myMap4
	fmt.Println(myMap5)
}
