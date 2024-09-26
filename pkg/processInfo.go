package psutils

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Process struct {
	PID         int
	ProcessName string
}

type RealProcLoader struct{}

func (l *RealProcLoader) Load(filePath string) (string, error) {
	return loadFile(filePath)
}

func setProcInfo(data string) (proc Process, running bool, err error) {
	lines := strings.SplitN(data, "\n", -1)

	for _, line := range lines {
		if strings.HasPrefix(line, "Name") {
			parts := strings.Split(line, ":")
			value := strings.TrimSpace(parts[1])
			proc.ProcessName = value

		} else if strings.HasPrefix(line, "Pid") {

			parts := strings.Split(line, ":")
			value := strings.TrimSpace(parts[1])
			proc.PID, err = strconv.Atoi(value)
		} else if strings.HasPrefix(line, "State") {

			parts := strings.Split(line, ":")
			value := strings.TrimSpace(parts[1])
			if strings.Contains(value, "R") {
				running = true
			}
		}
	}
	return
}

func GetProcessList() (procs []Process, err error) {
	var _ Loader = (*RealProcLoader)(nil)
	return getProcessList(&RealProcLoader{})
}

func getProcessList(loader Loader) (procs []Process, err error) {
	procs = make([]Process, 0)
	dir, err := os.Open("/proc")
	if err != nil {
		return
	}
	defer dir.Close()

	entries, err := dir.Readdirnames(0)
	if err != nil {
		return
	}

	for _, entry := range entries {

		_, err = strconv.Atoi(entry)
		if err != nil {
			continue
		}

		statusFile := filepath.Join("/proc", entry, "status")
		var data string
		data, err = loader.Load(statusFile)
		if err != nil {
			return
		}
		var process Process
		var running bool
		process, running, err = setProcInfo(data)
		if running {
			procs = append(procs, process)
		}
	}
	return
}
