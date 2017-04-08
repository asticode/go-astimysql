package astimysql

import "flag"

// Flags
var (
	DBAddr     = flag.String("db-addr", "", "the db addr")
	DBName     = flag.String("db-name", "", "the db name")
	DBPassword = flag.String("db-password", "", "the db password")
	DBUsername = flag.String("db-username", "", "the db username")
)

// Configuration represents the MySQL configuration
type Configuration struct {
	Addr     string `toml:"addr"`
	Name     string `toml:"name"`
	Password string `toml:"password"`
	Username string `toml:"username"`
}

// FlagConfig generates a Configuration based on flags
func FlagConfig() Configuration {
	return Configuration{
		Addr:     *DBAddr,
		Name:     *DBName,
		Password: *DBPassword,
		Username: *DBUsername,
	}
}
