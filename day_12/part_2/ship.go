package part_2

import "github.com/cmdaltent/advent-of-code-2020/day_12/common"

type Ship struct {
	waypoint common.Position
	Position common.Position
}

func (s *Ship) Apply(navigationInstruction common.NavigationInstruction) {
	switch navigationInstruction.Action {
	case common.N:
		s.waypoint = common.Position{
			Y: s.waypoint.Y + navigationInstruction.Value,
			X: s.waypoint.X,
		}
	case common.S:
		s.waypoint = common.Position{
			Y: s.waypoint.Y - navigationInstruction.Value,
			X: s.waypoint.X,
		}
	case common.E:
		s.waypoint = common.Position{
			Y: s.waypoint.Y,
			X: s.waypoint.X + navigationInstruction.Value,
		}
	case common.W:
		s.waypoint = common.Position{
			Y: s.waypoint.Y,
			X: s.waypoint.X - navigationInstruction.Value,
		}
	case common.F:
		s.Position = common.Position{
			Y: s.Position.Y + (navigationInstruction.Value * s.waypoint.Y),
			X: s.Position.X + (navigationInstruction.Value * s.waypoint.X),
		}
	case common.L:
		s.waypoint = s.waypoint.RotateCounterClockwise(navigationInstruction.Value)
	case common.R:
		s.waypoint = s.waypoint.RotateClockwise(navigationInstruction.Value)
	}
}

func NewShip(waypoint common.Position) Ship {
	return Ship{
		waypoint: waypoint,
	}
}
