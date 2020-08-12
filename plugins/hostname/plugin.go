package hostname

import (
	"github.com/coredhcp/coredhcp/handler"
	"github.com/coredhcp/coredhcp/logger"
	"github.com/coredhcp/coredhcp/plugins"
	"github.com/insomniacslk/dhcp/dhcpv4"
	"github.com/insomniacslk/dhcp/dhcpv6"
)

var log = logger.GetLogger("plugins/hostname")

var hostname string

// Plugin wraps the Hostname plugin information.
var Plugin = plugins.Plugin{
	Name:   "hostname",
	Setup6: setup6,
	Setup4: setup4,
}

func setup6(args ...string) (handler.Handler6, error) {
	return nil, nil
}

func setup4(args ...string) (handler.Handler4, error) {
	log.Println("loaded plugin for DHCPv4.")

	return Handler4, nil
}

func Handler6(req, resp dhcpv6.DHCPv6) (dhcpv6.DHCPv6, bool) {
	return nil, true
}

func Handler4(req, resp *dhcpv4.DHCPv4) (*dhcpv4.DHCPv4, bool) {
	if req.IsOptionRequested(dhcpv4.OptionHostName) {
		resp.Options.Update(dhcpv4.OptHostName(req.HostName()))
	}

	return resp, false
}
