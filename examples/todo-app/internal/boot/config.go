package boot

var _ Config = (*config)(nil)

type (
	Config interface {
		SQLiteDSN() string
		PostgreSQLDSN() string
	}

	config struct{}
)

func NewConfig() Config {
	return &config{}
}

func (c *config) PostgreSQLDSN() string {
	return "postgres://postgres:postgres@localhost:5432/postgres" // TODO load from envs
}

func (c *config) SQLiteDSN() string {
	return "file:./test.db?mode=memory&cache=shared" // TODO load from envs
}
