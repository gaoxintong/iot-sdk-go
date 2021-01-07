package storage

import (
	"fmt"
	"testing"
)

var localStorage = LocalStorage{}

func TestGet(t *testing.T) {
	a, err := localStorage.Get("a")
	if err != nil {
		panic(err)
	}
	fmt.Println(a)
}

func TestSet(t *testing.T) {
	err := localStorage.Set("c", 123)
	if err != nil {
		fmt.Println(err)
	}

	c, err := localStorage.Get("c")
	if err != nil || c != 123 {
		panic("c not equal to 123")
	}
}

func TestDel(t *testing.T) {
	err := localStorage.Del("b")
	if err != nil {
		panic(err)
	}

	b, err := localStorage.Get("b")
	if err != nil {
		panic(err)
	}
	if b != nil {
		panic("b not equal to nil")
	}
}

func TestOverride(t *testing.T) {
	err := localStorage.Set("c", 123)
	if err != nil {
		fmt.Println(err)
	}
	err = localStorage.Set("c", 333)
	if err != nil {
		fmt.Println(err)
	}

	c, err := localStorage.Get("c")
	if err != nil || c != 333 {
		panic("c not equal to 123")
	}
}
