package testImport

import "fmt"

var ToBeUsed string = "I am the imported variable"

func ToBeUsedFn() {
	fmt.Println(ToBeUsed)
}

func NewImportedFn() {
	fmt.Println("I am new")
}
