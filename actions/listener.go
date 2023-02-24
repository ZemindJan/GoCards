package actions

type ID string

// Interface for listeners that want to intercept changes to game state
// getID should be unique
type Listener[T Action] interface {
	Notify(action T)
	GetId() ID
}
