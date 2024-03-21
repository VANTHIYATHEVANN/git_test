package parkinglot

import (
	vehicle "parkmate/car"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEmptyParkingLotWhereTotalToBeGivenValueAndNameToBeGiven(t *testing.T) {
	name := "Airport"
	total := 500

	parkinglot, err := NewParkingLot(name, total)

	assert.NotNil(t, parkinglot)
	assert.Nil(t, err)
	assert.Equal(t, parkinglot.GetTotal(), total)
	assert.Equal(t, parkinglot.GetAvailable(), total)
	assert.Equal(t, parkinglot.GetName(), name)
}

func TestNewEmptyParkingLotWhereTotalToBeZeroAndNameToBeGiven(t *testing.T) {
	name := "Airport"
	total := 0

	parkinglot, err := NewParkingLot(name, total)

	assert.Nil(t, parkinglot)
	assert.ErrorIs(t, ErrParkingLotTotalIsZero, err)

}

func TestParkWhereParkingLotIsFull(t *testing.T) {
	name := "Airport"
	total := 1
	parkinglot, _ := NewParkingLot(name, total)
	car1 := vehicle.NewCar()
	parkinglot.ParkCar(car1)
	car2 := vehicle.NewCar()

	result, err := parkinglot.ParkCar(car2)

	assert.ErrorIs(t, err, ErrParkingLotIsFull)
	assert.Equal(t, false, result)

}
func TestParkWhenArgumentIsCarAndCarIsNotParked(t *testing.T) {
	name := "Airport"
	total := 500
	parkinglot, _ := NewParkingLot(name, total)
	available := parkinglot.GetAvailable()
	car := vehicle.NewCar()

	result, err := parkinglot.ParkCar(car)

	assert.Equal(t, true, result)
	assert.Nil(t, err)
	assert.Equal(t, available-1, parkinglot.GetAvailable())

}

func TestParkWhenArgumentIsCarAndCarIsAlreadyParked(t *testing.T) {
	name := "Airport"
	total := 500
	parkinglot, _ := NewParkingLot(name, total)
	car := vehicle.NewCar()

	parkinglot.ParkCar(car)

	result, err := parkinglot.ParkCar(car)
	assert.Equal(t, true, result)
	assert.ErrorIs(t, ErrCarIsAlreadyParked, err)

}

func TestUnParkWhenArguementIsCarAndCarIsParked(t *testing.T) {
	name := "Airport"
	total := 500
	parkinglot, _ := NewParkingLot(name, total)
	car := vehicle.NewCar()
	parkinglot.ParkCar(car)

	result, err := parkinglot.UnParkCar(car)

	assert.Equal(t, true, result)
	assert.Nil(t, err)
}

func TestUnParkWhenArguementIsCarAndCarIsAlreadyUnParked(t *testing.T) {
	name := "Airport"
	total := 500
	parkinglot, _ := NewParkingLot(name, total)
	car1 := vehicle.NewCar()
	car2 := vehicle.NewCar()
	parkinglot.ParkCar(car1)
	parkinglot.ParkCar(car2)
	parkinglot.UnParkCar(car2)

	result, err := parkinglot.UnParkCar(car2)

	assert.Equal(t, true, result)
	assert.ErrorIs(t, ErrCarIsAlreadyUnParked, err)
}
