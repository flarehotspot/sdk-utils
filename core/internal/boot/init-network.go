package boot

import (
	"log"

	"core/internal/network"
	"core/internal/utils/nftables"
	"core/internal/utils/ubus"
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
