package leetcode

import (
	"errors"
	"math"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

const (
	machineID          = int64(1)
	base         int64 = 62
	characterSet       = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

type Codec struct {
	sync.Mutex
	seq  int64
	base map[string]string
}

func NewCodec() Codec {
	return Codec{
		base: make(map[string]string),
		seq:  0,
	}

}

func (c *Codec) Encode(longUrl string) string {
	t := c.nextID()
	short := c.toBase62(t)
	c.Lock()
	defer c.Unlock()
	c.base[short] = longUrl
	return short
}

func (c *Codec) Decode(shortUrl string) string {
	return c.base[shortUrl]

}

func (Codec) toBase62(num int64) string {
	encoded := ""
	for num > 0 {
		r := num % base
		num /= base
		encoded = string(characterSet[r]) + encoded

	}
	return encoded
}

func (c *Codec) nextID() int64 {
	startTime := time.Date(2022, 11, 5, 0, 0, 0, 0, time.UTC)
	now := time.Now().UTC().UnixNano() / 1e6

	fromEpoch := now - startTime.UTC().UnixNano()/1e6

	var seq int64
	if c.seq > 1<<12-1 {
		seq = atomic.SwapInt64(&c.seq, 0)
	} else {
		seq = atomic.AddInt64(&c.seq, 1)
	}

	id := fromEpoch<<22 | machineID<<12 | seq
	return id
}

func (Codec) fromBase62(encoded string) (uint64, error) {
	var val uint64
	for i, char := range encoded {
		pow := len(encoded) - (i + 1)
		pos := strings.IndexRune(characterSet, char)
		if pos == -1 {
			return 0, errors.New("invalid character: " + string(char))
		}

		val += uint64(pos) * uint64(math.Pow(float64(base), float64(pow)))
	}

	return val, nil
}
