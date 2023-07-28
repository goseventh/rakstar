package models

const (
	MaxChatbubbleLength     int = 144

	InvalidActorId int = 0xFFFF
	NoTeam         int = 255

	MaxGangZones              int = 1024
	
	MaxMenus                  int = 128
	Max3dtextGlobal           int = 1024
	MaxPickups                int = 4096
	InvalidMenu               int = 0xFF
	InvalidGangZone           int = -1
	Invalid3dtextId           int = 0xFFFF
	ServerVartypeNone         int = 0
	ServerVartypeInt          int = 1
	ServerVartypeString       int = 2
	ServerVartypeFloat        int = 3


	//Weapon


	ClickSourceScoreboard int = 0
	EditResponseCancel    int = 0
	EditResponseFinal     int = 1
	EditResponseUpdate    int = 2

	BulletHitTypeNone   int = 0
	BulletHitTypePlayer int = 1

	DownloadRequestEmpty       int = 0
	DownloadRequestModelFile   int = 1
	DownloadRequestTextureFile int = 2
)
