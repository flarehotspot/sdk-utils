package sdkhttp

type VueRoutePath interface {
    GetTemplate() string
    URL(pairs ...string) string
}
