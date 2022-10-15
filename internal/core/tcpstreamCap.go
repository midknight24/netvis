package core

import (
	"fmt"
	"log"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func Tcpstream() {
	var handle *pcap.Handle
	var err error

	ifname := "\\Device\\NPF_{AE039499-7655-476B-9119-7BC13F7F3CEC}"
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
			fmt.Println(ip.SrcIP, ip.DstIP)
		case <-ticker:
			fmt.Println("tick")
		}
	}

}
