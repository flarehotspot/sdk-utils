package views

type ILayoutData interface {
	Helpers() IViewHelpers
	ContentPath() string
}
