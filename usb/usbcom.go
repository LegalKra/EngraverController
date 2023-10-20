package usb

import (
	"fmt"

	"github.com/google/gousb"
)

func ListAwaiablePorts() {
	ctx := gousb.NewContext()
	defer ctx.Close()
	vid, pid := gousb.ID(0x04f2), gousb.ID(0xb531)
	devs, err := ctx.OpenDevices(func(desc *gousb.DeviceDesc) bool {
		return desc.Vendor == vid && desc.Product == pid
	})
	if err != nil {
		fmt.Println(err)
		return

	}
	fmt.Println(devs)
}
func conectUsb() {
	//TODO
}
func StartComunication() {
	conectUsb()
	//TODO
}
func EndComunication() {
	//TODO
}
