package rsync
import (
	"errors"
	"fmt"
)

type RollingChecksum struct {
	data []byte
	blockSize int

	start int
	end int

	a uint16
	b uint16

	checksum uint32;
}

func newRollingChecksum(data []byte, blockSize int) *RollingChecksum {
	r := &RollingChecksum{}

	fmt.Println("Data: ", data)

	r.data = data[:];
	r.start = 0;
	r.end = min(blockSize, len(data));
	r.blockSize = blockSize

	for i := r.start; i < r.end; i++ {
		r.a += uint16(data[i]);
		r.b += uint16(r.a);
	}
	r.a, r.b = r.a % 16, r.b % 16
	r.calcChecksum()

	return r
}

func (r *RollingChecksum) Sum() uint32 {
	return r.checksum;
}

func (r *RollingChecksum) calcChecksum() {
	r.checksum = uint32(r.a) + (MOD_VALUE * uint32(r.b))
}

func (r* RollingChecksum) Shift() (uint32, error) {
	if r.end >= len(r.data) {
		return 0, errors.New("rollingChecksum: End is greater than size of data!")
	}

	r.a -= uint16(r.data[r.start])
	r.a += uint16(r.data[r.end + 1])

	r.b -= (uint16(r.blockSize) * uint16(r.data[r.start]))
	r.b += uint16(r.a)

	r.start++
	r.end++

	r.calcChecksum()

	return r.checksum, nil
}