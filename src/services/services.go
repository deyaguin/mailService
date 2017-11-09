package services

type Services struct {
	Template
	MailService
}

func NewServices(
	template Template,
	mail MailService,
) *Services {
	return &Services{
		template,
		mail,
	}
}
