package main

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)

var (
	typeMap = make(map[string]reflect.Type)
)

type foo struct{}

func (f *foo) speak(word string) {
	fmt.Println(word)
}

func registeType(object interface{}) {
	value := reflect.ValueOf(object)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	typeMap[value.Type().Name()] = value.Type()
}

func newObject(name string) (interface{}, error) {
	if t, ok := typeMap[name]; ok {
		return reflect.New(t).Interface(), nil
	}
	return nil, errors.New("no this type")
}

func main() {
	registeType(foo{})
	obj, err := newObject("foo")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reflect.TypeOf(obj))
	sample, ok := obj.(*foo)
	fmt.Println(ok)
	sample.speak("Hello, World.")
}
