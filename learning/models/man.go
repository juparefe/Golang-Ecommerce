package models

type Man struct {
	age       int
	height    float32
	weight    float32
	breathing bool
	thinking  bool
	eating    bool
	alive     bool
}

func (m *Man) Breathe()       { m.breathing = true }
func (m *Man) Think()         { m.thinking = true }
func (m *Man) Eat()           { m.eating = true }
func (m *Man) Gender() string { return "Man" }
