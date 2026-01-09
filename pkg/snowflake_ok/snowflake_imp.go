package snowflake_ok

import (
	"dxxproject/config_prepare/app_config"
	"dxxproject/config_prepare/start_config"
	"github.com/bwmarrin/snowflake"
	"time"
)

type SnowflakeIMPL struct {
	node *snowflake.Node
}

func (s *SnowflakeIMPL) GenSnowFlakeID() int64 {
	return s.node.Generate().Int64()
}

var _ SnowflakeIface = &SnowflakeIMPL{}

func NewSnowFlake(startCfg *start_config.StartConfig,
	appCfg *app_config.AppConfig) (snow *SnowflakeIMPL, err error) {

	startTime := appCfg.StartTime
	machineID := startCfg.MachineID

	var timeStart time.Time
	timeStart, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	snowflake.Epoch = timeStart.UnixNano() / 1000000
	node, err := snowflake.NewNode(machineID)
	if err != nil {
		return
	}

	snow = &SnowflakeIMPL{node}
	return
}
