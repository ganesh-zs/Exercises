package temperature

type CelToFer struct {
	Temp float64
}

type FerToCel struct {
	Temp float64
}

func (temp CelToFer) Convertion() float64 {
	return temp.Temp*9.0/5.0 + 32.0
}

func (temp FerToCel) Convertion() float64 {
	return (temp.Temp - 32.0) * (5.0 / 9.0)
}
