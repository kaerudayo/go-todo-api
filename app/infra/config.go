package infrastructure

type Config struct {
	DB struct {
		Production struct {
			Host     string
			Username string
			Password string
			DBName   string
		}
	}
	Routing struct {
		Port string
	}
}

func NewConfig() *Config {

	c := new(Config)

	c.DB.Production.Host = "mysql-container"
	c.DB.Production.Username = "root"
	c.DB.Production.Password = "todo"
	c.DB.Production.DBName = "todo"

	c.Routing.Port = ":3001"

	return c
}
