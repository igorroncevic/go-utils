# Hashmap

## Quick explanation

Requirement is that the capacity is to the power of 2 - [SO explanation](https://stackoverflow.com/questions/53526790/why-are-hashmaps-implemented-using-powers-of-two).

`"If the size is a power of two, the keys will be more evenly distributed across
the array with minimal collision leading to better retrieval performance
when compared with any other size which is not a power of 2."`

**Note**: same key always yields the same hash and return value of `getIndex()`, therefore
the we are putting in the map will always end up in either:

- the exact returned index (e.g. `3`)
- the nearest empty index to it (e.g. `5`, which we will iterate from the returned index)

### Example

`[{"foo": 1}, {"bar": 2}, nil, nil]`

`hashmap.Put("baz", 42)`

- `getIndex("baz")` --> index = 1 (`010101101001011 (hash) & 000000000110 (capacity - 1) == 00000000010 => 2 (index)`)
- try putting the value at `hashmap[1]`
  - `hashmap[1].filled == true` --> skip and index++
- try putting the value at `hashmap[2]`
  - `hashmap[2].filled == false` --> put the value there

Example proves that even with millions of values, we can skip a substantial amount of them
and place the value in the appropriate spot in the smallest amount of time.
