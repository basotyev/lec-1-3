package cat

type Cat struct {
	sound string
}

func NewCat(sound string) Cat {
	return Cat{
		sound: sound,
	}
}

func (c *Cat) SetSound(sound string) {
	c.sound = sound
}

func (c *Cat) MakeASound() string {
	return c.sound
}

func (c *Cat) GetName() string {
	return "Duck"
}

func (c *Cat) IsDuck() bool {
	return false
}

func (c *Cat) IsCat() bool {
	return true
}
