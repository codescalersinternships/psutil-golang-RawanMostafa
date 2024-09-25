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
	log.Fatalf("err %v",err)

}
