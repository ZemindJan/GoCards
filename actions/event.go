package actions

// A type of action that cannot be altered or blocked, and has no execution
type Event struct {
}

func (ev *Event) IsBlocked() bool             { return false }
func (ev *Event) SetBlocked(isBlocked bool)   {}
func (ev *Event) IsModified() bool            { return false }
func (ev *Event) SetModified(isModified bool) {}
func (ev *Event) Execute()                    {}
