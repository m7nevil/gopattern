package behavior
type EventListener interface {
	Update(data interface{})
}


type EventManager struct {
	listeners map[string][]EventListener
}

func NewEventManager() *EventManager {
	ls := make(map[string][]EventListener)
	return &EventManager{
		listeners: ls,
	}
}

func (m *EventManager) subscribe(eventType string, listener EventListener)  {
	listeners := m.listeners[eventType]
	if listeners == nil {
		m.listeners[eventType] = []EventListener{}
	}
	m.listeners[eventType] = append(listeners, listener)
}

func (m *EventManager) unsubscribe(eventType string)  {
	delete(m.listeners, eventType)
}

func (m *EventManager) notify(eventType string, data interface{})  {
	listeners := m.listeners[eventType]
	for _, listener := range listeners{
		listener.Update(data)
	}
}