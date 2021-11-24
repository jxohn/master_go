package reflect

import (
	"reflect"
	"sort"
)

type ArrayStruct struct {
	What        int
	Wha         []ArrayStruct
	IntSlice    []int
	Int32Slice  []int32
	StringSlice []string
}

func SortArrayFields(input interface{}) {
	value := reflect.ValueOf(input).Elem()
	for x := 0; x < value.NumField(); x++ {
		switch value.Field(x).Kind() {
		case reflect.Slice:
			sliceValue := value.Field(x).Interface()
			switch value.Field(x).Type().Elem().Kind() {
			case reflect.Int:
				tempVar := sliceValue.([]int)
				sort.Slice(tempVar, func(i, j int) bool { return tempVar[i] < tempVar[j] })
			case reflect.Int8:
				tempVar := sliceValue.([]int8)
				sort.Slice(tempVar, func(i, j int) bool { return tempVar[i] < tempVar[j] })
			case reflect.Int16:
				tempVar := sliceValue.([]int16)
				sort.Slice(tempVar, func(i, j int) bool { return tempVar[i] < tempVar[j] })
			case reflect.Int32:
				tempVar := sliceValue.([]int32)
				sort.Slice(tempVar, func(i, j int) bool { return tempVar[i] < tempVar[j] })
			case reflect.String:
				tempVar := sliceValue.([]string)
				sort.Slice(tempVar, func(i, j int) bool { return tempVar[i] < tempVar[j] })
			case reflect.Int64:
				tempVar := sliceValue.([]int64)
				sort.Slice(tempVar, func(i, j int) bool { return tempVar[i] < tempVar[j] })
			case reflect.Float32:
				tempVar := sliceValue.([]float32)
				sort.Slice(tempVar, func(i, j int) bool { return tempVar[i] < tempVar[j] })
			case reflect.Float64:
				tempVar := sliceValue.([]float64)
				sort.Slice(tempVar, func(i, j int) bool { return tempVar[i] < tempVar[j] })
			default:
			}
		default:

		}
	}
}
