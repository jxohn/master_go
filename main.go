package main

import (
	"fmt"

	"github.com/jxohn/master_go/reflect"
)

func main() {
	arrayStruct := &reflect.ArrayStruct{
		What:        1,
		IntSlice:    []int{3, 2, 1},
		Int32Slice:  []int32{2, 1, 3},
		StringSlice: []string{"hi", "hello"},
	}

	fmt.Printf("%+v\n", arrayStruct)
	reflect.SortArrayFields(arrayStruct)
	fmt.Printf("%+v\n", arrayStruct)
}
