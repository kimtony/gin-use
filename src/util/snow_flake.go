package util

import (
	"github.com/bwmarrin/snowflake"
	"fmt"
)

var node *snowflake.Node


// GenerateSnowflake generate Twitter Snowflake ID
func GenerateId() string {

	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
	}
	id := fmt.Sprintf("%s",node.Generate())
	return id
}

