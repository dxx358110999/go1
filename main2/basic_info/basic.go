package basic_info

import (
	"dxxproject/toolkit"
	"github.com/samber/do/v2"
)

type BasicInfo struct {
	LocalIp string
}

func NewBasicInfo(i do.Injector) (bi *BasicInfo, err error) {
	err, ip := toolkit.GetOutboundIP()
	if err != nil {
		panic(err)
	}

	bi = &BasicInfo{
		LocalIp: ip.String(),
	}

	return
}
