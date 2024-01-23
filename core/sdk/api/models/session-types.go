package sdkmodels

// List of possible session types.
const (
	SessionTypeTime SessionType = iota
	SessionTypeData
	SessionTypeTimeOrData
)

type SessionType uint8

func (self SessionType) ToUint8() uint8 {
	return uint8(self)
}

func (self SessionType) String() string {
	switch self {
	case SessionTypeTime:
		return "time"
	case SessionTypeData:
		return "data"
	case SessionTypeTimeOrData:
		return "time_or_data"
	default:
		return "invalid_session_type"
	}
}
