package psutils

import (
	"strconv"
	"strings"
)

type MemInfo struct {
	TotalMemoryKB     float64
	UsedMemoryKB      float64
	AvailableMemoryKB float64
}

type RealMemLoader struct{}

func (l *RealMemLoader) Load() (string, error) {
	return loadFile("/proc/meminfo")
}

func extractValue(line string) (valueNoKB string) {
	parts := strings.Split(line, ":")
	value := strings.Replace(parts[1], "kB", "", 1)
	valueNoKB = strings.TrimSpace(value)
	return
}

func setMemInfo(memData string) (memInfo MemInfo, err error) {
	lines := strings.SplitN(memData, "\n", -1)

	for _, line := range lines {
		if strings.HasPrefix(line, "MemTotal") {

			memTotalNoKB := extractValue(line)
			memInfo.TotalMemoryKB, err = strconv.ParseFloat(memTotalNoKB, 64)

		} else if strings.HasPrefix(line, "MemAvailable") {

			memavailabeNoKB := extractValue(line)
			memInfo.AvailableMemoryKB, err = strconv.ParseFloat(memavailabeNoKB, 64)

		} else if strings.HasPrefix(line, "MemFree") {

			var free float64
			MemFreeNoKB := extractValue(line)
			free, err = strconv.ParseFloat(MemFreeNoKB, 64)
			memInfo.UsedMemoryKB = memInfo.TotalMemoryKB - free
		}
	}
	return
}

func GetMemInfo() (memInfo MemInfo, err error) {
	var _ Loader = (*RealMemLoader)(nil)
	return getMemInfo(&RealMemLoader{})
}

func getMemInfo(loader Loader) (memInfo MemInfo, err error) {

	memData, err := loader.Load()
	if err != nil {
		return
	}
	memInfo, err = setMemInfo(memData)
	return
}
