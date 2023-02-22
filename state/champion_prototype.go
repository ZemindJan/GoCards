package state

type ChampionPrototype struct {
	Name      string
	MaxHealth int
	MaxCharge int
}

func (pr ChampionPrototype) GetName() string {
	return pr.Name
}
