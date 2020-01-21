package config

type listCheckOption struct {
	Name  string `validate:"gte=0,lte=100" json:"name"`
	Email string `validate:"email" json:"email"`
}
