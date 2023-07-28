package models

const (
	
	FightStyleNormal                 int = 4
	FightStyleBoxing                 int = 5
	FightStyleKungfu                 int = 6
	FightStyleKneehead               int = 7
	FightStyleGrabkick               int = 15
	FightStyleElbow                  int = 16
	

	
	
	


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
	VehicleParamsUnset               int = -1
	VehicleParamsOff                 int = 0
	VehicleParamsOn                  int = 1
	VehicleModelInfoSize             int = 1
	VehicleModelInfoFrontseat        int = 2
	VehicleModelInfoRearseat         int = 3
	VehicleModelInfoPetrolcap        int = 4
	VehicleModelInfoWheelsfront      int = 5
	VehicleModelInfoWheelsrear       int = 6
	VehicleModelInfoWheelsmid        int = 7
	VehicleModelInfoFrontBumperZ     int = 8
	VehicleModelInfoRearBumperZ      int = 9
	MaxClientMessage                 int = 144
	MaxVehicles                      int = 2000
	MaxActors                        int = 1000

	InvalidVehicleId                 int = 0xFFFF
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
	


	VehicleLandstalker               int = 400
	VehicleBravura                   int = 401
	VehicleBuffalo                   int = 402
	VehicleLinerunner                int = 403
	VehiclePerrenial                 int = 404
	VehicleSentinel                  int = 405
	VehicleDumper                    int = 406
	VehicleFiretruck                 int = 407
	VehicleTrashmaster               int = 408
	VehicleStretch                   int = 409
	VehicleManana                    int = 410
	VehicleInfernus                  int = 411
	VehicleVoodoo                    int = 412
	VehiclePony                      int = 413
	VehicleMule                      int = 414
	VehicleCheetah                   int = 415
	VehicleAmbulance                 int = 416
	VehicleLeviathan                 int = 417
	VehicleMoonbeam                  int = 418
	VehicleEsperanto                 int = 419
	VehicleTaxi                      int = 420
	VehicleWashington                int = 421
	VehicleBobcat                    int = 422
	VehicleMrwhoopee                 int = 423
	VehicleBfinjection               int = 424
	VehicleHunter                    int = 425
	VehiclePremier                   int = 426
	VehicleEnforcer                  int = 427
	VehicleSecuricar                 int = 428
	VehicleBanshee                   int = 429
	VehiclePredator                  int = 430
	VehicleBus                       int = 431
	VehicleRhino                     int = 432
	VehicleBarracks                  int = 433
	VehicleHotknife                  int = 434
	VehicleArticletrailer1           int = 435
	VehiclePrevion                   int = 436
	VehicleCoach                     int = 437
	VehicleCabbie                    int = 438
	VehicleStallion                  int = 439
	VehicleRumpo                     int = 440
	VehicleRcbandit                  int = 441
	VehicleRomero                    int = 442
	VehiclePacker                    int = 443
	VehicleMonster                   int = 444
	VehicleAdmiral                   int = 445
	VehicleSqualo                    int = 446
	VehicleSeasparrow                int = 447
	VehiclePizzaboy                  int = 448
	VehicleTram                      int = 449
	VehicleArticletrailer2           int = 450
	VehicleTurismo                   int = 451
	VehicleSpeeder                   int = 452
	VehicleReefer                    int = 453
	VehicleTropic                    int = 454
	VehicleFlatbed                   int = 455
	VehicleYankee                    int = 456
	VehicleCaddy                     int = 457
	VehicleSolair                    int = 458
	VehicleBerkleysrcvan             int = 459
	VehicleSkimmer                   int = 460
	VehiclePcj600                    int = 461
	VehicleFaggio                    int = 462
	VehicleFreeway                   int = 463
	VehicleRcbaron                   int = 464
	VehicleRcraider                  int = 465
	VehicleGlendale                  int = 466
	VehicleOceanic                   int = 467
	VehicleSanchez                   int = 468
	VehicleSparrow                   int = 469
	VehiclePatriot                   int = 470
	VehicleQuad                      int = 471
	VehicleCoastguard                int = 472
	VehicleDinghy                    int = 473
	VehicleHermes                    int = 474
	VehicleSabre                     int = 475
	VehicleRustler                   int = 476
	VehicleZr350                     int = 477
	VehicleWalton                    int = 478
	VehicleRegina                    int = 479
	VehicleComet                     int = 480
	VehicleBmx                       int = 481
	VehicleBurrito                   int = 482
	VehicleCamper                    int = 483
	VehicleMarquis                   int = 484
	VehicleBaggage                   int = 485
	VehicleDozer                     int = 486
	VehicleMaverick                  int = 487
	VehicleSannewsmaverick           int = 488
	VehicleRancher                   int = 489
	VehicleFbirancher                int = 490
	VehicleVirgo                     int = 491
	VehicleGreenwood                 int = 492
	VehicleJetmax                    int = 493
	VehicleHotringracer              int = 494
	VehicleSandking                  int = 495
	VehicleBlistacompact             int = 496
	VehiclePolicemaverick            int = 497
	VehicleBoxville                  int = 498
	VehicleBenson                    int = 499
	VehicleMesa                      int = 500
	VehicleRcgoblin                  int = 501
	VehicleHotringracera             int = 502
	VehicleHotringracerb             int = 503
	VehicleBloodringbanger           int = 504
	VehicleRancherlure               int = 505
	VehicleSupergt                   int = 506
	VehicleElegant                   int = 507
	VehicleJourney                   int = 508
	VehicleBike                      int = 509
	VehicleMountainbike              int = 510
	VehicleBeagle                    int = 511
	VehicleCropdust                  int = 512
	VehicleStuntplane                int = 513
	VehicleTanker                    int = 514
	VehicleRoadtrain                 int = 515
	VehicleNebula                    int = 516
	VehicleMajestic                  int = 517
	VehicleBuccaneer                 int = 518
	VehicleShamal                    int = 519
	VehicleHydra                     int = 520
	VehicleFcr900                    int = 521
	VehicleNrg500                    int = 522
	VehicleHpv1000                   int = 523
	VehicleCementtruck               int = 524
	VehicleTowtruck                  int = 525
	VehicleFortune                   int = 526
	VehicleCadrona                   int = 527
	VehicleFbitruck                  int = 528
	VehicleWillard                   int = 529
	VehicleForklift                  int = 530
	VehicleTractor                   int = 531
	VehicleCombine                   int = 532
	VehicleFeltzer                   int = 533
	VehicleRemington                 int = 534
	VehicleSlamvan                   int = 535
	VehicleBlade                     int = 536
	VehicleFreight                   int = 537
	VehicleBrownstreak               int = 538
	VehicleVortex                    int = 539
	VehicleVincent                   int = 540
	VehicleBullet                    int = 541
	VehicleClover                    int = 542
	VehicleSadler                    int = 543
	VehicleFiretruckla               int = 544
	VehicleHustler                   int = 545
	VehicleIntruder                  int = 546
	VehiclePrimo                     int = 547
	VehicleCargobob                  int = 548
	VehicleTampa                     int = 549
	VehicleSunrise                   int = 550
	VehicleMerit                     int = 551
	VehicleUtilityvan                int = 552
	VehicleNevada                    int = 553
	VehicleYosemite                  int = 554
	VehicleWindsor                   int = 555
	VehicleMonstera                  int = 556
	VehicleMonsterb                  int = 557
	VehicleUranus                    int = 558
	VehicleJester                    int = 559
	VehicleSultan                    int = 560
	VehicleStratum                   int = 561
	VehicleElegy                     int = 562
	VehicleRaindance                 int = 563
	VehicleRctiger                   int = 564
	VehicleFlash                     int = 565
	VehicleTahoma                    int = 566
	VehicleSavanna                   int = 567
	VehicleBandito                   int = 568
	VehicleFreightflattrailer        int = 569
	VehicleStreaktrailer             int = 570
	VehicleKart                      int = 571
	VehicleMower                     int = 572
	VehicleDuneride                  int = 573
	VehicleSweeper                   int = 574
	VehicleBroadway                  int = 575
	VehicleTornado                   int = 576
	VehicleAt400                     int = 577
	VehicleDft30                     int = 578
	VehicleHuntley                   int = 579
	VehicleStafford                  int = 580
	VehicleBf400                     int = 581
	VehicleNewsvan                   int = 582
	VehicleTug                       int = 583
	VehiclePetroltrailer             int = 584
	VehicleEmperor                   int = 585
	VehicleWayfarer                  int = 586
	VehicleEuros                     int = 587
	VehicleHotdog                    int = 588
	VehicleClub                      int = 589
	VehicleFreightboxtrailer         int = 590
	VehicleArticletrailer3           int = 591
	VehicleAndromada                 int = 592
	VehicleDodo                      int = 593
	VehicleRccam                     int = 594
	VehicleLaunch                    int = 595
	VehiclePolicecarlspd             int = 596
	VehiclePolicecarsfpd             int = 597
	VehiclePolicecarlvpd             int = 598
	VehiclePoliceranger              int = 599
	VehiclePicador                   int = 600
	VehicleSwat                      int = 601
	VehicleAlpha                     int = 602
	VehiclePhoenix                   int = 603
	VehicleGlendaleshit              int = 604
	VehicleSadlershit                int = 605
	VehicleBaggagetrailera           int = 606
	VehicleBaggagetrailerb           int = 607
	VehicleTugstairstrailer          int = 608
	VehicleBoxburg                   int = 609
	VehicleFarmtrailer               int = 610
	VehicleUtilitytrailer            int = 611
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
	BulletHitTypeVehicle             int = 2
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
