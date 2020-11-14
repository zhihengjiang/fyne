package binding

// FloatList supports binding a list of string values in a Fyne application
type FloatList interface {
	DataList

	Append(float64)
	Get(int) float64
}

// NewFloatList returns a bindable list of string values.
func NewFloatList() FloatList {
	return &floatListBind{}
}

type floatListBind struct {
	listBase
}

func (f *floatListBind) Append(val float64) {
	f.appendItem(BindFloat(&val))
}

func (f *floatListBind) Get(i int) float64 {
	if i > f.Length() {
		return 0.0
	}

	return f.GetItem(i).(Float).Get()
}
