package data

import (
	"context"
	"errors"
	"strconv"
	"yayar/internal/data/ent"
	"yayar/internal/data/ent/migrate"
	"yayar/pkg/utils"

	"entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Data struct {
	Client *ent.Client
	Driver *sql.Driver
}
type ConfigMysql struct {
	User     string
	Password string
	Database string
	Host     *string
	Port     *string
}
type ConfigRedis struct {
	Host *string
	Port *string
}
type InitialConf struct {
	Mysql *ConfigMysql
	Redis *ConfigRedis
}

func NewData(conf InitialConf) (*Data, error) {
	if conf.Mysql != nil {
		c, d, err := initMysql(*conf.Mysql)
		if err != nil {
			return nil, err
		}
		return &Data{Client: c, Driver: d}, nil
	}
	return nil, errors.New("DataBase Initialize Error: db config not found")
}

func initMysql(conf ConfigMysql) (*ent.Client, *sql.Driver, error) {
	db_name := "mysql"
	host := "localhost"
	port := 3306
	if conf.Host != nil {
		host = *conf.Host
	}
	if conf.Port != nil {
		_port, err := strconv.Atoi(*conf.Port)
		if err == nil {
			port = _port
		}
	}
	dsn := utils.DB.DSN(utils.DSNConfig{
		Username: conf.User,
		Password: conf.Password,
		Host:     host,
		Port:     port,
		Database: conf.Database,
		Pairs:    utils.DB.DefaultDSNPair(),
	})
	drv, err := sql.Open(db_name, dsn)
	if err != nil {
		return nil, nil, err
	}
	db := drv.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	client := ent.NewClient(ent.Driver(drv))
	ctx := context.Background()
	err = client.Schema.Create(
		ctx,
		migrate.WithDropColumn(true),
		migrate.WithDropIndex(true),
	)
	if err != nil {
		return nil, nil, err
	}
	return client, drv, nil
}
