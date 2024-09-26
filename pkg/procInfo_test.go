package psutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var procStatus string = `Name:	Anything
Umask:	0000
State:	R (Running)
Tgid:	4085954
Ngid:	0
Pid:	5
PPid:	2
TracerPid:	0
Uid:	0	0	0	0
Gid:	0	0	0	0
FDSize:	64`

type spyProcLoader struct{}

func (l *spyProcLoader) Load(string) (string, error) {
	return procStatus, nil
}
func TestGetProcList(t *testing.T) {
	var _ Loader = (*spyProcLoader)(nil)
	got, err := getProcessList(&spyProcLoader{})
	if err != nil {
		t.Errorf("%v", err)
	}
	expected := Process{
		PID:         5,
		ProcessName: "Anything",
	}
	assert.Equal(t, expected, got[0])

}
