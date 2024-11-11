package sdkfields

type IMultiField interface {
	GetStringValue(row int, name string) (string, error)
	GetIntValue(row int, name string) (int, error)
}
