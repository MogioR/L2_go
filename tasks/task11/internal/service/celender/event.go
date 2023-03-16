package celender

import (
	"sync"
	"time"

	celenderdomain "task11/internal/domain"

	"github.com/google/uuid"
)

type Event struct {
	uid   string
	date  time.Time
	title string
	m     sync.RWMutex
}

func NewEventByDomain(event celenderdomain.Event) *Event {
	date, _ := time.Parse("2006-01-02", event.Date)
	return &Event{
		uid:   event.Uid,
		date:  date,
		title: event.Title,
		m:     sync.RWMutex{},
	}
}

func NewEvent(date time.Time, title string) *Event {
	return &Event{
		uid:   uuid.New().String(),
		date:  date,
		title: title,
	}
}

func (e *Event) EventToDomain() *celenderdomain.Event {
	return &celenderdomain.Event{
		Uid:   e.uid,
		Date:  e.date.Format("2006-01-02"),
		Title: e.title,
	}
}

func (e *Event) GetUid() string {
	e.m.RLock()
	defer e.m.RUnlock()
	return e.uid
}

func (e *Event) GetDate() time.Time {
	e.m.RLock()
	defer e.m.RUnlock()
	return e.date
}

func (e *Event) GetTitle() string {
	e.m.RLock()
	defer e.m.RUnlock()
	return e.title
}

func (e *Event) setDate(date time.Time) {
	e.m.Lock()
	e.date = date
	e.m.Unlock()
}

func (e *Event) setTitle(title string) {
	e.m.Lock()
	e.title = title
	e.m.Unlock()
}
