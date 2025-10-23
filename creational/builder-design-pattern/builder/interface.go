package builder

import (
	cartype "github.com/roydevashish/golang-design-patterns/builder/car_type"
	"github.com/roydevashish/golang-design-patterns/builder/dashboard"
	"github.com/roydevashish/golang-design-patterns/builder/engine"
	gpsnavigator "github.com/roydevashish/golang-design-patterns/builder/gps_navigator"
	"github.com/roydevashish/golang-design-patterns/builder/wheels"
)

type Builder interface {
	Reset()
	SetCarType(typeOfCar cartype.CarType) Builder
	SetSeats(seats int) Builder
	IsConvertible(isConvertible bool) Builder
	SetEngine(engine *engine.Engine) Builder
	SetWheels(wheels *wheels.Wheels) Builder
	SetDashboard(dashboard *dashboard.Dashboard) Builder
	SetGPSNavigator(gpsNavigator *gpsnavigator.GPSNavigator) Builder
}
