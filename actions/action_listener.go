package actions

type ActionListener[T Action] interface {
	Notify(action T)
}

type ActionListenerFunction[T Action] struct {
	Function func(action T)
}

func (listener ActionListenerFunction[T]) Notify(action T) {
	listener.Function(action)
}
