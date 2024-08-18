package domain

// Player is the base domain object.
type Player struct {
	ID   int
	Name string
}

// Footballer is a type of Player with additional attributes.
type Footballer struct {
	Player
	Club string
}

// Cricketer is another type of Player with different attributes.
type Cricketer struct {
	Player
	BattingAverage float64
}

// Bowler is a type of Cricketer with an additional attribute.
type Bowler struct {
	Cricketer
	BowlingAverage float64
}
