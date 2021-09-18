package person

type Person struct {
	Name string
	Age  int
}

func (p *Person) GetName() string {
	return p.Name
}
