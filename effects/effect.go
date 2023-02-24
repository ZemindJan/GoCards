package effects

// Effects dispatch actions, usually when cards are played
type Effect interface {
	Trigger(context EffectContext)
}
