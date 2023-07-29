package dialog

import (
	"time"
)

type playerChannelPool chan *DialogResponse

var playerChannelsPool [1000][]playerChannelPool

func poolNext(id int) playerChannelPool {
	if len(playerChannelsPool[id]) == 0 {
		return nil
	}

	poolChanel := playerChannelsPool[id][0]

	if len(playerChannelsPool[id]) > 1 {
		playerChannelsPool[id] = playerChannelsPool[id][1:]
	} else {
		playerChannelsPool[id] = []playerChannelPool{}
	}

	return poolChanel
}

func poolPush(id int) playerChannelPool {
	channel := make(playerChannelPool)
	playerChannelsPool[id] = append(playerChannelsPool[id], channel)

	time.AfterFunc(time.Second*10, func() {
		playerChannelsLen := len(playerChannelsPool[id])

		for i := 0; i < playerChannelsLen; i++ {
			poolChannel := playerChannelsPool[id][i]

			if poolChannel != channel {
				continue
			}

			sliceA := i
			sliceB := i + 1

			if sliceB >= playerChannelsLen {
				sliceB = playerChannelsLen
			}

			poolChannel <- nil

			Builder().
				Select(id).
				Close()
			playerChannelsPool[id] = append(playerChannelsPool[id][0:sliceA], playerChannelsPool[id][sliceB:]...)

			break
		}
	})

	return channel
}
