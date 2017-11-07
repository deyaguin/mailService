package services

type Services struct {
	Template
	Mail
}

func NewServices(
	template Template,
	mail Mail,
) *Services {
	return &Services{
		template,
		mail,
	}
}
