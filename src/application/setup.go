package application

// Application settings encapsulated in a single component. It can be as complex as we want.
// For example, we can create another struct for database parameters and embed it here.
type AppConfig struct {
	Port int
}

func Setup(appConfig AppConfig) func() {
	logger, _ := makeLoggerForDevelopment()
	domainSetup := createDomainSetup(logger)
	configureHandlers := func() {
		setupPromotionsHandler(domainSetup.UseCase, logger)
	}
	setupMux(appConfig.Port, configureHandlers)
	return func() {
		logger.Sync()
	}
}
