package psutils

import (
	"os"
	"strconv"
	"strings"
)

type Loader interface {
	Load() (string, error)
}

type RealLoader struct{}

func (l *RealLoader) Load() (string, error) {
	return loadCpuFile()
}

type CpuInfo struct {
	NumCores  int
	VendorId  string
	ModelName string
	CacheSize string
	CpuMHZ    float32
}

func loadCpuFile() (cpuFile string, err error) {
	file, err := os.ReadFile("/proc/cpuinfo")
	cpuFile = string(file)

	return
}

func setCpuInfo(cpuFile string) (cpuInfo CpuInfo, err error) {
	var processorsNum int
	var totalFreq float32
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
				return cpuInfo, err
			}
			totalFreq += float32(freq)
		}
	}
	cpuInfo.NumCores = processorsNum
	cpuInfo.CpuMHZ = totalFreq / float32(cpuInfo.NumCores)
	return
}

func GetCpuInfo() (cpuInfo CpuInfo, err error) {
	var _ Loader = (*RealLoader)(nil)
	return getCpuInfo(&RealLoader{})
}

func getCpuInfo(loader Loader) (cpuInfo CpuInfo, err error) {

	cpuFile, err := loader.Load()
	if err != nil {
		return
	}
	cpuInfo, err = setCpuInfo(cpuFile)
	return
}
