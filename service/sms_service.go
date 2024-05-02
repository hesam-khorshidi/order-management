package service

type SmsService interface {
	SendSms(payload string, number int)
}

type SmsServiceImpl struct {
}

func (service *SmsServiceImpl) SendSms(payload string, number uint) {
	//Implement SMS sending service
}
