package builder

import (
	"github.com/roydevashish/golang-design-patterns/builder/car"
	cartype "github.com/roydevashish/golang-design-patterns/builder/car_type"
	"github.com/roydevashish/golang-design-patterns/builder/dashboard"
	"github.com/roydevashish/golang-design-patterns/builder/engine"
	gpsnavigator "github.com/roydevashish/golang-design-patterns/builder/gps_navigator"
	"github.com/roydevashish/golang-design-patterns/builder/wheels"
)

type CarBuilder struct {
	car *car.Car
}

func NewCarBuilder() *CarBuilder {
	carBuilder := &CarBuilder{}
	carBuilder.Reset()
	return carBuilder
}

func (c *CarBuilder) Reset() {
	c.car = car.NewCar()
}

func (c *CarBuilder) SetCarType(typeOfCar cartype.CarType) Builder {
	c.car.TypeOfCar = typeOfCar
	return c
}

func (c *CarBuilder) SetSeats(seats int) Builder {
	c.car.Seats = seats
	return c
}

func (c *CarBuilder) IsConvertible(isConvertible bool) Builder {
	c.car.IsConvertible = isConvertible
	return c
}

func (c *CarBuilder) SetEngine(engine *engine.Engine) Builder {
	c.car.Engine = engine
	return c
}

func (c *CarBuilder) SetWheels(wheels *wheels.Wheels) Builder {
	c.car.Wheels = wheels
	return c
}

func (c *CarBuilder) SetDashboard(dashboard *dashboard.Dashboard) Builder {
	c.car.Dashboard = dashboard
	return c
}

func (c *CarBuilder) SetGPSNavigator(gpsNavigator *gpsnavigator.GPSNavigator) Builder {
	c.car.GPSNavigator = gpsNavigator
	return c
}

func (c *CarBuilder) GetCar() *car.Car {
	car := c.car
	c.Reset()
	return car
}
