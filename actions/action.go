package actions

type Action interface {
	IsBlocked() bool
	IsModified() bool
	SetModified(isModified bool)
	SetBlocked(isBlocked bool)
	Execute()
}

type BaseAction struct {
	isBlocked  bool
	isModified bool
}

func (action BaseAction) IsBlocked() bool {
	return action.isBlocked
}

func (action BaseAction) IsModified() bool {
	return action.isModified
}

func (action BaseAction) SetModified(isModified bool) {
	action.isModified = isModified
}

func (action BaseAction) SetBlocked(isBlocked bool) {
	action.isBlocked = isBlocked
}
