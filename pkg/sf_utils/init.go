package sf_utils

import (
	"dxxproject/config_prepare/app_config"
	"dxxproject/config_prepare/start_config"
	"github.com/bwmarrin/snowflake"
	"github.com/samber/do/v2"
	"time"
)

func NewSnowFlake(injector do.Injector) (snow *SnowflakeIMPL, err error) {
	appConfig := do.MustInvoke[*app_config.AppConfig](injector)
	startConfig := do.MustInvoke[*start_config.StartConfig](injector)
	startTime := appConfig.StartTime
	machineID := startConfig.MachineID

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
