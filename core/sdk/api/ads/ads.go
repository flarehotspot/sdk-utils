package sdkads

// IAdsApi is used for displaying ads in the captive portal.
type IAdsApi interface {
  // Init initializes the ads API with the given app ID.
	Init(appId string)
}
