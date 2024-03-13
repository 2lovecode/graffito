package pattern

type Event struct {
	Data interface{}
}

type IObserver interface {
	OnNotify(eve Event)
}

type INotifier interface {
	Register(obs IObserver)
	Deregister(obs IObserver)
	Notify(eve Event)
}

type Observer struct {
	id   int
	Data interface{}
}

type Notifier struct {
	observers map[IObserver]struct{}
}

func (o *Observer) OnNotify(eve Event) {
	o.Data = eve.Data
}

func (n *Notifier) Register(obs IObserver) {
	n.observers[obs] = struct{}{}
}

func (n *Notifier) Deregister(obs IObserver) {
	delete(n.observers, obs)
}

func (n *Notifier) Notify(eve Event) {
	for o := range n.observers {
		o.OnNotify(eve)
	}
}

func NewNotifier() Notifier {
	return Notifier{observers: map[IObserver]struct{}{}}
}

func NewObserver(id int) *Observer {
	return &Observer{id: id}
}
