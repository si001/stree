package botton_box

type Action struct {
	ActName  string
	ActKey   string
	Callback func()
}

func (a Action) Name() string {
	return a.ActName
}

func (a Action) Key() string {
	return a.ActKey
}

func (a Action) Doing() func() {
	return a.Callback
}
