package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Data[T any] struct {
	First  T
	Second T
}

func (d *Data[_]) SayHello(name string) string {
	return "Hello " + name
}

func (d *Data[T]) ChangeFirst(first T) T {
	d.First = first
	return d.First
}

func TestData(t *testing.T) {
	data := Data[string]{
		First:  "Luthfi",
		Second: "Hanif",
	}
	fmt.Println(data)
}

func TestGenericMethod(t *testing.T) {
	data := Data[string]{
		First:  "Luthfi",
		Second: "Hanif",
	}
	assert.Equal(t, "Hanif", data.ChangeFirst("Hanif"))
	assert.Equal(t, "Hello Luthfi", data.SayHello("Luthfi"))
}
