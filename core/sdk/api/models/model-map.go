package models

func SessionToMap(model ISession) map[string]any {
	return map[string]any{
		"id":            model.Id(),
		"device_id":     model.DeviceId(),
		"session_type":  SessionTypeToStr(model.SessionType()),
		"time_secs":     model.TimeSecs(),
		"data_mbyte":    model.DataMbyte(),
		"time_consumed": model.TimeConsumed(),
		"data_consumed": model.DataConsumed(),
		"started_at":    model.StartedAt(),
		"exp_days":      model.ExpDays(),
		"expires_at":    model.ExpiresAt(),
		"down_mbits":    model.DownMbits(),
		"up_mbits":      model.UpMbits(),
	}
}

func DeviceToMap(model IDevice) map[string]any {
	return map[string]any{
		"id":       model.Id(),
		"mac_addr": model.MacAddress(),
		"ip_addr":  model.IpAddress(),
		"hostname": model.Hostname(),
	}
}
