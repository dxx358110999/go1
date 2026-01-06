package sms_svc

import (
	"context"
	"dxxproject/internal/svc/sms_svc/sms_provider"
	"github.com/samber/do/v2"
	"sync"
	"sync/atomic"
)

type SvcSmsFailOver struct {
	smsProviders  []sms_provider.SmsProviderIF
	lock          sync.Mutex
	providerIndex int64
	failCount     int64
	maxFail       int64
}

var _ SvcSmsIface = (*SvcSmsFailOver)(nil)

func (r *SvcSmsFailOver) checkFailCountAndChangeProvider() {
	r.lock.Lock()
	defer r.lock.Unlock()
	if r.failCount > r.maxFail {
		r.providerIndex = (r.providerIndex + 1) % int64(len(r.smsProviders)) //切换provider
		r.failCount = 0                                                      //重置count
	}
}

func (r *SvcSmsFailOver) SendSms(ctx context.Context, info sms_provider.SmsSendInfo) (err error) {
	r.checkFailCountAndChangeProvider()
	err = r.smsProviders[r.providerIndex].SendSms(ctx, info)
	if err != nil {
		//if errors.Is(err, context.DeadlineExceeded) { return }
		atomic.AddInt64(&r.failCount, 1) //原子操作,这里所有错误都切换
		return
	}
	return
}

func NewSmsFailOverSvc(injector do.Injector) (*SvcSmsFailOver, error) {
	smsProviders := do.MustInvoke[[]sms_provider.SmsProviderIF](injector)
	//if err != nil {
	//	return nil, err
	//}
	sms := &SvcSmsFailOver{
		smsProviders:  smsProviders,
		lock:          sync.Mutex{},
		providerIndex: 0,
		failCount:     0,
		maxFail:       1,
	}
	return sms, nil
}
