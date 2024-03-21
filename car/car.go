package vehicle

type Car struct {
	isParked bool
}

func NewCar() *Car {
	return &Car{
		isParked: false,
	}
}

func (car *Car) Park() {

	car.isParked = true
}

func (car *Car) UnPark() {
	car.isParked = false
}

func (car *Car) GetIsParked() bool {
	return car.isParked
}
