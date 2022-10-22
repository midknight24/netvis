package backend

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/midknight24/netvis/internal/core"
)

type FlowServer struct {
	fc      core.FlowCollector
	capInt  string
	address string
	port    int
}

func New(capInt string, options ...func(*FlowServer)) *FlowServer {
	svr := &FlowServer{
		fc:      *core.NewFlowCollector(),
		capInt:  capInt,
		address: "127.0.0.1:80",
	}
	for _, option := range options {
		option(svr)
	}
	return svr
}

func (fs *FlowServer) Run() {
	c := make(chan core.Flow)
	go fs.fc.Run(c)
	go core.Tcpstream(fs.capInt, c)
	log.Fatal(http.ListenAndServe(fs.address, fs.Handler()))
}
  
func (fs *FlowServer) Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		highFlows := []*core.Flow{}
		for _, flow := range fs.fc.Counter {
			if flow.Count > 5 {
				highFlows = append(highFlows, &flow.Flow)
			}
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		data, err := json.Marshal(highFlows)
		if err != nil {
			w.Write([]byte(err.Error()))
		}
		w.Write(data)
	}
}
