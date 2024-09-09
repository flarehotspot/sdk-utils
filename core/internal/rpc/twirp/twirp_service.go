package twirp

import (
	"context"
	core_machine_v0_0_1 "core/internal/rpc/machines/coremachines/v0_0_1"
	"log"
	"net/http"

	"github.com/twitchtv/twirp"
)

func GetCoreMachineTwirpServiceAndCtx() (core_machine_v0_0_1.CoreMachineService, context.Context) {
	isDev := true

	proto := "http"
	prefix := "v0.0.1"
	domain := "flarehotspot.com"
	subdomain := "rpc-machines"

	if isDev {
		domain = "flarehotspot-dev.com"
	}

	baseUrl := subdomain + "." + domain
	url := proto + "://" + baseUrl + "/" + prefix

	srv := core_machine_v0_0_1.NewCoreMachineServiceProtobufClient(url, &http.Client{})
	header := make(http.Header)
	header.Set("Authorization", "Bearer "+"xxxxxxxxxx")

	ctx := context.Background()
	ctx, err := twirp.WithHTTPRequestHeaders(ctx, header)
	if err != nil {
		log.Fatalf("twirp error setting headers: %s", err)
	}

	return srv, ctx
}
