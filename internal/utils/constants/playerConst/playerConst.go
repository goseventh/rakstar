package playerConst

const (
	MaxPlayerAttachedObjects         int = 10
	MaxPlayerName                    int = 24
	MaxPlayers                       int = 1000
	MaxPlayerTextDraws               int = 256
	Max3dtextPlayer                  int = 1024
	PlayerVartypeNone                int = 0
	PlayerVartypeInt                 int = 1
	PlayerVartypeString              int = 2
	PlayerVartypeFloat               int = 3
	PlayerRecordingTypeNone          int = 0
	PlayerRecordingTypeDriver        int = 1
	PlayerRecordingTypeOnfoot        int = 2
	PlayerStateNone                  int = 0
	PlayerStateOnfoot                int = 1
	PlayerStateDriver                int = 2
	PlayerStatePassenger             int = 3
	PlayerStateExitVehicle           int = 4
	PlayerStateEnterVehicleDriver    int = 5
	PlayerStateEnterVehiclePassenger int = 6
	PlayerStateWasted                int = 7
	PlayerStateSpawned               int = 8
	PlayerStateSpectating            int = 9
	PlayerMarkersModeOff             int = 0
	PlayerMarkersModeGlobal          int = 1
	PlayerMarkersModeStreamed        int = 2
	InvalidPlayerId                  int = 0xFFFF
)
