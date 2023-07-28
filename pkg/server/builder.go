package server

type ServerBuild struct {
	msgRestart string
	msgLoop    string
	tag        string
	playerID   int
	message    string
}

func Builder() *ServerBuild {
	b := new(ServerBuild)
	b.playerID = -1
	return b
}
