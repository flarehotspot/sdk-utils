package models

func SessionTypeToStr(t uint8) string {
	st := SessionType(t)
	return st.String()
}
