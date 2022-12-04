package configs

type Options struct {
	// DbDialect is the DB dialect (Gorm driver name: sqlite3)
	DbDialect string `long:"db-dialect" env:"SERVICE_DB_DIALECT"`

	// DbDsn is the DB connection info: ":memory:"
	DbDsn string `long:"db-dsn" env:"SERVICE_DB_DSN"`

	// DbSample enables filling DB by sampe data
	DbSample bool `long:"db-sample" env:"SERVICE_DB_SAMPLE"`

	// DbDebug enables DB debug messages
	DbDebug bool `long:"db-debug" env:"SERVICE_DB_DEBUG"`

	// JwtKey is Identity key to JWT
	JwtKey string `long:"jwt-key" env:"SERVICE_JWT_KEY"`

	// JwtExpireSec is the JWT expire duration in seconds
	JwtExpireSec string `long:"jwt-expire-sec" env:"SERVICE_JWT_EXPIRE_SEC"`
}

const TestingBasePath = ""
