package model

import "github.com/neerajsurjaye/spen/internal/enums"

type Collider interface {
	Type() enums.ColliderType
}