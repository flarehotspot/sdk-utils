package sdkmaps

func Merge[T map[any]any](maps ...T) T {
	merged := make(T)
	for _, m := range maps {
		for k, v := range m {
			merged[k] = v
		}
	}
	return merged
}
