package uplink

import (
	_ "fmt"
	_ "gtrack-cli/tracker"
	_ "gtrack-cli/uplink/protocol/headers/category"
	_ "gtrack-cli/uplink/protocol/headers/types"
	_ "net"
)

var (
	errDial = "err: Dialing error"
	buf     = [514]byte{}
)

// func DialSyncExact(address string) (error) {
// 	conn, err := net.Dial("tcp", address)
// 	if err != nil {
// 		fmt.Println(errDial)
// 		fmt.Println(err.Error())
// 		return err
// 	}

// 	buf[0] = category.REQ_TO_SEND
// 	buf
// }
