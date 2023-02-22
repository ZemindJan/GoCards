package actions

import (
	"GoCards/utils"
	"unsafe"
)

type Dispatcher[T Action] interface {
	Dispatch(action T)
	SubscribeOnStart(listener Listener[T])
	UnsubscribeOnStart(Listener Listener[T])
	SubscribeOnFinished(listener Listener[T])
	UnsubscribeOnFinished(Listener Listener[T])
	SubscribeOnBlocked(listener Listener[T])
	UnsubscribeOnBlocked(Listener Listener[T])
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
