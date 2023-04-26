package svc

import (
	"search/internal/config"
	"search/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config                config.Config
	SearchPermissionModel model.SearchPermissionModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                c,
		SearchPermissionModel: model.NewSearchPermissionModel(sqlx.NewMysql(c.Mysql.DataSource)),
	}
}
