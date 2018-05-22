package ip

import (
	"fmt"
	"net"

	"github.com/ipipdotnet/datx-go"
)

func Find(city *datx.City, ip string) {
	i := net.ParseIP(ip)
	if i == nil {
		fmt.Printf("Invalid IP address: %s\n", ip)
		return
	}

	r, err := city.Find(ip)
	if err != nil {
		fmt.Printf("IP query error: %s\n", err)
		return
	}

	fmt.Printf("%#v", r)
}
