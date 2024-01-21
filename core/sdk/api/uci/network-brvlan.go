package sdkuci

import "fmt"

type BrVlan struct {
	Device string
	VlanID int
}

func (self *BrVlan) String() string {
	return fmt.Sprintf("%s.%d", self.Device, self.VlanID)
}
