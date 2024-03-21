package parkinglot

import (
	"errors"
	vehicle "parkmate/car"
)

var (
	ErrCarIsAlreadyParked    = errors.New("car is already parked")
	ErrCarIsAlreadyUnParked  = errors.New("car is already unparked")
	ErrNoCarParked           = errors.New("no car parked")
	ErrParkingLotTotalIsZero = errors.New("parking lot total is zero")
	ErrParkingLotIsFull      = errors.New("parking lot is full")
)

type ParkingLot struct {
	total     int
	available int
	name      string
}

func NewParkingLot(name string, total int) (*ParkingLot, error) {
	if total <= 0 {
		return nil, ErrParkingLotTotalIsZero
	}
	return &ParkingLot{
		total:     total,
		available: total,
		name:      name,
	}, nil
}
func (parkingLot *ParkingLot) GetTotal() int {
	return parkingLot.total
}

func (parkingLot *ParkingLot) GetName() string {
	return parkingLot.name
}

func (parkingLot *ParkingLot) GetAvailable() int {
	return parkingLot.available
}

func (parkingLot *ParkingLot) ParkCar(car *vehicle.Car) (bool, error) {

	if car.GetIsParked() {
		return true, ErrCarIsAlreadyParked
	}
	if parkingLot.available <= 0 {
		return false, ErrParkingLotIsFull
	}
	car.Park()
	parkingLot.available -= 1
	return true, nil
}

func (parkingLot *ParkingLot) UnParkCar(car *vehicle.Car) (bool, error) {
	if !car.GetIsParked() {
		return true, ErrCarIsAlreadyUnParked
	} else if parkingLot.available == parkingLot.total {
		return false, ErrNoCarParked
	}
	car.UnPark()
	parkingLot.available += 1
	return true, nil
}
