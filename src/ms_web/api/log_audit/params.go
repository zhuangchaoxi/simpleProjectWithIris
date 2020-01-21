package log_audit

type listCheckOption struct {
	LoginUserName string `validate:"max=30" json:"login_username"`
}
