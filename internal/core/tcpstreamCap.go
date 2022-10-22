package core

import (
	"fmt"
	"log"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func Tcpstream(capInt string, c chan Flow) {
	var handle *pcap.Handle
	var err error

	ifname := capInt
	handle, err = pcap.OpenLive(ifname, 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	if err := handle.SetBPFFilter("tcp"); err != nil {
		log.Fatal(err)
	}
	log.Println("reading in packets")
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	packets := packetSource.Packets()
	ticker := time.Tick(time.Minute)
	for {
		select {
		case packet := <-packets:
			if packet == nil {
				return
			}
			if packet.NetworkLayer() == nil || packet.TransportLayer() == nil || packet.TransportLayer().LayerType() != layers.LayerTypeTCP {
				continue
			}
			ip := packet.NetworkLayer().(*layers.IPv4)
			// drops incoming packets
			if ip.DstIP.IsPrivate() {
				continue
			}
			c <- Flow{
				DstIP: ip.DstIP,
				SrcIP: ip.SrcIP,
				Hash:  FlowHash(ip.SrcIP, ip.DstIP),
			}
		case <-ticker:
			fmt.Println("tick")
		}
	}

}
