package entities

import (
	"net"

	"github.com/containers/libpod/pkg/specgen"
	"github.com/cri-o/ocicni/pkg/ocicni"
)

type Container struct {
	IdOrNamed
}

type Volume struct {
	Identifier
}

type Report struct {
	Id  []string
	Err map[string]error
}

type PodDeleteReport struct{ Report }

type VolumeDeleteOptions struct{}
type VolumeDeleteReport struct{ Report }

// NetOptions reflect the shared network options between
// pods and containers
type NetOptions struct {
	AddHosts     []string
	CNINetworks  []string
	DNSHost      bool
	DNSOptions   []string
	DNSSearch    []string
	DNSServers   []net.IP
	Network      specgen.Namespace
	NoHosts      bool
	PublishPorts []ocicni.PortMapping
	StaticIP     *net.IP
	StaticMAC    *net.HardwareAddr
}

// All CLI inspect commands and inspect sub-commands use the same options
type InspectOptions struct {
	Format string `json:",omitempty"`
	Latest bool   `json:",omitempty"`
	Size   bool   `json:",omitempty"`
}
