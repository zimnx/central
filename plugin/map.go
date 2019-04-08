package myplugin

import (
	"fmt"
)

type MyMap struct {
	inner map[string]string
}

func NewMyMap() *MyMap {
	return &MyMap{
		inner: map[string]string{},
	}
}

func (m *MyMap) Add(k string, v string) {
	fmt.Println(k, "=", v)
	m.inner[k] = v
}



