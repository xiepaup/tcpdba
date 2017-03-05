package main

/**
 *
 * Author  :  xiean
 * EMail   :  xiepaup@163.com
 * Date    :  2017-03-05
 * Project :  tcpdba
 **/

import (
	"log"
	"time"
	"xiepaup.com/tcpdba/com"
	"github.com/google/gopacket/pcap"
	"flag"
)


var (
	device string = "en0"
	snapshotLen int32 = 1024
	promiscuous bool = false
	err error
	timeout time.Duration = 30 * time.Second
	handle      *pcap.Handle
)



func main() {
	log.Print("Hello TCPDBA !")

	tcpdbContext := com.GetTcpdbaContext()

	flag.StringVar(&tcpdbContext.DeviceName,"DeviceName","eth0","Network device name [eht0,en0..]")
	flag.StringVar(&tcpdbContext.TcpPort,"port","3306","Instance listen port")
	flag.StringVar(&tcpdbContext.IntervalSeconds,"interval",2,"Status window ,seconds")
	flag.StringVar(&tcpdbContext.DurationSeconds,"duration",300,"total run time ,seconds")

	handle, err = pcap.OpenLive(device, snapshotLen, promiscuous, timeout)

	if err != nil {

		log.Fatal(err)
	}

	defer handle.Close()
}