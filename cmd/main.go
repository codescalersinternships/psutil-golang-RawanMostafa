package main

import (
	"log"

	psutils "github.com/codescalersinternships/psutil-golang-RawanMostafa/pkg"
)

func main() {
	cpuInfo, _ := psutils.GetCpuInfo()
	log.Printf("%v",cpuInfo)
	memInfo, err := psutils.GetMemInfo()
	log.Printf("%v",memInfo)
	log.Printf("err %v",err)
	procs, err := psutils.GetProcessList()
	log.Printf("%v\n",procs)
	log.Printf("err %v",err)

}
