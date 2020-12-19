package part_1

import "github.com/cmdaltent/advent-of-code-2020/day_12/common"

type ship struct {
	position common.Position
	heading  common.Heading
}

func (s *ship) Apply(ni common.NavigationInstruction) {
	switch ni.Action {
	case common.N:
		s.position = common.Position{
			Y: s.position.Y + ni.Value,
			X: s.position.X,
		}
	case common.S:
		s.position = common.Position{
			Y: s.position.Y - ni.Value,
			X: s.position.X,
		}
	case common.E:
		s.position = common.Position{
			Y: s.position.Y,
			X: s.position.X + ni.Value,
		}
	case common.W:
		s.position = common.Position{
			Y: s.position.Y,
			X: s.position.X - ni.Value,
		}
	case common.F:
		s.Apply(common.NavigationInstruction{
			Action: s.heading.ToAction(),
			Value:  ni.Value,
		})
	case common.L:
		s.heading = s.heading.TurnLeft(ni.Value)
	case common.R:
		s.heading = s.heading.TurnRight(ni.Value)
	}
}

func newShip(h common.Heading) ship {
	return ship{
		heading: h,
	}
}
