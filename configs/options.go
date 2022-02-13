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

	// AuthIdentityKey is Identity key to JWT
	JwtKey string `long:"auth-identity-key" env:"SERVICE_JWT_KEY"`
}

const TestingBasePath = "/v1"
