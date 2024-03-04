package models

type Women struct {
	Man
}

func (w *Women) Breathe()       { w.breathing = true }
func (w *Women) Think()         { w.thinking = true }
func (w *Women) Eat()           { w.eating = true }
func (w *Women) Gender() string { return "Women" }
