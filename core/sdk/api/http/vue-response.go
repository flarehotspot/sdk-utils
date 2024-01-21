package sdkhttp

type VueResponseWriter interface {
	JsonData(data any)
	Redirect(routename string, pairs ...string)
}
