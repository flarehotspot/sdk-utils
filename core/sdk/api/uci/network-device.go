package sdkuci

type IDevice interface {
	Name() string
	Type() string
	BrPorts() []string
}
