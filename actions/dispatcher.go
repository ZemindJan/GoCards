package actions

type Dispatcher[T Action] interface {
	Dispatch(action T)
	SubscribeOnStart(listener Listener[T])
	UnsubscribeOnStart(Listener Listener[T])
	SubscribeOnFinished(listener Listener[T])
	UnsubscribeOnFinished(Listener Listener[T])
	SubscribeOnBlocked(listener Listener[T])
	UnsubscribeOnBlocked(Listener Listener[T])
}
