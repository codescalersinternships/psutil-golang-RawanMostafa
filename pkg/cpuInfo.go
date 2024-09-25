package psutils

import (
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

var errKeyNotExist = errors.New("key doesn't exist")

type CpuInfo struct {
	NumCores  int
	VendorId  string
	ModelName string
	CacheSize string
	CpuMHZ    float32
}

func loadCpuFile() (cpuInfo string, err error) {
	file, err := os.ReadFile("/proc/cpuinfo")
	cpuInfo = string(file)

	return
}

func GetCpuInfo() (cpuInfo CpuInfo, err error) {
	var processorsNum int
	var totalFreq float32

	cpuFile, err := loadCpuFile()
	if err != nil {
		return
	}
	lines := strings.SplitN(cpuFile, "\n", -1)
	for _, line := range lines {
		if strings.HasPrefix(line, "processor") {
			processorsNum++
		} else if strings.HasPrefix(line, "vendor_id") {
			parts := strings.Split(line, ":")
			cpuInfo.VendorId = strings.TrimSpace(parts[1])
		} else if strings.HasPrefix(line, "model name") {
			parts := strings.Split(line, ":")
			cpuInfo.ModelName = strings.TrimSpace(parts[1])
		} else if strings.HasPrefix(line, "cache size") {
			parts := strings.Split(line, ":")
			cpuInfo.CacheSize = strings.TrimSpace(parts[1])
		} else if strings.HasPrefix(line, "cpu MHz") {
			parts := strings.Split(line, ":")
			freq, err := strconv.ParseFloat(strings.TrimSpace(parts[1]), 32)
			if err != nil {
				log.Print("ERRRRORRRRR")
				return cpuInfo, err
			}
			totalFreq += float32(freq)
		}
	}
	cpuInfo.NumCores = processorsNum
	cpuInfo.CpuMHZ = totalFreq / float32(cpuInfo.NumCores)
	return
}
