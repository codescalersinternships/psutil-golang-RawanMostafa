// psutils package implements a human-friendly lib for querying processes, memory info and cpu info
package psutils

import (
	"os"
	"strconv"
	"strings"
)

type Loader interface {
	load(filePath string) (string, error)
}

type realCpuLoader struct{}

func (l *realCpuLoader) load(filePath string) (string, error) {
	return loadFile(filePath)
}

// CpuInfo holds some details about the CPU like number of cores, vendorID, model name, cache size, average CPU frequency
type CpuInfo struct {
	NumCores  int
	VendorId  string
	ModelName string
	CacheSize string
	CpuMHZ    float32
}

func loadFile(filepath string) (fileData string, err error) {
	file, err := os.ReadFile(filepath)
	fileData = string(file)

	return
}

func setCpuInfo(cpuData string) (cpuInfo CpuInfo, err error) {
	var processorsNum int
	var totalFreq float32
	lines := strings.SplitN(cpuData, "\n", -1)
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
				return cpuInfo, err
			}
			totalFreq += float32(freq)
		}
	}
	cpuInfo.NumCores = processorsNum
	cpuInfo.CpuMHZ = totalFreq / float32(cpuInfo.NumCores)
	return
}

// GetCpuInfo returns a CpuInfo type about the CPU in this moment
func GetCpuInfo() (cpuInfo CpuInfo, err error) {
	var _ Loader = (*realCpuLoader)(nil)
	return getCpuInfo(&realCpuLoader{})
}

func getCpuInfo(loader Loader) (cpuInfo CpuInfo, err error) {

	cpuData, err := loader.load("/proc/cpuinfo")
	if err != nil {
		return
	}
	cpuInfo, err = setCpuInfo(cpuData)
	return
}
