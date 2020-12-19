package part_1

import "github.com/cmdaltent/advent-of-code-2020/day_12/common"

type Ship struct {
	Position common.Position
	Heading  common.Heading
}

func (s *Ship) Apply(ni common.NavigationInstruction) {
	switch ni.Action {
	case common.N:
		s.Position = common.Position{
			Y: s.Position.Y + ni.Value,
			X: s.Position.X,
		}
	case common.S:
		s.Position = common.Position{
			Y: s.Position.Y - ni.Value,
			X: s.Position.X,
		}
	case common.E:
		s.Position = common.Position{
			Y: s.Position.Y,
			X: s.Position.X + ni.Value,
		}
	case common.W:
		s.Position = common.Position{
			Y: s.Position.Y,
			X: s.Position.X - ni.Value,
		}
	case common.F:
		s.Apply(common.NavigationInstruction{
			Action: s.Heading.ToAction(),
			Value:  ni.Value,
		})
	case common.L:
		s.Heading = s.Heading.TurnLeft(ni.Value)
	case common.R:
		s.Heading = s.Heading.TurnRight(ni.Value)
	}
}

func NewShip(h common.Heading) Ship {
	return Ship{
		Heading: h,
	}
}
