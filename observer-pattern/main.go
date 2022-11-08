package main

func main() {
	n := notifier{
		observers: map[Observer]struct{}{},
	}

	n.Register(&observers{1})
	n.Register(&observers{2})

	n.Notify(Event{Data: 10})
	n.Notify(Event{Data: 100})
}
