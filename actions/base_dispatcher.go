package actions

import (
	"GoCards/utils"
)

type BaseDispatcher[T Action] struct {
	onStartListeners    []Listener[T]
	onFinishedListeners []Listener[T]
	onBlockedListeners  []Listener[T]
}

// PUBLIC

func (dispatcher *BaseDispatcher[T]) Dispatch(action T) {
	dispatcher.dispatchStart(action)

	if action.IsBlocked() {
		dispatcher.dispatchBlocked(action)
	} else {
		dispatcher.dispatchFinished(action)
	}
}

func (dispatcher *BaseDispatcher[T]) SubscribeOnStart(listener Listener[T]) {
	dispatcher.onStartListeners = append(dispatcher.onStartListeners, listener)
}

func (dispatcher *BaseDispatcher[T]) SubscribeOnFinished(listener Listener[T]) {
	dispatcher.onFinishedListeners = append(dispatcher.onFinishedListeners, listener)
}

func (dispatcher *BaseDispatcher[T]) SubscribeOnBlocked(listener Listener[T]) {
	dispatcher.onBlockedListeners = append(dispatcher.onBlockedListeners, listener)
}

func (dispatcher *BaseDispatcher[T]) UnsubscribeOnStart(listener Listener[T]) {
	utils.RemoveFirst(dispatcher.onStartListeners, func(item Listener[T]) bool {
		return listener.GetId() == item.GetId()
	})
}

func (dispatcher *BaseDispatcher[T]) UnsubscribeOnFinished(listener Listener[T]) {
	utils.RemoveFirst(dispatcher.onFinishedListeners, func(item Listener[T]) bool {
		return listener.GetId() == item.GetId()
	})
}

func (dispatcher *BaseDispatcher[T]) UnsubscribeOnBlocked(listener Listener[T]) {
	utils.RemoveFirst(dispatcher.onBlockedListeners, func(item Listener[T]) bool {
		return listener.GetId() == item.GetId()
	})
}

func NewBaseDispatcher[T Action]() *BaseDispatcher[T] {
	return &BaseDispatcher[T]{
		onStartListeners:    make([]Listener[T], 0),
		onFinishedListeners: make([]Listener[T], 0),
		onBlockedListeners:  make([]Listener[T], 0),
	}
}

// PRIVATE

func (dispatcher *BaseDispatcher[T]) dispatchBlocked(action T) {
	for _, listener := range dispatcher.onBlockedListeners {
		listener.Notify(action)
	}
}

func (dispatcher *BaseDispatcher[T]) dispatchFinished(action T) {
	for _, listener := range dispatcher.onFinishedListeners {
		listener.Notify(action)
	}
}

func (dispatcher *BaseDispatcher[T]) dispatchStart(action T) {
	encountered := make([]ID, 0)

	for {
		for _, listener := range dispatcher.onStartListeners {
			// Don't notify encountered listeners
			if utils.Contains(encountered, listener.GetId()) {
				continue
			}

			action.SetModified(false)

			listener.Notify(action)

			if action.IsModified() {
				encountered = append(encountered, listener.GetId())
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
