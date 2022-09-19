package database

import (
	"context"
	"fmt"

	"github.com/danieldavidpm/generallogin/settings"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/jmoiron/sqlx"
)

func New(ctx context.Context, s *settings.Settings) (*sqlx.DB, error) {
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		s.DB.User,
		s.DB.Password,
		s.DB.Host,
		s.DB.Port,
		s.DB.Name,
	)

	return sqlx.ConnectContext(ctx, "mssql", connectionString)
}
