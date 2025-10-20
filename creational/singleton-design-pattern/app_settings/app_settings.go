package appsettings

var instance *appSettings

type appSettings struct {
	settings map[string]string
}

func GetInstance() *appSettings {
	if instance == nil {
		instance = newAppSettings()
		return instance
	}

	return instance
}

func newAppSettings() *appSettings {
	return &appSettings{
		settings: make(map[string]string),
	}
}

func (a *appSettings) Get(key string) string {
	val, found := a.settings[key]
	if !found {
		return ""
	}

	return val
}

func (a *appSettings) Set(key, value string) {
	a.settings[key] = value
}
