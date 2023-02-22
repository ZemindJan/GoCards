package actions

import "GoCards/utils"

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
		return listener.getId() == item.getId()
	})
}

func (dispatcher *BaseDispatcher[T]) UnsubscribeOnFinished(listener Listener[T]) {
	utils.RemoveFirst(dispatcher.onFinishedListeners, func(item Listener[T]) bool {
		return listener.getId() == item.getId()
	})
}

func (dispatcher *BaseDispatcher[T]) UnsubscribeOnBlocked(listener Listener[T]) {
	utils.RemoveFirst(dispatcher.onBlockedListeners, func(item Listener[T]) bool {
		return listener.getId() == item.getId()
	})
}

func NewBaseDispatcher[T Action]() *BaseDispatcher[T] {
	return &BaseDispatcher[T]{
		onStartListeners:    make([]Listener[T], 0),
		onFinishedListeners: make([]Listener[T], 0),
		onBlockedListeners:  make([]Listener[T], 0),
	}
}
