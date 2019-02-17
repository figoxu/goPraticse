package ormsample

import (
	"fmt"
	"github.com/figoxu/Figo"
	"testing"
)

func TestPolymorphic(t *testing.T) {
	cat := Cat{Name: "Mr. Bigglesworth", Toy: Toy{Name: "cat toy"}}
	dog := Dog{Name: "Pluto", Toys: []Toy{{Name: "dog toy 1"}, {Name: "dog toy 2"}}}
	env.db.Save(&cat).Save(&dog)
}

func TestPolymorphicDog(t *testing.T) {
	dog := &Dog{}
	env.db.Where("id=?", 1).Preload("Toys").Find(dog)
	fmt.Println(Figo.JsonString(dog))
}
