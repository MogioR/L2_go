package celender

import (
	"fmt"
	"sync"
	"time"
)

type EventRepository struct {
	m    sync.RWMutex
	data map[int]map[string]*Event
}

var (
	Instanse *EventRepository
)

// Создание инстанса
func Create() {
	Instanse = &EventRepository{
		sync.RWMutex{},
		make(map[int]map[string]*Event),
	}
}

// Добавить событие
func (er *EventRepository) Add(userID int, event *Event) {
	er.m.Lock()
	if _, ok := er.data[userID]; !ok {
		er.data[userID] = make(map[string]*Event)
	}
	er.data[userID][event.GetUid()] = event
	er.m.Unlock()
}

// Удалить событие
func (er *EventRepository) Del(userID int, eventUID string) (ok bool) {
	er.m.Lock()
	if events, ok := er.data[userID]; ok {
		if _, ok = events[eventUID]; ok {
			delete(er.data[userID], eventUID)
		}
	}
	er.m.Unlock()
	return ok
}

// Получить событие
func (er *EventRepository) Get(userID int, eventUID string) *Event {
	er.m.RLock()
	defer er.m.RUnlock()
	return er.data[userID][eventUID]
}

// Обновить событие
func (er *EventRepository) Update(userID int, e *Event) (ok error) {
	currentEvent := er.Get(userID, e.GetUid())

	if currentEvent == nil {
		return fmt.Errorf("event from user %d with id with: %s - not found", userID, e.GetUid())
	}
	if e.GetTitle() != "" {
		currentEvent.setTitle(e.GetTitle())
	}
	if !e.GetDate().IsZero() {
		currentEvent.setDate(e.GetDate())
	}
	return nil
}

// Получить события на промежутке
func (er *EventRepository) GetFromTimeInterval(userID int, start, end time.Time) (ev []*Event) {
	er.m.RLock()
	result := make([]*Event, 0)
	if events, ok := er.data[userID]; ok {
		for _, event := range events {
			if (event.GetDate() == start || event.GetDate().After(start)) && event.GetDate().Before(end) {
				result = append(result, event)
			}
		}
	}
	er.m.RUnlock()
	return result
}
