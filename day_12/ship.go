package main

type heading uint8

const (
	north heading = iota
	east
	south
	west
)

func (h heading) ToAction() action {
	switch h {
	case north:
		return N
	case south:
		return S
	case west:
		return W
	case east:
		return E
	}
	return F
}

func (h heading) TurnLeft(degree int) heading {
	if degree == 360 {
		return h
	}
	return heading((int(h) + ((360 - degree) * 4 / 360)) % 4)
}

func (h heading) TurnRight(degree int) heading {
	if degree == 360 {
		return h
	}
	return heading((int(h) + (degree * 4 / 360)) % 4)
}

type position struct {
	y int // north/south
	x int // east/west
}

type ship struct {
	pos  position
	head heading
}

func (s *ship) Apply(ni navigationInstruction) {
	switch ni.act {
	case N:
		s.pos = position{
			y: s.pos.y + ni.value,
			x: s.pos.x,
		}
	case S:
		s.pos = position{
			y: s.pos.y - ni.value,
			x: s.pos.x,
		}
	case E:
		s.pos = position{
			y: s.pos.y,
			x: s.pos.x + ni.value,
		}
	case W:
		s.pos = position{
			y: s.pos.y,
			x: s.pos.x - ni.value,
		}
	case F:
		s.Apply(navigationInstruction{
			act:   s.head.ToAction(),
			value: ni.value,
		})
	case L:
		s.head = s.head.TurnLeft(ni.value)
	case R:
		s.head = s.head.TurnRight(ni.value)
	}
}

func newShip(h heading) ship {
	return ship{
		head: h,
	}
}

type action uint8

const (
	N action = iota // move north
	S               // move south
	E               // move east
	W               // move west
	L               // turn left
	R               // turn right
	F               // head forward
)

type navigationInstruction struct {
	act   action
	value int
}
