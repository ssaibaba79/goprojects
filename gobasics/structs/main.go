package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Person struct {
	Name string
	Age  int
	Id   int
}

type Department struct {
	Name string
}

// struct with embedded field and pointer field
type Staff struct {
	Person
	Department *Department
	DOJ        time.Time
}

type Student struct {
	Person
	Year string
}

func NewStudent(name string, age int, year string) *Student {
	id := rand.Int()
	return &Student{Person{name, age, id}, year}
}

// anonymous struct example
func Log(l struct {
	level   string
	message string
	err     error
}) {
	var msg string
	switch {
	case l.err != nil:
		msg = fmt.Sprintf("%v %s [message:%s][error:%s]", time.Now(), l.level,
			l.message,
			l.err)
	default:
		msg = fmt.Sprintf("%v %s [message:%s]", time.Now(), l.level,
			l.message)
	}

	fmt.Println(msg)
}

func main() {

	// using order and type of fields
	p := Person{"Sam", 10, 1000001}
	s := Staff{Person: Person{Name: "Rob Pike", Age: 40, Id: 1000002}, Department: &Department{"CS"}}

	st := NewStudent("Shan", 19, "Freshmen")
	fmt.Println(p, s, *st)
	Log(struct {
		level   string
		message string
		err     error
	}{level: "info", message: "test log message"})

	Log(struct {
		level   string
		message string
		err     error
	}{level: "error", message: "test err message", err: errors.New("method not allowed")})
}
