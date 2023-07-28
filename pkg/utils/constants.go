package models

const (

	MaxChatbubbleLength              int = 144
	MapiconLocal                     int = 0
	MapiconGlobal                    int = 1
	MapiconLocalCheckpoint           int = 2
	MapiconGlobalCheckpoint          int = 3
	CameraCut                        int = 2
	CameraMove                       int = 1
	SpectateModeNormal               int = 1
	SpectateModeFixed                int = 2
	SpectateModeSide                 int = 3
	CarmodtypeSpoiler                int = 0
	CarmodtypeHood                   int = 1
	CarmodtypeRoof                   int = 2
	CarmodtypeSideskirt              int = 3
	CarmodtypeLamps                  int = 4
	CarmodtypeNitro                  int = 5
	CarmodtypeExhaust                int = 6
	CarmodtypeWheels                 int = 7
	CarmodtypeStereo                 int = 8
	CarmodtypeHydraulics             int = 9
	CarmodtypeFrontBumper            int = 10
	CarmodtypeRearBumper             int = 11
	CarmodtypeVentRight              int = 12
	CarmodtypeVentLeft               int = 13

	


	InvalidActorId                   int = 0xFFFF
	NoTeam                           int = 255
	MaxObjects                       int = 1000
	InvalidObjectId                  int = 0xFFFF
	MaxGangZones                     int = 1024
	MaxTextDraws                     int = 2048
	MaxMenus                         int = 128
	Max3dtextGlobal                  int = 1024
	MaxPickups                       int = 4096
	InvalidMenu                      int = 0xFF
	InvalidTextDraw                  int = 0xFFFF
	InvalidGangZone                  int = -1
	Invalid3dtextId                  int = 0xFFFF
	ServerVartypeNone                int = 0
	ServerVartypeInt                 int = 1
	ServerVartypeString              int = 2
	ServerVartypeFloat               int = 3
	TextDrawFontSpriteDraw           int = 4
	TextDrawFontModelPreview         int = 5
	DialogStyleMsgbox                int = 0
	DialogStyleInput                 int = 1
	DialogStyleList                  int = 2
	DialogStylePassword              int = 3
	DialogStyleTablist               int = 4
	DialogStyleTablistHeaders        int = 5

	#Weapon
	


	

	KeyAction                        int = 1
	KeyCrouch                        int = 2
	KeyFire                          int = 4
	KeySprint                        int = 8
	KeySecondaryAttack               int = 16
	KeyJump                          int = 32
	KeyLookRight                     int = 64
	KeyHandbrake                     int = 128
	KeyLookLeft                      int = 256
	KeySubmission                    int = 512
	KeyLookBehind                    int = 512
	KeyWalk                          int = 1024
	KeyAnalogUp                      int = 2048
	KeyAnalogDown                    int = 4096
	KeyAnalogLeft                    int = 8192
	KeyAnalogRight                   int = 16384
	KeyYes                           int = 65536
	KeyNo                            int = 131072
	KeyCtrlBack                      int = 262144
	KeyUp                            int = -128
	KeyDown                          int = 128
	KeyLeft                          int = -128
	KeyRight                         int = 128
	BodyPartTorso                    int = 3
	BodyPartGroin                    int = 4
	BodyPartLeftArm                  int = 5
	BodyPartRightArm                 int = 6
	BodyPartLeftLeg                  int = 7
	BodyPartRightLeg                 int = 8
	BodyPartHead                     int = 9
	ClickSourceScoreboard            int = 0
	EditResponseCancel               int = 0
	EditResponseFinal                int = 1
	EditResponseUpdate               int = 2
	SelectObjectGlobalObject         int = 1
	SelectObjectPlayerObject         int = 2
	BulletHitTypeNone                int = 0
	BulletHitTypePlayer              int = 1

	BulletHitTypeObject              int = 3
	BulletHitTypePlayerObject        int = 4
	DownloadRequestEmpty             int = 0
	DownloadRequestModelFile         int = 1
	DownloadRequestTextureFile       int = 2
	ObjectMaterialSize32x32          int = 10
	ObjectMaterialSize64x32          int = 20
	ObjectMaterialSize64x64          int = 30
	ObjectMaterialSize128x32         int = 40
	ObjectMaterialSize128x64         int = 50
	ObjectMaterialSize128x128        int = 60
	ObjectMaterialSize256x32         int = 70
	ObjectMaterialSize256x64         int = 80
	ObjectMaterialSize256x128        int = 90
	ObjectMaterialSize256x256        int = 100
	ObjectMaterialSize512x64         int = 110
	ObjectMaterialSize512x128        int = 120
	ObjectMaterialSize512x256        int = 130
	ObjectMaterialSize512x512        int = 140
	ObjectMaterialTextAlignLeft      int = 0
	ObjectMaterialTextAlignCenter    int = 1
	ObjectMaterialTextAlignRight     int = 2
)
