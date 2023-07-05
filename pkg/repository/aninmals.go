package repository

type Aninmal struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type AninmalRepository struct{}

func (r AninmalRepository) GetAninmals() []Aninmal {
	aninmals := []Aninmal{
		{ID: "1", Name: "Aquatic Sock Puppet"},
		{ID: "2", Name: "Patience Monkey"},
	}

	return aninmals
}
