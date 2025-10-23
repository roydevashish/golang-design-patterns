package dashboard

type Dashboard struct {
	hasRevCounter bool
}

func NewDashboard(hasRevCounter bool) *Dashboard {
	return &Dashboard{
		hasRevCounter: hasRevCounter,
	}
}
