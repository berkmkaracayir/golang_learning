//Map verileri anahtar-değer olarak tutar, değerleri anahtarlarına göre tutar ve getirir. 

package main

import "fmt"

func main() {
	fmt.Println("Define a map")
	m := make(map[string]int)	 			// create an empty map
	m["one"] = 1              				// set the value of the key "one" to 1
	m["two"] = 2             		 		// set the value of the key "two" to 2
	fmt.Println(m["one"])
	
	grocery_list := make(map[int]string)	// create an empty map
	grocery_list[0] = "apple"              // set the value of the key 0 to "apple" 
	grocery_list[1] = "banana"             // set the value of the key "two" to "banana" 
	fmt.Println(grocery_list[0])

	myMap := map[string]int{"a": 1, "b": 2, "c": 3}
	delete(myMap, "b") 						// Delete the key "b"
	fmt.Println(myMap) 						// Output: map[a:1 c:3]
	
}