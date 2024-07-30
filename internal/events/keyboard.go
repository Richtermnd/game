package events

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type KeyboardEventType string

const (
	KeyPressed  KeyboardEventType = "KeyPressed"
	KeyReleased KeyboardEventType = "KeyRealesed"
)

type KeyboardEvent struct {
	EventType KeyboardEventType // Not just "Type" to avoid name collision with method
	Key       ebiten.Key
}

func (ke *KeyboardEvent) Type() string {
	return string(ke.EventType) + ke.Key.String()
}

func SubscribeKeyboard(handler EventHandler, t KeyboardEventType, keys ...ebiten.Key) {
	events := make([]Event, 0, len(keys))
	for _, key := range keys {
		events = append(events, &KeyboardEvent{
			EventType: t,
			Key:       key,
		})
	}
	Subscribe(handler, events...)
}

func AcceptKeyboard() {
	const keysCapacity = 10
	keys := make([]ebiten.Key, 0, keysCapacity)
	keys = inpututil.AppendJustPressedKeys(keys)
	for _, key := range keys {
		event := &KeyboardEvent{
			EventType: KeyPressed,
			Key:       key,
		}
		AddEvent(event)
	}
	keys = keys[:0]
	keys = inpututil.AppendJustReleasedKeys(keys)
	for _, key := range keys {
		event := &KeyboardEvent{
			EventType: KeyReleased,
			Key:       key,
		}
		AddEvent(event)
	}
}
