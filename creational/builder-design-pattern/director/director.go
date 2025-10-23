package director

import (
	"github.com/roydevashish/golang-design-patterns/builder/builder"
	cartype "github.com/roydevashish/golang-design-patterns/builder/car_type"
	"github.com/roydevashish/golang-design-patterns/builder/dashboard"
	"github.com/roydevashish/golang-design-patterns/builder/engine"
	gpsnavigator "github.com/roydevashish/golang-design-patterns/builder/gps_navigator"
	"github.com/roydevashish/golang-design-patterns/builder/wheels"
)

type Director struct{}

func NewDirector() *Director {
	return &Director{}
}

func (d *Director) ConstructSportsCar(builder builder.Builder) {
	builder.SetCarType(cartype.SPORTS).SetSeats(4).SetEngine(engine.NewEngine()).SetDashboard(dashboard.NewDashboard(true)).SetWheels(wheels.NewWheels(20)).SetGPSNavigator(gpsnavigator.NewGPSNavigator())
}

func (d *Director) ConstructSUVCar(builder builder.Builder) {
	builder.SetCarType(cartype.SUV).SetSeats(7).SetEngine(engine.NewEngine()).SetDashboard(dashboard.NewDashboard(false)).SetWheels(wheels.NewWheels(19)).SetGPSNavigator(gpsnavigator.NewGPSNavigator())
}
