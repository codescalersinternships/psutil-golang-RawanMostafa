package psutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var memInfo string = `MemTotal:       16239040 kB
MemFree:         1313720 kB
MemAvailable:    8580468 kB
Buffers:          292824 kB
Cached:          7358304 kB
SwapCached:        24172 kB
Active:          6177748 kB
Inactive:        6296272 kB
`

type spyMemLoader struct{}

func (l *spyMemLoader) Load(string) (string, error) {
	return memInfo, nil
}
func TestGetMemInfo(t *testing.T) {
	var _ Loader = (*spyMemLoader)(nil)
	got, _ := getMemInfo(&spyMemLoader{})
	expected := MemInfo{
		TotalMemoryKB:     16239040,
		UsedMemoryKB:      14925320,
		AvailableMemoryKB: 8580468,
	}
	assert.Equal(t, expected, got)

}
