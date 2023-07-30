package player

import "github.com/goseventh/rakstar/internal/samp"

func (pb *PlayerBuilder) Connected(status *bool) *PlayerBuilder {
	*status = samp.IsPlayerConnected(pb.ID)
	return pb
}
