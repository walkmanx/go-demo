package example14.model

type Product struct {
	name  string
	price float32
}

func (this *Product) setName(name string) {
	this.name = name
}

func (this *Product) getName() string {
	return this.name
}
