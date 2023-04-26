package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ SearchPermissionModel = (*customSearchPermissionModel)(nil)

type (
	// SearchPermissionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSearchPermissionModel.
	SearchPermissionModel interface {
		searchPermissionModel
	}

	customSearchPermissionModel struct {
		*defaultSearchPermissionModel
	}
)

// NewSearchPermissionModel returns a model for the database table.
func NewSearchPermissionModel(conn sqlx.SqlConn) SearchPermissionModel {
	return &customSearchPermissionModel{
		defaultSearchPermissionModel: newSearchPermissionModel(conn),
	}
}
