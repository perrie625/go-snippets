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

type foo struct {
	Name    string
	private string
}

func (f *foo) speak(word string) {
	fmt.Println(word)
}

func (f *foo) sayName() {
	if f.Name == "" {
		fmt.Println("I have no name.")
		return
	}
	fmt.Println(f.Name)
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

func setAttr(object interface{}, field string, value string) bool {
	objectValue := reflect.ValueOf(object)
	if objectValue.Kind() != reflect.Ptr {
		return false
	}
	objectValue = objectValue.Elem()
	if !objectValue.CanSet() {
		return false
	}
	fieldValue := objectValue.FieldByName(field)
	fieldValue.SetString(value)
	return true
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	registeType(foo{})
	obj, err := newObject("foo")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reflect.TypeOf(obj))
	sample, ok := obj.(*foo)
	fmt.Println(ok)
	sample.speak("Hello, World.")

	fmt.Printf("set successful ? %t\n", setAttr(*sample, "Name", "foo"))
	sample.sayName()

	fmt.Printf("set successful ? %t\n", setAttr(sample, "Name", "foo"))
	sample.sayName()

	setAttr(sample, "private", "hahaha")
}
