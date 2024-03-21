package flight

import (
	"parkmate/person"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFlightWherePassengerListToBeEmpty(t *testing.T) {
	flight := NewFlight()
	assert.Equal(t, len(flight.GetPassengerList()), 0)
}

func TestBoardingOnePasssengerInFlight(t *testing.T) {
	flight := NewFlight()
	person := person.NewPerson("Aaditya")
	passenger := NewPassenger(person)
	flight.Board(passenger)
	passengerList := flight.GetPassengerList()
	currentBoardedPassengerIndex := len(passengerList) - 1
	currentBoarderPassenger := passengerList[currentBoardedPassengerIndex]
	assert.Equal(t, currentBoarderPassenger, passenger)
}
