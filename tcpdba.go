package main


import (
	"time"
	"xiepaup.com/tcpdba/com"
	"github.com/google/gopacket/pcap"
	"flag"
	"fmt"
	"os"
	"xiepaup.com/tcpdba/log"
)

/**
 *
 * Author  :  xiean
 * EMail   :  xiepaup@163.com
 * Date    :  2017-03-05
 * Project :  tcpdba
 **/


var AppVersion string

var (
	device string = "en0"
	snapshotLen int32 = 1024
	promiscuous bool = false
	err error
	timeout time.Duration = 30 * time.Second
	handle      *pcap.Handle
)

func main() {

	tcpdbContext := com.GetTcpdbaContext()

	flag.StringVar(&tcpdbContext.DeviceName, "device", "eth0", "Network device name [eht0,en0..] ,default : eth0")
	flag.StringVar(&tcpdbContext.TcpPort, "port", 3306, "Instance listen port ,default : 3306")
	flag.StringVar(&tcpdbContext.PackageType, "type", "mysql", "instance type,for now support mysql and redis ,default : mysql ")
	flag.StringVar(&tcpdbContext.IntervalSeconds, "interval", 2, "Status window ,seconds ,default : 2")
	flag.StringVar(&tcpdbContext.DurationSeconds, "duration", 300, "total run time ,seconds ,default : 300")
	flag.StringVar(&tcpdbContext.FilterSrcIp,"srcip","","only capture srcip package ,default None")

	quiet := flag.Bool("quiet", false, "quiet,only show error msg !")
	verbose := flag.Bool("verbose", false, "show detail infomation !")
	debug := flag.Bool("debug", false, "debug mode (very verbose)")
	stack := flag.Bool("stack", false, "add stack trace upon error")
	help := flag.Bool("help", false, "Display usage")
	version := flag.Bool("version", false, "Print version & exit")

	flag.Parse()

	if *help {
		fmt.Fprintf(os.Stderr, "Usage of tcpdba:\n")
		flag.PrintDefaults()
		return
	}
	if *version {
		appVersion := AppVersion
		if appVersion == "" {
			appVersion = "unversioned"
		}
		fmt.Println(appVersion)
		return
	}
	log.SetLevel(log.ERROR)
	if *verbose {
		log.SetLevel(log.INFO)
	}
	if *debug {
		log.SetLevel(log.DEBUG)
	}
	if *stack {
		log.SetPrintStackTrace(*stack)
	}
	if *quiet {
		// Override!!
		log.SetLevel(log.ERROR)
	}

	// Check args
	checkParametersMustInput(&tcpdbContext)

	handle, err = pcap.OpenLive(device, snapshotLen, promiscuous, timeout)

	if err != nil {

		log.Fatal(err)
	}

	defer handle.Close()
}

func checkParametersMustInput(tcpdbContext *com.TcpdbaContext) {

	if (tcpdbContext.DeviceName == "") {
		log.Fatalf("--device must speicfied !")
	}

	if (tcpdbContext.TcpPort == "") {
		log.Fatalf("--port must specified !")
	}
	if (tcpdbContext.PackageType == "") {
		log.Fatalf("--type must specified ,you can specifie mysql or redis !")
	}
}