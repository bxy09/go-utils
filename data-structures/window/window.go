package indicator

type Type float64

type TypeWindow struct {
	values       []Type
	size         int
	index        int
	lastPopValid bool
	lastPop      Type
}

func New(capacity int) *TypeWindow {
	if capacity < 1 {
		return nil
	}
	return &TypeWindow{
		values: make([]Type, capacity),
	}
}

func (w *TypeWindow) Push(value Type) Type {
	if w == nil {
		var defaultValue Type
		return defaultValue
	}
	capacity := len(w.values)
	if w.size < capacity {
		w.size++
		w.lastPopValid = false
	} else {
		w.lastPopValid = true
		w.lastPop = w.values[w.index]
	}
	w.values[w.index] = value
	w.index = (w.index + 1) % capacity
	return w.lastPop
}

func (w *TypeWindow) UnPush() (Type, bool) {
	if w != nil && w.size > 0 {
		w.index = (len(w.values) + w.index - 1) % len(w.values)
		tmp := w.values[w.index]
		if w.lastPopValid {
			w.values[w.index] = w.lastPop
			w.lastPopValid = false
		} else {
			w.size--
		}
		return tmp, true
	}
	var defaultValue Type
	return defaultValue, false
}

func (w *TypeWindow) Foreach(f func(int, Type)) {
	capacity := len(w.values)
	head := (w.index - w.size + capacity) % capacity
	for i := 0; i < w.size; i++ {
		f(i, w.values[(head+i)%capacity])
	}
}

func (w *TypeWindow) Head() int {
	capacity := len(w.values)
	return (w.index - w.size + capacity) % capacity
}

func (w *TypeWindow) Len() int {
	return w.size
}

func (w *TypeWindow) Capacity() int {
	return len(w.values)
}

func (w *TypeWindow) Raw() []Type {
	return w.values
}
