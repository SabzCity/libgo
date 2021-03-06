/* For license and copyright information please see LEGAL file in repository */

package ipv6

import (
	"net"
)

// GetLocalIP : Returns the non loopback local IP of the host.
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	for _, address := range addrs {
		// Check the address type and if it is not a loopback the display it.
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}

			return "\"[" + ipnet.IP.String() + "]\""
		}
	}

	return ""
}
