package config

type config struct {
	Mysql mysql
	Redis redis
	Email email
}

type mysql struct {
	Addr     string
	Database string
	Username string
	Password string
	Charset  string
}

type redis struct {
	Addr     string
	Password string
}

type email struct {
	SmtpHost  string
	SmtpEmail string
	SmtpPass  string
}
