package basic

type Invisible struct {
	Id int
}

func New(id int) Invisible {
	return Invisible{id}
}
