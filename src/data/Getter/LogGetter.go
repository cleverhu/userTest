package Getter

import (
	"userTest/src/dbs"
	"userTest/src/models/LogModel"
)

var LogGetter ILogGetter

func init() {
	LogGetter = NewLogGetterImpl()
}

type ILogGetter interface {
	GetLogList() []*LogModel.LogImpl
}

type LogGetterImpl struct {
}

func NewLogGetterImpl() *LogGetterImpl {
	return &LogGetterImpl{}
}

func (*LogGetterImpl) GetLogList() (logs []*LogModel.LogImpl) {
	dbs.Orm.Find(&logs)
	return
}
