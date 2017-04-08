package astimysql

import (
	"fmt"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// New creates a new DB based on a configuration
func New(c Configuration) (*sqlx.DB, error) {
	return sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&loc=UTC&time_zone=%s", c.Username, c.Password, c.Addr, c.Name, url.QueryEscape("'+00:00'")))
}
