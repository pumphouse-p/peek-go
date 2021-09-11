package net

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"

	"github.com/pumphouse-p/peek-go/pkg/apiutils"
)

type NetStatus struct {
	Hostname   string                    `json:"hostname"`
	Interfaces map[int]map[string]string `json:"interfaces"`
}

type Net struct{}

func New() *Net {
	return &Net{}
}

func CollectHostname() (h string) {
	h, _ = os.Hostname()
	return h
}

func CollectNetInfo() map[int]map[string]string {
	netinfo := make(map[int]map[string]string)

	interfaces, err := net.Interfaces()
	if err != nil {
		log.Printf("Could not retrieve network interfaces")
	}

	for _, e := range interfaces {
		log.Printf("Collecting info for interface: %v", e.Name)
		iface := make(map[string]string)
		iface["mtu"] = strconv.Itoa(e.MTU)
		iface["name"] = e.Name
		iface["mac-address"] = e.HardwareAddr.String()
		log.Printf("Collecting IP addresses")
		addr, _ := e.Addrs()
		for i, a := range addr {
			k := fmt.Sprintf("ip-%d", i)
			iface[k] = a.String()
		}
		netinfo[e.Index] = iface
	}

	return netinfo
}

func (n *Net) APIGet(w http.ResponseWriter, r *http.Request) {
	status := NetStatus{}

	status.Hostname = CollectHostname()
	status.Interfaces = CollectNetInfo()

	apiutils.ServeJSON(w, status)
}
