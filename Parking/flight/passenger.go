package flight

import "parkmate/person"

type Passenger struct {
	person    *person.Person
	isBoarded bool
}

func NewPassenger(person *person.Person) *Passenger {
	return &Passenger{
		person:    person,
		isBoarded: false,
	}
}
func (passenger *Passenger) GetIsBoarded() bool {
	return passenger.isBoarded
}
