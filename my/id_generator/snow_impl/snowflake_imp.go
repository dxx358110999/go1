package snow_impl

import (
	"dxxproject/config_prepare/app_config"
	"dxxproject/config_prepare/start_config"
	"dxxproject/my/id_generator"
	"github.com/bwmarrin/snowflake"
	"time"
)

type SnowflakeIMPL struct {
	node *snowflake.Node
}

func (s *SnowflakeIMPL) GenSnowFlakeID() int64 {
	return s.node.Generate().Int64()
}

var _ id_generator.IdGenIface = new(SnowflakeIMPL)

func NewSnowFlake(startCfg *start_config.Config,
	appCfg *app_config.Config) (snow *SnowflakeIMPL, err error) {

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
