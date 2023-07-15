package hashmap_test

import (
	"fmt"
	"math/rand"
	"testing"

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
	stdmap := make(map[uint64]uint32)
	hmap := hashmap.New[uint64, uint32](uint64(rand.Intn(1024)), util.Equals[uint64], util.HashUint64)

	const nops = 1000

	for i := 0; i < nops; i++ {
		key := uint64(rand.Intn(100))
		val := rand.Uint32()
		op := rand.Intn(2)

		switch op {
		case 0:
			stdmap[key] = val
			hmap.Put(key, val)
		case 1:
			var del uint64
			for k := range stdmap {
				del = k
				break
			}
			delete(stdmap, del)
			hmap.Remove(del)
		}

		checkeq(hmap, func(k uint64) (uint32, bool) {
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
	// Output:
	// 42 true
	// 0 false
	// 0 false
	// 13 true
	// 0 false
	// 0 false
}
