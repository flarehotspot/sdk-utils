package sdktrans

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"encoding/json"
	"github.com/flarehotspot/core/sdk/utils/paths"
)

const (
	Label MsgType = "label"
	Info  MsgType = "info"
	Error MsgType = "error"
)

var (
	coretrnsdir = filepath.Join(sdkpaths.AppDir, "core/resources/translations")
	sdktrnsdir  = filepath.Join(sdkpaths.AppDir, "sdk/resources/translations")

	Core = NewTranslator(sdkpaths.CoreDir)
	Sdk  = NewTranslator(sdkpaths.SdkDir)
)

type MsgType string

type TranslateFn func(msgtype MsgType, msgkey string, params ...string) string

type LangCfg struct {
	Lang string `yaml:"lang"`
}

func NewTranslator(rootdir string) TranslateFn {
	return func(msgtype MsgType, msgkey string, params ...string) string {
		lang := getLang()
		return translate(rootdir, lang, msgtype, msgkey, params...)
	}
}

func getLang() string {
	cfgPath := filepath.Join(sdkpaths.AppDir, "config/application.json")
	bytes, err := os.ReadFile(cfgPath)
	if err != nil {
		return err.Error()
	}

	var cfg LangCfg
	err = json.Unmarshal(bytes, &cfg)
	if err != nil {
		return err.Error()
	}

	return cfg.Lang
}

func translate(rootdir string, lang string, msgtype MsgType, msgk string, params ...string) string {
	trnsdir := filepath.Join(rootdir, "resources/translations")
	msg := getMsg(trnsdir, lang, string(msgtype), msgk)
	msg = msgWithParams(msg, params...)
	return strings.TrimSpace(msg)
}

func msgWithParams(msg string, params ...string) string {
	for i, p := range params {
		k := fmt.Sprintf("$%d", i)
		msg = strings.ReplaceAll(msg, k, p)
	}
	return msg
}

func getMsg(trnsdir string, lang string, msgtype string, msgk string) string {
	f := filepath.Join(trnsdir, lang, msgtype, msgk+".txt")
	bytes, err := os.ReadFile(f)
	if err != nil {
		return err.Error()
	}
	return string(bytes)
}
