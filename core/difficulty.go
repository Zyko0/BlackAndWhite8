package core

type Difficulty byte

const (
	DifficultyNormal Difficulty = iota
	DifficultyHard
)

func (d Difficulty) String() string {
	switch d {
	case DifficultyNormal:
		return "Normal"
	case DifficultyHard:
		return "Hard"
	default:
		return ""
	}
}
