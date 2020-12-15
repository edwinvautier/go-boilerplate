package database

// Config is the structure used to load db credentials from the environment.
type Config struct {
	DbUser     string `env:"DB_USER"`
	DbPassword string `env:"DB_PASSWORD"`
	DbPort     int    `env:"DB_PORT" envDefault:"5432"`
	DbHost     string `env:"DB_HOST"`
	DbName     string `env:"DB_NAME"`
}