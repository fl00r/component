package component

import (
	"fmt"
	"testing"
)

type SomeComponent struct {
	number int
}

func NewSomeComponent(numbers ...interface{}) Lifecycle {
	c := SomeComponent{numbers[0].(int)}
	return &c
}

func (comp *SomeComponent) Start(components ...interface{}) Lifecycle {
	for _, c := range components {
		comp.number += c.(*SomeComponent).number
	}
	fmt.Println(comp.number)
	return comp
}

func (comp *SomeComponent) Stop() Lifecycle {
	comp.number = 0
	return comp
}

func TestSystem(t *testing.T) {
	system := NewSystem()
	system.
		NewComponent("component-1").
		Constructor(NewSomeComponent).
		Args(1)
	system.
		NewComponent("component-2").
		Constructor(NewSomeComponent).
		Args(2).
		Dependencies("component-1")

	component3 :=
		system.
			NewComponent("component-3").
			Constructor(NewSomeComponent).
			Args(3).
			Dependencies("component-1", "component-2")
	err := system.Start()
	if err != nil {
		panic(err)
	}
	exp := 7
	n := component3.entity.(*SomeComponent).number
	if exp != n {
		t.Errorf("%d != %d", exp, n)
		t.Fail()
	}

	system.Stop()
	exp = 0
	n = component3.entity.(*SomeComponent).number
	if exp != n {
		t.Errorf("%d != %d", exp, n)
		t.Fail()
	}
}
