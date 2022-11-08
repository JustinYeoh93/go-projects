package main

type notifier struct {
	observers map[Observer]struct{}
}

func (n *notifier) Register(o Observer) {
	n.observers[o] = struct{}{}
}

func (n *notifier) Unregister(o Observer) {
	delete(n.observers, o)
}

func (n *notifier) Notify(e Event) {
	for observer := range n.observers {
		observer.OnNotify(e)
	}
}
