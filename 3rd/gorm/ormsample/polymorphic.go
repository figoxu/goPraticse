package ormsample

type Cat struct {
	Id   int
	Name string
	Toy  Toy `gorm:"polymorphic:Owner;"`
}

type Dog struct {
	Id   int
	Name string
	Toys []Toy `gorm:"polymorphic:Owner;"`
}

type Toy struct {
	Id        int
	Name      string
	OwnerId   int
	OwnerType string
}