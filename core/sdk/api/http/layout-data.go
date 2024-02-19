package sdkhttp

type ILayoutData interface {
	Helpers() HttpHelpers
	ContentPath() string
}
