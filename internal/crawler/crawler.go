package crawler

import (
	"fmt"
	"reflect"
	"strings"
)

type (
	Field struct {
		Type    reflect.Type
		Value   any
		Name    string
		TypeStr string
		// TODO:
		// 	Add tags supports, and other things.
	}

	Node struct {
		Field    *Field
		Children []*Node
	}
)

// Crawl is a function that will receive any value and return
// all details from that value & type.
func Crawl(v any) {
	crawlDepth(v, 0)
}

func crawlDepth(v any, depth int) {
	vType := reflect.TypeOf(v)
	if vType.Kind() != reflect.Struct {
		return
	}

	fieldNum := vType.NumField()
	if fieldNum <= 0 {
		return
	}

	vValue := reflect.ValueOf(v)
	tabs := "- " + strings.Repeat("	", depth)

	for i := 0; i < fieldNum; i++ {
		fieldValue := vValue.Field(i)
		fieldType := vType.Field(i)

		fmt.Println(tabs + fieldType.Name + ": " + fieldValue.String())

		fType := reflect.TypeOf(fieldValue)
		fValue := reflect.ValueOf(fieldValue)

		if fType.Kind() == reflect.Struct {
			crawlDepth(fieldValue.Interface(), depth+1)
		}

		_ = fType
		_ = fValue
	}

	_ = vType
	_ = vValue
}
