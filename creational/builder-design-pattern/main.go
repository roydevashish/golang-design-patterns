package main

import (
	"github.com/roydevashish/golang-design-patterns/builder/builder"
	"github.com/roydevashish/golang-design-patterns/builder/director"
)

func main() {
	// cb := builder.NewCarBuilder()
	// cb.SetCarType(cartype.SPORTS).SetSeats(4).SetEngine(engine.NewEngine()).SetDashboard(dashboard.NewDashboard(true)).SetWheels(wheels.NewWheels(20)).SetGPSNavigator(gpsnavigator.NewGPSNavigator())
	// sportsCar := cb.GetCar()
	// sportsCar.Fule = 100

	// mb := builder.NewManualBuilder()
	// mb.SetCarType(cartype.SPORTS).SetSeats(4).SetEngine(engine.NewEngine()).SetDashboard(dashboard.NewDashboard(true)).SetWheels(wheels.NewWheels(20)).SetGPSNavigator(gpsnavigator.NewGPSNavigator())
	// sportsCarManual := mb.GetManual()
	// sportsCarManual.Print()

	cb := builder.NewCarBuilder()
	mb := builder.NewManualBuilder()
	d := director.NewDirector()

	d.ConstructSportsCar(cb)
	sportsCar := cb.GetCar()
	sportsCar.Fule = 100

	d.ConstructSUVCar(cb)
	suvCar := cb.GetCar()
	suvCar.Fule = 80

	d.ConstructSportsCar(mb)
	sportsCarManual := mb.GetManual()
	sportsCarManual.Print()

	d.ConstructSUVCar(mb)
	suvCarManual := mb.GetManual()
	suvCarManual.Print()
}
