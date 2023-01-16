package store

type travel interface {
	move() int
}
type car struct {
	name string
}

func (c car) move() int {
	return 2
}

func main() {
	var c travel
	c = car{name: "camry"}
	c.move()
}
