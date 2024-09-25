package psutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cpuinfo string = `processor	: 0
vendor_id	: GenuineIntel
cpu family	: 6
model		: 165
model name	: Intel(R) Core(TM) i7-10750H CPU @ 2.60GHz
stepping	: 2
microcode	: 0xfc
cpu MHz		: 1461.816
cache size	: 12288 KB
physical id	: 0
siblings	: 12
core id		: 0
cpu cores	: 6

processor	: 1
vendor_id	: GenuineIntel
cpu family	: 6
model		: 165
model name	: Intel(R) Core(TM) i7-10750H CPU @ 2.60GHz
stepping	: 2
microcode	: 0xfc
cpu MHz		: 1533.192
cache size	: 12288 KB
physical id	: 0
siblings	: 12
core id		: 1
cpu cores	: 6

processor	: 2
vendor_id	: GenuineIntel
cpu family	: 6
model		: 165
model name	: Intel(R) Core(TM) i7-10750H CPU @ 2.60GHz
stepping	: 2
microcode	: 0xfc
cpu MHz		: 1400.059
cache size	: 12288 KB
physical id	: 0
siblings	: 12
core id		: 2
cpu cores	: 6

processor	: 3
vendor_id	: GenuineIntel
cpu family	: 6
model		: 165
model name	: Intel(R) Core(TM) i7-10750H CPU @ 2.60GHz
stepping	: 2
microcode	: 0xfc
cpu MHz		: 1800.220
cache size	: 12288 KB
physical id	: 0
siblings	: 12
core id		: 3
cpu cores	: 6

processor	: 4
vendor_id	: GenuineIntel
cpu family	: 6
model		: 165
model name	: Intel(R) Core(TM) i7-10750H CPU @ 2.60GHz
stepping	: 2
microcode	: 0xfc
cpu MHz		: 1800.806
cache size	: 12288 KB
physical id	: 0
siblings	: 12
core id		: 4
cpu cores	: 6

processor	: 5
vendor_id	: GenuineIntel
cpu family	: 6
model		: 165
model name	: Intel(R) Core(TM) i7-10750H CPU @ 2.60GHz
stepping	: 2
microcode	: 0xfc
cpu MHz		: 1808.303
cache size	: 12288 KB
physical id	: 0
siblings	: 12
core id		: 5
cpu cores	: 6

processor	: 6
vendor_id	: GenuineIntel
cpu family	: 6
model		: 165
model name	: Intel(R) Core(TM) i7-10750H CPU @ 2.60GHz
stepping	: 2
microcode	: 0xfc
cpu MHz		: 900.094
cache size	: 12288 KB
physical id	: 0
siblings	: 12
core id		: 0
cpu cores	: 6

processor	: 7
vendor_id	: GenuineIntel
cpu family	: 6
model		: 165
model name	: Intel(R) Core(TM) i7-10750H CPU @ 2.60GHz
stepping	: 2
microcode	: 0xfc
cpu MHz		: 899.999
cache size	: 12288 KB
physical id	: 0
siblings	: 12
core id		: 1
cpu cores	: 6

processor	: 8
vendor_id	: GenuineIntel
cpu family	: 6
model		: 165
model name	: Intel(R) Core(TM) i7-10750H CPU @ 2.60GHz
stepping	: 2
microcode	: 0xfc
cpu MHz		: 800.000
cache size	: 12288 KB
physical id	: 0
siblings	: 12
core id		: 2
cpu cores	: 6

processor	: 9
vendor_id	: GenuineIntel
cpu family	: 6
model		: 165
model name	: Intel(R) Core(TM) i7-10750H CPU @ 2.60GHz
stepping	: 2
microcode	: 0xfc
cpu MHz		: 800.000
cache size	: 12288 KB
physical id	: 0
siblings	: 12
core id		: 3
cpu cores	: 6

processor	: 10
vendor_id	: GenuineIntel
cpu family	: 6
model		: 165
model name	: Intel(R) Core(TM) i7-10750H CPU @ 2.60GHz
stepping	: 2
microcode	: 0xfc
cpu MHz		: 800.000
cache size	: 12288 KB
physical id	: 0
siblings	: 12
core id		: 4
cpu cores	: 6

processor	: 11
vendor_id	: GenuineIntel
cpu family	: 6
model		: 165
model name	: Intel(R) Core(TM) i7-10750H CPU @ 2.60GHz
stepping	: 2
microcode	: 0xfc
cpu MHz		: 800.000
cache size	: 12288 KB
physical id	: 0
siblings	: 12
core id		: 5
cpu cores	: 6
`

type spyCpuLoader struct{}

func (l *spyCpuLoader) Load() (string, error) {
	return cpuinfo, nil
}
func TestGetCpuInfo(t *testing.T) {
	var _ Loader = (*spyCpuLoader)(nil)
	got, _ := getCpuInfo(&spyCpuLoader{})
	expected := CpuInfo{
		NumCores:  12,
		VendorId:  "GenuineIntel",
		ModelName: "Intel(R) Core(TM) i7-10750H CPU @ 2.60GHz",
		CacheSize: "12288 KB",
		CpuMHZ:    1233.7074,
	}
	assert.Equal(t, expected, got)

}
