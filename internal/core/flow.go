package core

import (
	"fmt"
	"net"
	"time"
)

type FlowCollector struct {
	Counter map[string]*FlowRecord
	Locater IPLocater
}

type Flow struct {
	SrcIP, DstIP net.IP
	DstGeo       GeoInfo
	Hash         string
	Located      bool
}

type FlowRecord struct {
	Flow
	Count int
}

func NewFlowCollector(options ...func(*FlowCollector)) *FlowCollector {
	fc := &FlowCollector{
		Counter: make(map[string]*FlowRecord),
		Locater: GeoIP2Locater{
			dbFile: "GeoLite2-City.mmdb",
		},
	}
	for _, option := range options {
		option(fc)
	}
	return fc
}

func FlowHash(from, to net.IP) string {
	return fmt.Sprintf("%v ==> %v", from.String(), to.String())
}

func (f *FlowCollector) Run(c chan Flow) {
	ticker := time.Tick(time.Second)
	for {
		select {
		case flow := <-c:
			if !flow.Located {
				flow.Located = true
				flow.DstGeo = f.Locater.IPLocation(flow.DstIP)
			}
			f.Add(flow)
		case <-ticker:
			f.Reduce()
		}
	}
}

func (f *FlowCollector) Add(flow Flow) {
	if _, ok := f.Counter[flow.Hash]; !ok {
		f.Counter[flow.Hash] = &FlowRecord{Flow: flow}
	}
	f.Counter[flow.Hash].Count++
}

func (f *FlowCollector) Reduce() {
	for flow := range f.Counter {
		f.Counter[flow].Count--
		if f.Counter[flow].Count < 0 {
			delete(f.Counter, flow)
		}
	}
}
