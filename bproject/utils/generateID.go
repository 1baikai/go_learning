package utils

import (
	"log"

	"github.com/bwmarrin/snowflake"
)

// snow flake id
var node *snowflake.Node

func init() {
	// Create a new NodeID with a NodeID number of 1
	var err error
	node, err = snowflake.NewNode(1)
	if err != nil {
		log.Fatalln("init snowflake node with err", err)
	}
}

// GenID generate a snowflake ID
func GenID() int64 {
	return node.Generate().Int64()
}
