package actions

import (
	"GoCards/utils"
	"unsafe"
)

type IDispatcher[T Action] interface {
	Dispatch(action T)
	SubscribeOnStart(listener Listener[T])
	UnsubscribeOnStart(Listener Listener[T])
	SubscribeOnFinished(listener Listener[T])
	UnsubscribeOnFinished(Listener Listener[T])
	SubscribeOnBlocked(listener Listener[T])
	UnsubscribeOnBlocked(Listener Listener[T])
}

type Dispatcher[T Action] struct {
	onStartListeners    []Listener[T]
	onFinishedListeners []Listener[T]
	onBlockedListeners  []Listener[T]
}

// PUBLIC

func (dispatcher *Dispatcher[T]) Dispatch(action T) {
	dispatcher.dispatchStart(action)

	if action.IsBlocked() {
		dispatcher.dispatchBlocked(action)
	} else {
		dispatcher.dispatchFinished(action)
	}
}

func (dispatcher *Dispatcher[T]) SubscribeOnStart(listener Listener[T]) {
	dispatcher.onStartListeners = append(dispatcher.onStartListeners, listener)
}

func (dispatcher *Dispatcher[T]) SubscribeOnFinished(listener Listener[T]) {
	dispatcher.onFinishedListeners = append(dispatcher.onFinishedListeners, listener)
}

func (dispatcher *Dispatcher[T]) SubscribeOnBlocked(listener Listener[T]) {
	dispatcher.onBlockedListeners = append(dispatcher.onBlockedListeners, listener)
}

func (dispatcher *Dispatcher[T]) UnsubscribeOnStart(listener Listener[T]) {
	utils.RemoveFirst(dispatcher.onStartListeners, func(item Listener[T]) bool {
		return listener.getId() == item.getId()
	})
}

func (dispatcher *Dispatcher[T]) UnsubscribeOnFinished(listener Listener[T]) {
	utils.RemoveFirst(dispatcher.onFinishedListeners, func(item Listener[T]) bool {
		return listener.getId() == item.getId()
	})
}

func (dispatcher *Dispatcher[T]) UnsubscribeOnBlocked(listener Listener[T]) {
	utils.RemoveFirst(dispatcher.onBlockedListeners, func(item Listener[T]) bool {
		return listener.getId() == item.getId()
	})
}

// PRIVATE

func (dispatcher *Dispatcher[T]) dispatchBlocked(action T) {
	for _, listener := range dispatcher.onBlockedListeners {
		listener.Notify(action)
	}
}

func (dispatcher *Dispatcher[T]) dispatchFinished(action T) {
	for _, listener := range dispatcher.onFinishedListeners {
		listener.Notify(action)
	}
}

func (dispatcher *Dispatcher[T]) dispatchStart(action T) {
	encountered := make([]unsafe.Pointer, 0)

	for {
		for _, listener := range dispatcher.onStartListeners {
			// Don't notify encountered listeners
			if utils.Contains(encountered, unsafe.Pointer(&listener)) {
				continue
			}

			action.SetModified(false)

			listener.Notify(action)

			if action.IsModified() {
				encountered = append(encountered, unsafe.Pointer(&listener))
				break
			}

		}

		// Only stop once actions aren't modified
		if !action.IsModified() {
			break
		} else {
			action.SetModified(false)
		}
	}
}
