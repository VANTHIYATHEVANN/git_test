package person

type Person struct {
	name string
}

func NewPerson(name string) *Person {
	return &Person{
		name: name,
	}
}
func (person *Person) GetName() string {
	return person.name
}
