package configuration

type Configuration interface {
	GetPort() int
	GetHost() string
	GetDebug() bool

	GetDBUserName() string
	GetDBPassword() string
	GetDBHost() string
	GetDBPort() int
	GetDBName() string
}
