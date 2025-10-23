package manual

import (
	"fmt"

	cartype "github.com/roydevashish/golang-design-patterns/builder/car_type"
	"github.com/roydevashish/golang-design-patterns/builder/dashboard"
	"github.com/roydevashish/golang-design-patterns/builder/engine"
	gpsnavigator "github.com/roydevashish/golang-design-patterns/builder/gps_navigator"
	"github.com/roydevashish/golang-design-patterns/builder/wheels"
)

type Manual struct {
	TypeOfCar     cartype.CarType
	Seats         int
	IsConvertible bool
	Engine        *engine.Engine
	Dashboard     *dashboard.Dashboard
	Wheels        *wheels.Wheels
	GPSNavigator  *gpsnavigator.GPSNavigator
}

func NewManual() *Manual {
	return &Manual{}
}

func (m *Manual) Print() {
	fmt.Println("car type:", m.TypeOfCar)
	fmt.Println("seats:", m.Seats)
	fmt.Println("convertible:", m.IsConvertible)
	fmt.Println("engine:", m.Engine)
	fmt.Println("dashboard:", m.Dashboard)
	fmt.Println("wheels:", m.Wheels)
	fmt.Println("gps navigator", m.GPSNavigator)
}
