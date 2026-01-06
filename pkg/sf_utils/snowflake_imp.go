package sf_utils

import (
	"github.com/bwmarrin/snowflake"
)

type SnowflakeIMPL struct {
	node *snowflake.Node
}

func (s *SnowflakeIMPL) GenSnowFlakeID() int64 {
	return s.node.Generate().Int64()
}

var _ SnowflakeIF = &SnowflakeIMPL{}
