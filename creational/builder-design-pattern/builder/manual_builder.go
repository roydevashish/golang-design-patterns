package builder

import (
	cartype "github.com/roydevashish/golang-design-patterns/builder/car_type"
	"github.com/roydevashish/golang-design-patterns/builder/dashboard"
	"github.com/roydevashish/golang-design-patterns/builder/engine"
	gpsnavigator "github.com/roydevashish/golang-design-patterns/builder/gps_navigator"
	"github.com/roydevashish/golang-design-patterns/builder/manual"
	"github.com/roydevashish/golang-design-patterns/builder/wheels"
)

type ManualBuilder struct {
	manual *manual.Manual
}

func NewManualBuilder() *ManualBuilder {
	manualBuilder := &ManualBuilder{}
	manualBuilder.Reset()
	return manualBuilder
}

func (m *ManualBuilder) Reset() {
	m.manual = manual.NewManual()
}

func (m *ManualBuilder) SetCarType(typeOfCar cartype.CarType) Builder {
	m.manual.TypeOfCar = typeOfCar
	return m
}

func (m *ManualBuilder) SetSeats(seats int) Builder {
	m.manual.Seats = seats
	return m
}

func (m *ManualBuilder) IsConvertible(isConvertible bool) Builder {
	m.manual.IsConvertible = isConvertible
	return m
}

func (m *ManualBuilder) SetEngine(engine *engine.Engine) Builder {
	m.manual.Engine = engine
	return m
}

func (m *ManualBuilder) SetWheels(wheels *wheels.Wheels) Builder {
	m.manual.Wheels = wheels
	return m
}

func (m *ManualBuilder) SetDashboard(dashboard *dashboard.Dashboard) Builder {
	m.manual.Dashboard = dashboard
	return m
}

func (m *ManualBuilder) SetGPSNavigator(gpsNavigator *gpsnavigator.GPSNavigator) Builder {
	m.manual.GPSNavigator = gpsNavigator
	return m
}

func (m *ManualBuilder) GetManual() *manual.Manual {
	manual := m.manual
	m.Reset()
	return manual
}
