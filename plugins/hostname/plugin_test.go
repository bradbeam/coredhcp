package hostname

import (
	"fmt"
	"net"
	"testing"

	"github.com/coredhcp/coredhcp/logger"
	"github.com/insomniacslk/dhcp/dhcpv4"
	"github.com/stretchr/testify/assert"
)

func TestAddHostNameServer4(t *testing.T) {
	// Disable logging for tests
	logger.WithNoStdOutErr(log)

	var mac = net.HardwareAddr{0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff}
	testCases := []struct {
		desc          string
		hostname      string
		expectedError bool
	}{
		{
			desc:     "with hostname",
			hostname: "coredhcp",
		},
		{
			desc:     "without hostname",
			hostname: "",
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%q", tc.desc), func(t *testing.T) {
			var err error

			_, err = setup4()
			assert.NoError(t, err)

			mods := []dhcpv4.Modifier{
				dhcpv4.WithOption(dhcpv4.OptHostName(tc.hostname)),
				dhcpv4.WithRequestedOptions(dhcpv4.OptionHostName),
			}

			req, err := dhcpv4.NewDiscovery(mac, mods...)
			assert.NoError(t, err)

			stub, err := dhcpv4.NewReplyFromRequest(req)
			assert.NoError(t, err)

			resp, stop := Handler4(req, stub)
			if resp == nil {
				t.Fatal("plugin did not return a message")
			}
			if stop {
				t.Error("plugin interrupted processing")
			}

			// Validate response
			assert.Equal(t, tc.hostname, resp.HostName())
		})
	}
}
