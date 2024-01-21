package sdkhttp

type ILayoutData interface {
	Helpers() IHelpers
	ContentPath() string
}
