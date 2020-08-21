package service

// Init system component.
func (a *Application) initCore() {
	// Init Config
	a.config.Init(a.GetAttribute("ConfigPath"))

	// Init logger.
	pattern := ""
	if pattern = a.GetConfig("server.error.pattern").ToString(); pattern == "" {
		pattern = "local"
	}
	a.logger.Init(pattern, a.GetAttribute("RuntimeLogPath"))
}