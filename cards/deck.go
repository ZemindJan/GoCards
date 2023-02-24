package cards

type Deck struct {
	name  string
	Cards []Card
}

func (dk *Deck) GetName() string {
	return dk.name
}

func (dk *Deck) GetCards() []Card {
	return dk.Cards
}
