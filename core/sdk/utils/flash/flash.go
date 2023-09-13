package flash

import (
	"log"
	"net/http"
	"time"

	"github.com/flarehotspot/core/sdk/utils/cookie"
)

const (
	Success FlashType = "success"
	Info    FlashType = "info"
	Warning FlashType = "warning"
	Error   FlashType = "error"
)

type FlashType string

func SetFlashMsg(w http.ResponseWriter, t FlashType, msg string) {
	log.Println("set flash_"+t, msg)
	name := "flash_" + string(t)
	cookie.SetCookie(w, name, msg)
}

func GetFlashMsg(w http.ResponseWriter, r *http.Request, t string) string {
	name := "flash_" + t
	c, err := r.Cookie(name)
	if err != nil {
		switch err {
		case http.ErrNoCookie:
			return ""
		default:
			return ""
		}
	}

  log.Println("Delete cookie: " + name)
	dc := &http.Cookie{Name: name, MaxAge: -1, Expires: time.Unix(1, 0), Path: "/"}
	http.SetCookie(w, dc)

	return c.Value
}
