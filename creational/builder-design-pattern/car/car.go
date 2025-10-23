package car

import (
	cartype "github.com/roydevashish/golang-design-patterns/builder/car_type"
	"github.com/roydevashish/golang-design-patterns/builder/dashboard"
	"github.com/roydevashish/golang-design-patterns/builder/engine"
	gpsnavigator "github.com/roydevashish/golang-design-patterns/builder/gps_navigator"
	"github.com/roydevashish/golang-design-patterns/builder/wheels"
)

type Car struct {
	TypeOfCar     cartype.CarType
	Seats         int
	IsConvertible bool
	Engine        *engine.Engine
	Dashboard     *dashboard.Dashboard
	Wheels        *wheels.Wheels
	GPSNavigator  *gpsnavigator.GPSNavigator
	Fule          int
}

func NewCar() *Car {
	return &Car{}
}
