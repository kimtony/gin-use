package util

import (
	"github.com/bwmarrin/snowflake"
	"fmt"
)

var node *snowflake.Node


// GenerateSnowflake generate Twitter Snowflake ID
func GenerateId() int64 {

	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
	}
	id := node.Generate().Int64()

	return id
}

