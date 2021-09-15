package net

import (
	"log"
	"net"
	"net/http"
	"os"

	"github.com/pumphouse-p/peek-go/pkg/apiutils"
)

type IPAddress struct {
	IP   string `json:"ip"`
	Zone string `json:"zone,omitempty"`
}

type NetworkInterface struct {
	Index        int         `json:"index"`
	MTU          int         `json:"mtu"`
	Name         string      `json:"name"`
	HardwareAddr string      `json:"hwaddr"`
	Flags        *net.Flags  `json:"flags,omitempty"`
	IPAddress    []IPAddress `json:"ipaddr"`
}

type NetStatus struct {
	Hostname   string             `json:"hostname"`
	Interfaces []NetworkInterface `json:"interfaces,omitempty"`
}

type Net struct{}

func New() *Net {
	return &Net{}
}

func GetHostname() (h string) {
	h, _ = os.Hostname()
	return h
}

func GetInterfaces() []NetworkInterface {

	interfaces, err := net.Interfaces()
	var sanitized []NetworkInterface

	if err != nil {
		log.Printf("Could not retrieve network interfaces")
	}

	log.Printf("Got %d interfaces", len(interfaces))

	for _, e := range interfaces {
		log.Printf("Collecting info for interface: %v", e.Name)
		iface := NetworkInterface{}
		iface.MTU = e.MTU
		iface.Name = e.Name
		iface.Index = e.Index
		iface.HardwareAddr = e.HardwareAddr.String()

		log.Printf("Collecting IP addresses")
		addr, _ := e.Addrs()

		for _, a := range addr {
			ip := IPAddress{
				IP: a.String(),
			}
			iface.IPAddress = append(iface.IPAddress, ip)
		}

		sanitized = append(sanitized, iface)
	}

	return sanitized
}

func (n *Net) APIGet(w http.ResponseWriter, r *http.Request) {
	status := NetStatus{}

	status.Hostname = GetHostname()
	status.Interfaces = GetInterfaces()

	apiutils.ServeJSON(w, status)
}
