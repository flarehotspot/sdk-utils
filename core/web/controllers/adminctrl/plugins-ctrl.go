package adminctrl

// import (
// 	"fmt"
// 	"io"
// 	"log"
// 	"net/http"
// 	"os"
// 	"os/exec"
// 	"path/filepath"
// 	"sync"

// 	"github.com/flarehotspot/core/accounts"
// 	"github.com/flarehotspot/core/plugins"
// 	"github.com/flarehotspot/core/sdk/utils/contexts"
// 	"github.com/flarehotspot/core/sdk/utils/paths"
// 	"github.com/flarehotspot/core/sdk/utils/strings"
// 	"github.com/flarehotspot/core/web/response"
// 	"github.com/flarehotspot/core/web/router"
// 	"github.com/flarehotspot/core/web/routes/names"
// )

// type InstallOut struct{ acct *accounts.Account }

// func (i *InstallOut) Write(p []byte) (n int, err error) {
// 	status := map[string]any{"msg": string(p)}
// 	i.acct.Emit("plugin:install:progress", status)
// 	return len(p), nil
// }

// type PluginCtrl struct {
// 	mu     sync.RWMutex
// 	pmgr   *plugins.PluginsMgr
// 	result *plugins.InstPrgrs
// 	capi   *plugins.PluginApi
// }

// func NewPluginsCtrl(pmgr *plugins.PluginsMgr, capi *plugins.PluginApi) *PluginCtrl {
// 	return &PluginCtrl{
// 		mu:   sync.RWMutex{},
// 		pmgr: pmgr,
// 		capi: capi,
// 	}
// }

// func (self *PluginCtrl) Index(w http.ResponseWriter, r *http.Request) {
// 	plugins := self.pmgr.All()
// 	newPluginUrl, _ := router.UrlForRoute(routenames.RouteAdminPluginsNew)

// 	data := map[string]any{
// 		"newPluginUrl": newPluginUrl,
// 		"plugins":      plugins,
// 	}

// 	self.capi.HttpApi().Respond().AdminView(w, r, "plugins/index.html", data)
// }

// func (self *PluginCtrl) New(w http.ResponseWriter, r *http.Request) {
// 	uploadUrl, _ := router.UrlForRoute(routenames.RouteAdminPluginUpload)
// 	data := map[string]any{"uploadUrl": uploadUrl}
// 	self.capi.HttpApi().Respond().AdminView(w, r, "plugins/upload.html", data)
// }

// func (self *PluginCtrl) Upload(w http.ResponseWriter, r *http.Request) {
// 	self.mu.Lock()
// 	defer self.mu.Unlock()

// 	// Parse our multipart form, 10 << 20 specifies a maximum
// 	// upload of 10 MB files.
// 	r.ParseMultipartForm(10 << 20)
// 	file, _, err := r.FormFile("plugin_zip")
// 	if err != nil {
// 		fmt.Println("Error Retrieving the File")
// 		fmt.Println(err)
// 		self.uploadErr(w, r, err)
// 		return
// 	}
// 	defer file.Close()

// 	uploadDir := filepath.Join(paths.TmpDir, "uploads/plugins")
// 	os.MkdirAll(uploadDir, os.ModePerm)

// 	tempFile, err := os.CreateTemp(uploadDir, "plugin-*.zip")
// 	if err != nil {
// 		fmt.Println(err)
// 		self.uploadErr(w, r, err)
// 		return
// 	}
// 	defer tempFile.Close()

// 	// read all of the contents of our uploaded file into a
// 	// byte array
// 	fileBytes, err := io.ReadAll(file)
// 	if err != nil {
// 		fmt.Println(err)
// 		self.uploadErr(w, r, err)
// 		return
// 	}
// 	// write this byte array to our temporary file
// 	_, err = tempFile.Write(fileBytes)
// 	if err != nil {
// 		log.Println(err)
// 		self.uploadErr(w, r, err)
// 		return
// 	}

// 	log.Println("Zip file written to: ", tempFile.Name())

// 	extDir := filepath.Join(paths.TmpDir, strings.Rand(32))

// 	unzipCmd := exec.Command("unzip", tempFile.Name(), "-d", extDir)
// 	err = unzipCmd.Run()
// 	if err != nil {
// 		self.uploadErr(w, r, err)
// 		return
// 	}

// 	go func() {
// 		tags := r.PostFormValue("tags")
// 		log.Println("tags: ", tags)
// 		acctSym := r.Context().Value(contexts.SysAcctCtxKey)
// 		acct := acctSym.(*accounts.Account)
// 		out := &InstallOut{acct: acct}
// 		err = plugins.Build(out, extDir, "-tags", tags)
// 		if err != nil {
// 			self.uploadErr(w, r, err)
// 			return
// 		}
// 	}()

// 	w.WriteHeader(http.StatusOK)
// }

// func (self *PluginCtrl) uploadErr(w http.ResponseWriter, r *http.Request, err error) {
// 	w.WriteHeader(http.StatusInternalServerError)
// 	response.Json(w, map[string]any{"error": err.Error()}, 500)
// }
