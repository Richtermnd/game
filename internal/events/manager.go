package events

const eventQueueCapacity = 20

var (
	eventQueue  = make([]Event, 0, eventQueueCapacity)
	subscribers = make(map[string][]EventHandler)
)

type EventHandler func(e Event)

type Event interface {
	Type() string // Unique type for storing in eventManager map of subscribers
}

func AddEvent(event Event) {
	eventQueue = append(eventQueue, event)
}

func Notify() {
	for _, event := range eventQueue {
		subs := subscribers[event.Type()]
		for _, handler := range subs {
			handler(event)
		}
	}
	eventQueue = eventQueue[:0]
}

func Subscribe(handler EventHandler, events ...Event) {
	for _, event := range events {
		subscribers[event.Type()] = append(subscribers[event.Type()], handler)
	}
}
