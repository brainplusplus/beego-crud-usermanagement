package util

import (
	"github.com/bwmarrin/snowflake"
	"strconv"
)

var node *snowflake.Node

// InitSnowflake initiate Snowflake node singleton.
func InitSnowflake() error {
	// Get node number from env TIX_NODE_NO
	key := "1"
	// Parse node number
	nodeNo, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		return err
	}
	// Create snowflake node
	n, err := snowflake.NewNode(nodeNo)
	if err != nil {
		return err
	}
	// Set node
	node = n
	return nil
}

// GenerateSnowflake generate Twitter Snowflake ID
func GenerateSnowflake() string {
	return node.Generate().String()
}