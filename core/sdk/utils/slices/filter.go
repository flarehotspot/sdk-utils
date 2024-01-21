package sdkslices

func Filter[T any](collection []T, filterFn func(item T) bool) []T {
	var newArr []T = []T{}
	for _, a := range collection {
		ok := filterFn(a)
		if ok {
			newArr = append(newArr, a)
		}
	}
	return newArr
}
