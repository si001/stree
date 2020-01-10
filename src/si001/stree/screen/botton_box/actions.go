package botton_box

type Action struct {
	ActName  string
	ActKey   string
	Callback func()
}

//type Action interface {
//	Name() string
//	Key() string
//	Doing() func()
//}

func (a Action) Name() string {
	return a.ActName
}

func (a Action) Key() string {
	return a.ActKey
}

func (a Action) Doing() func() {
	return a.Callback
}
