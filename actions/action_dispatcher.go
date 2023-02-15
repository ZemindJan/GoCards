package actions

import (
	"GoCards/slice"
)

type ActionDispatcher[T Action] struct {
	onStartListeners     []*ActionListener[T] // These are the only listeners which may modify
	onCompletedListeners []*ActionListener[T]
	onBlockedListeners   []*ActionListener[T]
}

func (dispatcher ActionDispatcher[T]) Dispatch(action T) {
	var encountered_listeners = make([]*ActionListener[T], 0)

	for {
		action.SetModified(false)

		for _, ptr := range dispatcher.onStartListeners {
			if ptr == nil {
				return
			}

			listener := (*ptr)

			// Skip Encountered Listeners
			if slice.Contains(encountered_listeners, ptr) {
				continue
			}

			listener.Notify(action)

			if action.IsModified() {
				encountered_listeners = append(encountered_listeners, ptr)
				break
			}
		}

		if !action.IsModified() {
			break
		}
	}

	var final_listeners = dispatcher.onCompletedListeners

	if action.IsBlocked() {
		final_listeners = dispatcher.onBlockedListeners
	}

	if final_listeners == nil {
		return
	}

	for _, ptr := range final_listeners {
		if ptr == nil {
			return
		}

		var listener = *ptr
		listener.Notify(action)
	}
}

func (dispatcher ActionDispatcher[T]) OnStart(listener *ActionListener[T]) {
	dispatcher.onStartListeners = append(dispatcher.onStartListeners, listener)
}

func (dispatcher ActionDispatcher[T]) OnCompleted(listener *ActionListener[T]) {
	dispatcher.onCompletedListeners = append(dispatcher.onStartListeners, listener)
}

func (dispatcher ActionDispatcher[T]) OnBlocked(listener *ActionListener[T]) {
	dispatcher.onBlockedListeners = append(dispatcher.onStartListeners, listener)
}
