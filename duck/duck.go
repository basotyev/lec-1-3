package duck

import "fmt"

type Duck struct {
	Sound string
}

func (d *Duck) MakeASound() string {
	return d.Sound
}

func (d *Duck) GetName() string {
	return "Duck"
}

func (d *Duck) IsDuck() bool {
	return true
}

func (d *Duck) secretFunction() {
	fmt.Println("QQQQQQQQQQRRRRRRRRRRRRRRYAAAAAAAAAAAAAAAAAAAAA")
}

func (d *Duck) SecretFunction() {
	fmt.Println("QQQQQQQQQQRRRRRRRRRRRRRRYAAAAAAAAAAAAAAAAAAAAA")
}
