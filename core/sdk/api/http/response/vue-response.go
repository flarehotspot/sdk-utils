package response

type IVueResponse interface {
	VueComponent(componentPath string, data any)
	VueRedirect(routename string)
}
