package uicontrol

type StringOrBool interface {
	~string | ~bool
}

type UIControl[T StringOrBool] interface {
	Get() T
	Set(value T)
}
