package plugins

type AdsApi struct {
	plugin *PluginApi
}

func (ads *AdsApi) Init(appId string) {

}

func NewAdsApi(plugin *PluginApi) *AdsApi {
	return &AdsApi{plugin}
}
