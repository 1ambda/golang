package interfaces

import (
	"fmt"
)

type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func ExampleIPAddrStringer() {
	addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, addr := range addrs {
		fmt.Printf("%v: %v\n", name, addr)
	}

	// Output might be
	// loopback: 127.0.0.1
	// googleDNS: 8.8.8.8
}
