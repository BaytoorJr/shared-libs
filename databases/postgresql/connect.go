package postgresql

import (
	"context"
	"github.com/BaytoorJr/shared-libs/errors"
	"github.com/jackc/pgx/v4/pgxpool"
	"strconv"
)

// InitConnect
// initialize PostgreSQL connection
func InitConnect(ctx context.Context, maxConns int, host, port, user, pass, dbName string) (*pgxpool.Pool, error) {
	dsn := "postgres://" +
		user + ":" +
		pass + "@" +
		host + ":" +
		port + "/" +
		dbName +
		"?sslmode=disable" +
		"&pool_max_conns=" + strconv.Itoa(maxConns)

	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, errors.InternalServerError.SetDevMessage(err.Error())
	}

	conn, err := pgxpool.ConnectConfig(ctx, cfg)
	if err != nil {
		return nil, errors.DBConnectError.SetDevMessage(err.Error())
	}

	return conn, nil
}
