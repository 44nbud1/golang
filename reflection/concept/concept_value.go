package concept

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string `name:"name" tag2:"tag2"`
	Age  int
}

func ValuesConcept() {

	fmt.Println("===\t\t Value concept \t\t===")
	s := Student{
		Name: "Aan budi",
		Age:  21,
	}

	// string
	greet := "greeting"

	gRValue := reflect.ValueOf(greet)
	fmt.Println(gRValue.Kind())

	// nilainya harus pointer supaya bisa diubah
	gPRValue := reflect.ValueOf(&greet)
	if gPRValue.Kind() == reflect.Ptr {
		fmt.Println("Type : ", gRValue)
		gPRValue.Elem().SetString("Hello")
	}

	// tampilkan nilai baru
	fmt.Println(greet)

	// struct val
	vs := reflect.ValueOf(s)
	fmt.Println(vs.NumField())
	fmt.Println(vs.NumMethod())

	for i := 0; i < vs.NumField(); i++ {
		fmt.Println("=====")
		fmt.Println("Tags ", vs.Type().Field(i).Tag)
		fmt.Println("Type ", vs.Type().Field(i).Type)
		fmt.Println("Name ", vs.Type().Field(i).Name)
	}

	tv := reflect.ValueOf(&s)

	if tv.Kind() == reflect.Ptr {

		// set elem
		tv = tv.Elem()
		for i := 0; i < tv.NumField(); i++ {
			if tv.Kind() == reflect.Struct {
				if tv.Type().Field(i).Name == "Name" {
					if tv.Field(i).Kind() == reflect.String {
						tv.Field(i).SetString("change name")
					}
				}
			}
		}
	}
	fmt.Println("===\t\t Value concept \t\t===")

	fmt.Println(s)

}
