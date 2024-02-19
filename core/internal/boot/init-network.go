package boot

import (
	"log"

	"github.com/flarehotspot/core/internal/network"
	"github.com/flarehotspot/core/internal/utils/nftables"
	"github.com/flarehotspot/core/internal/utils/ubus"
)

func InitNetwork() (err error) {

	err = nftables.Setup()
	if err != nil {
		log.Println(err)
		return err
	}

	err = network.SetupLanInterfaces()
	if err != nil {
		log.Println(err)
		return err
	}

	ubus.Listen()

	return nil
}
