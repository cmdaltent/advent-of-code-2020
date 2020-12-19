package common

type Heading uint8

const (
	North Heading = iota
	East
	South
	West
)

func (h Heading) ToAction() Action {
	switch h {
	case North:
		return N
	case South:
		return S
	case West:
		return W
	case East:
		return E
	}
	return F
}

func (h Heading) TurnLeft(degree int) Heading {
	if degree == 360 {
		return h
	}
	return Heading((int(h) + ((360 - degree) * 4 / 360)) % 4)
}

func (h Heading) TurnRight(degree int) Heading {
	if degree == 360 {
		return h
	}
	return Heading((int(h) + (degree * 4 / 360)) % 4)
}

type Position struct {
	Y int // North/South
	X int // East/West
}

const (
	N Action = iota // move North
	S               // move South
	E               // move East
	W               // move West
	L               // turn left
	R               // turn right
	F               // head forward
)

type Action uint8

type NavigationInstruction struct {
	Action Action
	Value  int
}
