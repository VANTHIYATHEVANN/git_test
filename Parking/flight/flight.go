package flight

type Flight struct {
	passengerList []*Passenger
}

func NewFlight() *Flight {
	flight := Flight{
		passengerList: []*Passenger{},
	}
	return &flight
}

func (flight *Flight) GetPassengerList() []*Passenger {
	return flight.passengerList
}

func (flight *Flight) Board(passenger *Passenger) {
	passenger.isBoarded = true
	flight.passengerList = append(flight.passengerList, passenger)
}
