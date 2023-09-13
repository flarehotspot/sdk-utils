package boot

import (
	"log"

	"github.com/flarehotspot/core/network"
	"github.com/flarehotspot/core/utils/nftables"
	"github.com/flarehotspot/core/utils/ubus"
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
