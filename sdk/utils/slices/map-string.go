package sdkslices

func MapString(ts []string, f func(string) string) []string {
	us := make([]string, len(ts))
	for i := range ts {
		us[i] = f(ts[i])
	}
	return us
}
