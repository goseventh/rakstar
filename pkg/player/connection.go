package player

import "github.com/goseventh/rakstar/internal/natives"

func (pb *PlayerBuilder) Connected(status *bool) *PlayerBuilder {
	*status = natives.IsPlayerConnected(pb.ID)
	return pb
}
