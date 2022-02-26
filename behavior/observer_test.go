package behavior

type EventManager struct {
	listeners map[string]EventListener
}

type EventListener interface {
	Update(filename string)
}

func (m *EventManager) subscribe(eventType string, listener EventListener)  {
	m.listeners[eventType] = listener
}

func (m *EventManager) unsubscribe(eventType string)  {
	delete(m.listeners, eventType)
}