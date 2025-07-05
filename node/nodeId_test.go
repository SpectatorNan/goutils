package unique

import (
	"testing"
)

func TestGetUniqueNodeID(t *testing.T) {

}

func Test_getOutboundIP(t *testing.T) {

	ip, err := getOutboundIP()
	if err != nil {
		t.Fatalf("Failed to get outbound IP: %v", err)
	}
	t.Logf("Outbound IP: %v", ip)

	nodeID := int64(ip[len(ip)-1]) % 1024
	t.Logf("Using IP-based node ID: %d", nodeID)
	t.Logf("Node ID: %d\n", nodeID)
}
