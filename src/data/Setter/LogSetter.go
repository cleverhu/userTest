package Setter

import (
	"userTest/src/models/LogModel"
	"userTest/src/result"
)

var LogSetter ILogSetter

func init() {
	LogSetter = NewLogSetterImpl()
}

type ILogSetter interface {
	AddLog(log *LogModel.LogImpl) *result.ErrorResult
}

type LogSetterImpl struct {
}

func NewLogSetterImpl() *LogSetterImpl {
	return &LogSetterImpl{}
}

func (*LogSetterImpl) AddLog(log *LogModel.LogImpl) *result.ErrorResult {
	return nil
}
