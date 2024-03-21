package vehicle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCarWhereIsParkedToBeFalse(t *testing.T) {
	car := NewCar()
	assert.NotNil(t, car)
	isParked := false
	assert.Equal(t, car.GetIsParked(), isParked)

}
func TestParkWhereIsParkedToBeTrue(t *testing.T) {
	car := NewCar()
	car.Park()
	assert.Equal(t, car.GetIsParked(), true)
}

func TestUnParkWhereIsParkedToBeTrue(t *testing.T) {
	car := NewCar()

	car.Park()
	car.UnPark()
	assert.Equal(t, car.GetIsParked(), false)
}
