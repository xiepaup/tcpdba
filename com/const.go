package com

import "time"

/**
 *
 * Author  :  xiean
 * EMail   :  xiepaup@163.com 
 * Date    :  2017-03-05
 * Project :  tcpdba
 **/

type pcakageHeader struct {
	srcIp   string
	srcPort int32
	dstIp   string
	dstPort int32
	pkgTime time.Time
	pkgseq  int64
}

type pkg4MysqlBody struct {
	myUser      string
	mySql       string
	myTable     string
	myDatabase  string
	myOperation string
}

type pkg4RedisBody struct {
	dbSeq   int32
	keyName string
	keyType string
	optCmd  string
}

