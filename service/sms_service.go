package service

type SmsService interface {
	SendSms(payload string, number uint)
}

type SmsServiceImpl struct {
}

func (service *SmsServiceImpl) SendSms(payload string, number uint) {
	//TODO Implement SMS sending service
}

func NewSmsService() SmsService {
	return &SmsServiceImpl{}
}
