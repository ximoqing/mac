package mac

import (
	"crypto/rand"
	"fmt"
)

// single mac address
func MacSingle() (mac string) {
	buf := make([]byte, 6)
	_, err := rand.Read(buf)
	if err != nil {
		return
	}
	buf[0] |= 2
	mac = fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x", buf[0], buf[1], buf[2], buf[3], buf[4], buf[5])
	return mac
}

// Bulk mac addresses
