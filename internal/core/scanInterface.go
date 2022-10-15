package core

import (
	"log"
	"net"

	"github.com/google/gopacket/pcap"
)

type deviceInterface struct {
	Name, Description string
	IPs               []IPAddress
}

type IPAddress struct {
	IP   net.IP
	Mask net.IPMask
}

// scanInterface returns all interfaces in the local machine
func ScanInterface() []deviceInterface {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}
	interfaces := make([]deviceInterface, 0)
	for _, device := range devices {
		ips := make([]IPAddress, 0)
		for _, address := range device.Addresses {
			ips = append(ips, IPAddress{address.IP, address.Netmask})
		}
		interfaces = append(interfaces, deviceInterface{
			Name:        device.Name,
			Description: device.Description,
			IPs:         ips,
		})
	}
	return interfaces
}
