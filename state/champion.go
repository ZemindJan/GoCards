package state

type Champion struct {
	Prototype *ChampionPrototype
	Health    BoundedStat[int]
	Charge    BoundedStat[int]
	Offering  Offering
}

func (ch *Champion) GetName() string {
	return ch.Prototype.GetName()
}

func NewChampion(prototype *ChampionPrototype) *Champion {
	return &Champion{
		Prototype: prototype,
		Health:    MakeBoundedStat(80, 0, 80),
		Charge:    MakeBoundedStat(7, 0, 7),
	}
}
