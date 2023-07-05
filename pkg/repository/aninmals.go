package repository

type Aninmal struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type AninmalRepository struct{}

func (r AninmalRepository) GetAninmals() ([]Aninmal, error) {
	aninmals := []Aninmal{
		{ID: "1", Name: "Aquatic Sock Puppet"},
		{ID: "2", Name: "Patience Monkey"},
	}
	return aninmals, nil
}

func (r AninmalRepository) GetAninmalById(int) (Aninmal, error) {
	aninmal := Aninmal{ID: "1", Name: "Aquatic Sock Puppet"}
	return aninmal, nil
}
