package binding

// DataList is the base interface for all bindable data lists.
type DataList interface {
	AddListener(DataListListener)
	RemoveListener(DataListListener)
	GetItem(int) DataItem
	Length() int
}

// DataListListener is any object that can register for changes in a bindable DataItem.
// See NewDataListListener to define a new listener using just an inline function.
type DataListListener interface {
	DataChanged()
}

// NewDataListListener is a helper function that creates a new listener type from a list callback function.
func NewDataListListener(fn func()) DataListListener {
	return &listListener{fn}
}

type listListener struct {
	callback func()
}

func (l *listListener) DataChanged() {
	l.callback()
}

type listBase struct {
	listeners []DataListListener
	val       []DataItem
}

// AddListener allows a data listener to be informed of changes to this list.
func (b *listBase) AddListener(l DataListListener) {
	b.listeners = append(b.listeners, l)
}

// GetItem returns the dataitem at the specified index.
func (b *listBase) GetItem(i int) DataItem {
	if i >= len(b.val) {
		return nil
	}

	return b.val[i]
}

// Length returns the number of items in this data list.
func (b *listBase) Length() int {
	return len(b.val)
}

// RemoveListener should be called if the listener is no longer interested in being informed of data change events.
func (b *listBase) RemoveListener(l DataListListener) {
	for i, listen := range b.listeners {
		if listen != l {
			continue
		}

		if i == len(b.listeners)-1 {
			b.listeners = b.listeners[:len(b.listeners)-1]
		} else {
			b.listeners = append(b.listeners[:i], b.listeners[i+1:]...)
		}
	}
}

func (b *listBase) appendItem(i DataItem) {
	b.val = append(b.val, i)

	b.trigger()
}

func (b *listBase) trigger() {
	for _, listen := range b.listeners {
		queueItem(listen.DataChanged)
	}
}
