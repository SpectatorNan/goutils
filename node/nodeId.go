package unique

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"net"
	"os"
	"strconv"
	"time"
)

func GetUniqueNodeID() int64 {
	return getUniqueNodeID()
}

// Replace the existing snowflake initialization
func getUniqueNodeID() int64 {
	// Try environment variable first (most reliable for containerized environments)
	if nodeIDStr := os.Getenv("SNOWFLAKE_NODE_ID"); nodeIDStr != "" {
		nodeID, err := strconv.ParseInt(nodeIDStr, 10, 64)
		if err == nil && nodeID >= 0 && nodeID < 1024 {
			logx.Infof("Using node ID from environment: %d", nodeID)
			fmt.Printf("Node ID: %d\n", nodeID)
			return nodeID
		}
		logx.Errorf("Invalid node ID from environment: %s", nodeIDStr)
		fmt.Printf("Invalid node ID from environment: %s\n", nodeIDStr)
	}

	// Fall back to IP-based node ID
	ip, err := getOutboundIP()
	if err == nil {
		// Use last octet of IP as node ID
		nodeID := int64(ip[len(ip)-1]) % 1024
		logx.Infof("Using IP-based node ID: %d", nodeID)
		fmt.Printf("Node ID: %d\n", nodeID)
		return nodeID
	}

	timestamp := time.Now().Unix()
	// Default fallback
	fmt.Printf("Using default node ID based on timestamp: %d\n", timestamp)
	logx.Infof("Using default node ID based on timestamp: %d", timestamp)
	return timestamp
}

func getOutboundIP() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP, nil
}
