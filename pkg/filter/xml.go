package filter

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type NmapRun struct {
	XMLName xml.Name `xml:"nmaprun"`
	Hosts   []Host   `xml:"host"`
}

type Host struct {
	Addresses []Address  `xml:"address"`
	Hostnames []Hostname `xml:"hostnames>hostname"`
	Ports     []Port     `xml:"ports>port"`
}

type Address struct {
	Addr     string `xml:"addr,attr"`
	AddrType string `xml:"addrtype,attr"`
}

type Hostname struct {
	Name string `xml:"name,attr"`
	Type string `xml:"type,attr"`
}

type Port struct {
	PortID   string `xml:"portid,attr"`
	Protocol string `xml:"protocol,attr"`
	State    State  `xml:"state"`
}

type State struct {
	State string `xml:"state,attr"`
}

func NmapXML(input io.Reader) {
	inputData, err := io.ReadAll(input)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading XML input stream:", err)
		return
	}

	var run NmapRun
	err = xml.Unmarshal(inputData, &run)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error parsing XML data structure:", err)
		return
	}

	for _, host := range run.Hosts {
		var hostIdentifier string
		for _, hName := range host.Hostnames {
			if hName.Name != "" {
				hostIdentifier = hName.Name
				break
			}
		}

		// Fallback to the IPv4 address if no hostname node exists
		if hostIdentifier == "" {
			for _, addr := range host.Addresses {
				if addr.AddrType == "ipv4" {
					hostIdentifier = addr.Addr
					break
				}
			}
		}

		// Skip processing if absolutely no identifier was found
		if hostIdentifier == "" {
			continue
		}

		for _, port := range host.Ports {
			if port.State.State == "open" {
				fmt.Printf("%s:%s\n", hostIdentifier, port.PortID)
			}
		}
	}
}
