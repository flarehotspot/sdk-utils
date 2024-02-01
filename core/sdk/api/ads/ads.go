package sdkads

// AdsApi is used for displaying ads in the captive portal.
type AdsApi interface {
  // Init initializes the ads API with the given app ID.
	Init(appId string)
}
