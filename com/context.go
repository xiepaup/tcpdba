package com

import "time"

/**
 *
 * Author  :  xiean
 * EMail   :  xiepaup@163.com 
 * Date    :  2017-03-05
 * Project :  tcpdba
 **/

type TcpdbaContext struct {
	DeviceName      string
	TcpPort         int32
	IntervalSeconds int32  // 2 s
	DurationSeconds int32  // 5*60 s

	OpenTimeOut     int32

	PackageType     string // MySQL or Redis

	OfflineName     string

	OutPutType      int32  // Logfile , MySQLStore


	startTime       time.Time
	endTime         time.Time
}

const (
	STAT_OUTPUT_LOG = 1001
	STAT_OUTPUT_DB = 1003
)

var tcpdbaContext TcpdbaContext

func init() {
	tcpdbaContext = &TcpdbaContext{
		IntervalSeconds:2,
		DurationSeconds:5 * 60,
		OutPutType:1001,
	}
}

func GetTcpdbaContext() *TcpdbaContext {
	return tcpdbaContext
}