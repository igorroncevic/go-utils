package hashmap_test

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/igorroncevic/go-utils/hashmap"
	"github.com/igorroncevic/go-utils/util"
)

// Check if value from hashmap exists in standard map using the provided get function.
func checkeq[K any, V comparable](hmap *hashmap.Map[K, V], get func(k K) (V, bool), t *testing.T) {
	hmap.Each(func(key K, val V) {
		if ov, ok := get(key); !ok {
			t.Fatalf("key %v should exist", key)
		} else if val != ov {
			t.Fatalf("value mismatch: %v != %v", val, ov)
		}
	})
}

func TestHashmapVsStandardMapCrossCheck(t *testing.T) {
	randSize, err := util.RandomInt64(1024)
	assert.NoError(t, err, "error generating random hmap size")
	assert.NotEqual(t, -1, randSize, "randSize is empty")

	stdmap := make(map[int64]int64)
	hmap := hashmap.New[int64, int64](uint64(randSize), util.Equals[int64], util.HashInt64)

	const nops = 1000

	for i := 0; i < nops; i++ {
		key, err := util.RandomInt64(1000)
		assert.NoError(t, err, "error generating key")
		assert.NotEqual(t, -1, key, "key is empty")

		val, err := util.RandomInt64(100000)
		assert.NoError(t, err, "error generating value")
		assert.NotEqual(t, -1, val, "val is empty")

		op, err := util.RandomInt64(2)
		assert.NoError(t, err, "error generating op")
		assert.NotEqual(t, -1, op, "op is empty")

		switch op {
		case 0:
			stdmap[key] = val
			hmap.Put(key, val)
		case 1:
			var del int64
			for k := range stdmap {
				del = k
				break
			}

			delete(stdmap, del)
			hmap.Remove(del)
		}

		checkeq(hmap, func(k int64) (int64, bool) {
			v, ok := stdmap[k]
			return v, ok
		}, t)
	}
}

func TestHashmapCopy(t *testing.T) {
	orig := hashmap.New[uint64, uint32](1, util.Equals[uint64], util.HashUint64)

	for i := uint32(0); i < 10; i++ {
		orig.Put(uint64(i), i)
	}

	cpy := orig.Copy()

	checkeq(cpy, orig.Get, t)

	cpy.Put(0, 42)

	if v, _ := cpy.Get(0); v != 42 {
		t.Fatal("didn't get 42")
	}
}

func TestHashmapFlow(t *testing.T) {
	hmap := hashmap.New[string, int](1, util.Equals[string], util.HashString)

	hmap.Put("foo", 42)
	hmap.Put("bar", 13)

	fmt.Println(hmap.Get("foo"))
	fmt.Println(hmap.Get("baz"))

	hmap.Remove("foo")

	fmt.Println(hmap.Get("foo"))
	fmt.Println(hmap.Get("bar"))

	hmap.Clear()

	fmt.Println(hmap.Get("foo"))
	fmt.Println(hmap.Get("bar"))
}
