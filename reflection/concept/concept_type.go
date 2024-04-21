package concept

import (
	"fmt"
	"reflect"
)

type TypeConcept struct {
	Name string
	Age  int
}

func TypeConceptFunc() {

	typeConcept := TypeConcept{
		Name: "Aan",
		Age:  21,
	}
	// will return type
	t := reflect.TypeOf(typeConcept)

	// akan menampilkan nama dari type nya, dalam hal ini struct nya
	fmt.Println(t.Name())
	fmt.Println(t.Kind())
	fmt.Println(t.Kind())

	// jika value nya []slice atau pointer namanya akan kosong

	p := &TypeConcept{
		Name: "Aan",
		Age:  21,
	}

	fmt.Println(reflect.TypeOf(p).Name())

	// kind

	k := reflect.TypeOf(typeConcept).Kind()

	if k == reflect.Ptr {

	}

}
