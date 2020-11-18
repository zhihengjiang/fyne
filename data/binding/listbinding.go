package binding

// DataList is the base interface for all bindable data lists.
type DataList interface {
	DataItem
	GetItem(int) DataItem
	Length() int
}

type listListener struct {
	callback func()
}

func (l *listListener) DataChanged() {
	l.callback()
}

type listBase struct {
	base
	val []DataItem
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

func (b *listBase) appendItem(i DataItem) {
	b.val = append(b.val, i)

	b.trigger()
}

func (b *listBase) trigger() {
	for _, listen := range b.listeners {
		queueItem(listen.DataChanged)
	}
}
