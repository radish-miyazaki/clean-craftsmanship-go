package bowring

type Game struct {
	rolls []int
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) isSpare(fi int) bool {
	return g.rolls[fi]+g.rolls[fi+1] == 10
}

func (g *Game) spareBonus(fi int) int {
	return g.rolls[fi+2]
}

func (g *Game) isStrike(fi int) bool {
	return g.rolls[fi] == 10
}

func (g *Game) strikeBonus(fi int) int {
	return g.rolls[fi+1] + g.rolls[fi+2]
}

func (g *Game) twoBallsInFrame(fi int) int {
	return g.rolls[fi] + g.rolls[fi+1]
}

func (g *Game) Roll(p int) {
	g.rolls = append(g.rolls, p)
}

func (g *Game) Score() int {
	s := 0
	fi := 0
	for frame := 0; frame < 10; frame++ {
		if g.isStrike(fi) {
			s += 10 + g.strikeBonus(fi)
			fi++
		} else if g.isSpare(fi) {
			s += 10 + g.spareBonus(fi)
			fi += 2
		} else {
			s += g.twoBallsInFrame(fi)
			fi += 2
		}
	}

	return s
}
