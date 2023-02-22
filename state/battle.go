package state

type Battle struct {
	Player    *Champion
	Adversary *Champion
	turn      int
}

// Returns the oppponent of a champion
func (bd *Battle) Opponent(champion *Champion) *Champion {
	switch champion {
	case bd.Player:
		return bd.Adversary
	case bd.Adversary:
		return bd.Adversary
	default:
		return nil
	}
}
