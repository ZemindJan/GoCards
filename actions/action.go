package actions

// Actions are changes in the game state
// Listeners may modify them before they are executed
// Or react when they are executed
type Action interface {
	// Allows the action to be blocked from executing
	IsBlocked() bool
	SetBlocked(isBlocked bool)

	// Allows the action to be modified before executing
	// Must be tracked to disallow double modification
	IsModified() bool
	SetModified(isModified bool)

	// Makes the change in game state
	Execute()
}

// Basic implementation of blocking and modification functions for actions
// Still requires Execute() to be a fully implemented action
type ActionBase struct {
	isBlocked  bool
	isModified bool
}

func (action *ActionBase) IsBlocked() bool {
	return action.isBlocked
}

func NewActionBase() *ActionBase {
	return &ActionBase{
		isBlocked:  false,
		isModified: false,
	}
}

// This also sets isModified to true
func (action *ActionBase) SetBlocked(isBlocked bool) {
	action.isBlocked = isBlocked
	action.isModified = true
}

func (action *ActionBase) IsModified() bool {
	return action.isModified
}

func (action *ActionBase) SetModified(isModified bool) {
	action.isModified = isModified
}
