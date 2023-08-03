package service


import "storie/pkg/domain"


type SendMail interface {
	Send(mail domain.Mail) error
}

var providers = map[string]SendMail{}

func init() {
	providers["gmail"] = Gmail{}
}

func GetProvider(provider string) SendMail {
	return providers[provider]
}
