package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name    string
	Age     int
	Address string
}

func inspectStruct(person interface{}) {
	valueOf := reflect.ValueOf(person)
	typeOf := reflect.TypeOf(person)

	fmt.Printf("Type: %v\n", typeOf)

	if valueOf.Kind() == reflect.Struct {
		fmt.Printf("Fields:\n")
		for i := 0; i < valueOf.NumField(); i++ {
			field := valueOf.Field(i)
			fieldName := typeOf.Field(i).Name
			fieldType := field.Type()
			fieldValue := field.Interface()

			fmt.Printf("%s (%v): %v\n", fieldName, fieldType, fieldValue)
		}
	}
}

func setField(person interface{}, fieldName string, newValue interface{}) {
	valueOf := reflect.ValueOf(person).Elem()

	field := valueOf.FieldByName(fieldName)

	if field.IsValid() && field.CanSet() {
		newFieldValue := reflect.ValueOf(newValue)
		if newFieldValue.Type() == field.Type() {
			field.Set(newFieldValue)
		} else {
			fmt.Printf("Error: Type mismatch for field %s\n", fieldName)
		}
	} else {
		fmt.Printf("Error: Field %s not found or cannot be set\n", fieldName)
	}
}

func main() {
	person := Person{
		Name:    "Joan Clarke",
		Age:     30,
		Address: "123 Main St",
	}

	inspectStruct(person)

	setField(&person, "Age", 35)
	inspectStruct(person)
}