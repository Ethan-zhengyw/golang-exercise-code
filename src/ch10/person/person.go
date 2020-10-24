package person

import "sync"

type Person struct {
	Mu sync.Mutex
	firstName string
	lastName string
}

func (p *Person) SetFirstName(newFirstName string)  {
	p.firstName = newFirstName
}

func (p *Person) SetLastName(newLastName string)  {
	p.lastName = newLastName
}

func (p *Person) FirstName() string {
	return p.firstName
}

func (p *Person) LastName() string {
	return p.lastName
}
