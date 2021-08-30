package safety

import "sync"

// Slice is a thread-safe []interface{} implementation. Order is not
// guaranteed.
type Slice struct {
	sync.RWMutex
	slice []interface{}
}

// NewSlice will return a pointer to a new Slice instance.
func NewSlice() *Slice {
	return &Slice{}
}

// Append will append a new entry to the slice.
func (a *Slice) Append(v interface{}) {
	a.Lock()
	defer a.Unlock()

	a.slice = append(a.slice, v)
}

// Clear will delete all slice entries.
func (a *Slice) Clear() {
	a.Lock()
	defer a.Unlock()

	a.slice = []interface{}{}
}

// Delete will delete a slice entry and return the deleted entry.
func (a *Slice) Delete(i int) (v interface{}) {
	a.Lock()
	defer a.Unlock()

	if (i >= 0) && (i < len(a.slice)) {
		v = a.slice[i]
		a.slice[i] = a.slice[len(a.slice)-1]
		a.slice = a.slice[:len(a.slice)-1]
	}

	return
}

// Get will return a slice entry.
func (a *Slice) Get(i int) (v interface{}) {
	a.RLock()
	defer a.RUnlock()

	if (i >= 0) && (i < len(a.slice)) {
		v = a.slice[i]
	}

	return
}

// Length will return the length of the slice. There is no guarantee
// that it will remain accurate for any amount of time.
func (a *Slice) Length() int {
	a.RLock()
	defer a.RUnlock()

	return len(a.slice)
}

// Range will loop over the slice and run the specified function for
// each entry. The return value determines whether or not to break the
// loop. You should not add or delete entries within Range, and you
// should avoid calling other Slice functions or you may cause
// deadlock. Range should be safe to nest for any read operations.
func (a *Slice) Range(f func(idx int, val interface{}) bool) {
	a.RLock()
	defer a.RUnlock()

	for i, v := range a.slice {
		if f(i, v) {
			break
		}
	}
}

// RangeChange will loop over the slice and run the specified function
// for each entry, storing the first return value as the new entry
// value. The second return value determines whether or not to break
// the loop. You should not add or delete entries within RangeUpdate,
// and you should avoid calling other Slice functions or you may cause
// deadlock. Range should be safe to nest for any read operations.
func (a *Slice) RangeChange(
	f func(idx int, val interface{}) (interface{}, bool),
) {
	var stop bool

	a.Lock()
	defer a.Unlock()

	for i, v := range a.slice {
		a.slice[i], stop = f(i, v)

		if stop {
			break
		}
	}
}
