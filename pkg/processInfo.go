// psutils package implements a human-friendly lib for querying processes, memory info and cpu info
package psutils

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Process holds the basic information about a process, like its Pid and name
type Process struct {
	PID         int
	ProcessName string
}

// ProcessDetails holds more detailed information about a process, like its State, PPid, Tgid
type ProcessDetails struct {
	State string
	PPID  int
	Tgid  int
}

type realProcLoader struct{}

func (l *realProcLoader) load(filePath string) (string, error) {
	return loadFile(filePath)
}

func setProcInfo(data string) (proc Process, procDetails ProcessDetails, err error) {
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
			procDetails.State = value
		} else if strings.HasPrefix(line, "PPid") {

			parts := strings.Split(line, ":")
			value := strings.TrimSpace(parts[1])
			procDetails.PPID, err = strconv.Atoi(value)
		} else if strings.HasPrefix(line, "Tgid") {

			parts := strings.Split(line, ":")
			value := strings.TrimSpace(parts[1])
			procDetails.Tgid, err = strconv.Atoi(value)
		}
	}
	return
}

// GetProcessList returns a list of all currently running processes and an error if exists
func GetProcessList() (procs []Process, err error) {
	var _ Loader = (*realProcLoader)(nil)
	return getProcessList(&realProcLoader{})
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
		data, err = loader.load(statusFile)
		if err != nil {
			return
		}
		var process Process
		var procDetails ProcessDetails
		process, procDetails, err = setProcInfo(data)
		if strings.Contains(procDetails.State, "R") {
			procs = append(procs, process)
		}
	}
	return
}

// GetProcessDetails takes a PID as an argument
// and returns its details as ProcessDetails type, and an error if exists
func GetProcessDetails(PID int) (ProcessDetails ProcessDetails, err error) {
	var _ Loader = (*realProcLoader)(nil)
	return getProcessDetails(PID, &realProcLoader{})
}

func getProcessDetails(PID int, loader Loader) (ProcessDetails ProcessDetails, err error) {

	statusFile := filepath.Join("/proc", fmt.Sprint(PID), "status")
	var data string
	data, err = loader.load(statusFile)
	if err != nil {
		return
	}
	_, ProcessDetails, err = setProcInfo(data)

	return
}
