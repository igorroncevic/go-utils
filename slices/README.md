# Slices

## Map

`Map` mutates every element in the `slice` using the mapper func and returns the mutated `slice`.

```
numbers := []int{1, 2, 3}
mappedNumbers := Map(numbers, func(elem int) { return a + 1 }) // returns [2, 3, 4]
```

---

## Filter

`Filter` filters out elements that do not satisfy the requirement specified in the `filtrator` func and returns the filtrated `slice`.

```
allPositive := []int{1, 2, 3}
result := Filter(allPositive, func(elem int){ return elem > 0 }) // returns [1, 2, 3]

allPositiveButOne := []int{-10, 6, 14}
result := Filter(allPositiveButOne, func(elem int){ return elem > 0 }) // returns [6, 14]
```

---

## Reduce

`Reduce` executes a `reducer` func on each element of the `slice`, where the return value is only one element.

```
numbers := []int{1, 2, 3}
result := Reduce(numbers, func(acc, elem int) { return acc + elem }, 0) // returns 6
```

---

## Find

`Find` attempts to find the specified element in the `slice`.

- If the element is found, it will be returned,
- Otherwise, `ErrNotFound` is returned.

```
words := []string{"somebody", "that i", "used to know"}
found, err := Find(words, "somebody") // returns '"somebody", nil'
```

---

## Every

`Every` checks whether all elements of the `slice` satisfy the requirement specified in the `filtrator` func.

```
allPositive := []int{1, 2, 3}
isEvery := Every(allPositive, func(elem int){ return elem > 0 }) // returns true

allPositiveButOne := []int{-10, 6, 14}
wontBeEvery := Every(allPositiveButOne, func(elem int){ return elem > 0 }) // returns false
```

---

## Some

`Some` checks whether any of elements of the `slice` satisfy the requirement specified in the `filtrator` func.

```
onlyOnePositive := []int{-100, -99, 1}
isSome := Some(onlyOnePositive, func(elem int){ return elem > 0 }) // returns true

allNegative := []int{-10, -84, -36}
wontBeAny := Some(allNegative, func(elem int){ return elem > 0 }) // returns false
```

---
