package main

import (
	"log"

	psutils "github.com/codescalersinternships/psutil-golang-RawanMostafa/pkg"
)

func main() {
	cpuinfo, _ := psutils.GetCpuInfo()
	log.Printf("%v",cpuinfo)

}
