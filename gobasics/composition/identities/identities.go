package identities

import (
	"math/rand"
	"strconv"
)

type Pid string

func newPid() Pid {
	return Pid("U-" + strconv.FormatUint(rand.Uint64(), 10))
}

var departments = map[string]*Department{}

func init() {
	departments["cs"] = NewDepartment("cs", "ComputerScience")
	departments["math"] = NewDepartment("math", "Mathematics")
	departments["admin"] = NewDepartment("admin", "Administration")
}

type Department struct {
	Id   string
	Name string
	Code string
}

func GetDepartment(code string) *Department {
	return departments[code]
}

func NewDepartment(code string, name string) *Department {
	id := "D-" + strconv.FormatUint(rand.Uint64(), 10)
	return &Department{id, name, code}
}

type Email string

type Phone string

type Person struct {
	id        Pid
	Name      string
	Email     Email
	ContactNo Phone
}

func (p *Person) GetId() Pid {
	return p.id
}

func NewPerson(name string, email string, contactNo string) *Person {
	return &Person{id: newPid(), Name: name, Email: Email(email), ContactNo: Phone(contactNo)}
}

type Staff struct {
	Person
	Department *Department
}

func NewStaff(name string, email string, contactNo string, dept string) *Staff {

	return &Staff{*NewPerson(name, email, contactNo), GetDepartment(dept)}

}
