package dialogbox

type DialogBox interface {
	Changed(uiControl any)
	SimulateUserInteraction()
}
