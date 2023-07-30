package callbacks

/*
#cgo windows CFLAGS: -I../lib -I../lib/amx -Wno-attributes -Wno-implicit-function-declaration
#cgo windows CFLAGS: -DHAVE_INTTYPES_H -DHAVE_MALLOC_H -DHAVE_STDINT_H -DWIN32
#cgo windows LDFLAGS: -Wl,--subsystem,windows,--kill-at

#cgo linux CFLAGS: -I../lib -I../lib/amx -Wno-attributes -Wno-implicit-function-declaration
#cgo linux CFLAGS: -DHAVE_INTTYPES_H -DHAVE_MALLOC_H -DHAVE_STDINT_H -DLINUX -D_GNU_SOURCE
#cgo linux LDFLAGS: -ldl

#ifndef GOLANG_APP
#define GOLANG_APP

#include "main.h"

#endif
*/
import "C"
import (
	"fmt"
	"unsafe"

	"github.com/goseventh/rakstar/internal/natives"
	"github.com/goseventh/rakstar/internal/serverlog"
)

type EventType int

const (
	Repeat EventType = iota
	OnceOnly
)

type event struct {
	Handler interface{}
	Type    EventType
}

var Events = make(map[string]event)
var mainEvent func() = nil

func On(eventName string, handler interface{}) error {
	_, ok := Events[eventName]
	if ok {
		return fmt.Errorf("this handler already exists")
	}

	Events[eventName] = event{Handler: handler, Type: Repeat}
	_ = serverlog.Print(fmt.Sprintf("Registered %s event", eventName))

	return nil
}

func Once(eventName string, handler interface{}) error {
	_, ok := Events[eventName]
	if ok {
		return fmt.Errorf("this handler already exists")
	}

	Events[eventName] = event{Handler: handler, Type: OnceOnly}
	_ = serverlog.Print(fmt.Sprintf("Registered %s event", eventName))

	return nil
}

//export onTick
func onTick() {
	evt, ok := Events["tick"]
	if !ok {
		return
	}

	fn, ok := evt.Handler.(func())
	if !ok {
		return
	}
	fn()

	if evt.Type == OnceOnly {
		// If this event was registered only once, then reassign this event to a new blank struct.
		// PoC code. Still needs to be implemented completely.
	}

	return
}

//export callEvent
func callEvent(amx *C.AMX, funcName *C.char_t, format *C.char_t, params []C.cell) bool {
	name := C.GoString(C.constToNonConst(funcName))
	specifiers := C.GoString(C.constToNonConst(format))

	evt, ok := Events[name]
	if !ok {
		_ = serverlog.Print(fmt.Sprintf("rakstar: Called an event ('%s') that is not registered by ", name))
		return false
	}

	_ = serverlog.Print("callEvent (1)")
	specifiersLen := len(specifiers)

	if specifiersLen == 0 {
		fn, ok := evt.Handler.(func())
		if !ok {
			_ = serverlog.Print(fmt.Sprintf("rakstar: Event ('%s') failed to call", name))
			return false
		}
		_ = serverlog.Print("callEvent (2)")
		fn()
	} else {
		fn, ok := evt.Handler.(func([]interface{}))
		if !ok {
			_ = serverlog.Print(fmt.Sprintf("rakstar: Event ('%s') failed to call", name))
			return false
		}
		in := make([]interface{}, specifiersLen)
		param_offset := 0
		for i := 0; i < specifiersLen; i++ {
			var index int = i + param_offset + 3
			switch specifiers[i] {
			case 'i', 'd':
				_ = serverlog.Print("It is an int")
				in[i] = int(params[index])
			case 'f':
				_ = serverlog.Print("It is a float")
				in[i] = float32(params[index])
			case 's':
				_ = serverlog.Print("It is a string")
				var maddr *C.cell
				var len C.int = 0
				if C.amx_GetAddr(amx, params[index], &maddr) == C.AMX_ERR_NONE || &maddr != nil {
					C.amx_StrLen(maddr, &len)
					len++
					sval := C.malloc(C.uint(C.sizeof_char * (len)))
					defer C.free(unsafe.Pointer(sval))
					param_offset += int(len)
					if C.amx_GetString((*C.char)(sval), maddr, C.int(0), C.uint(len)) == C.AMX_ERR_NONE {
						in[i] = C.GoString((*C.char)(sval))
					}
				}
			}
		}
		_ = serverlog.Print("callEvent (2)")
		fn(in)
	}
	return true
}

//export onGameModeInit
func onGameModeInit() bool {
	evt, ok := Events["goModeInit"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func() bool)
	if !ok {
		return false
	}
	fn()
	return true
}

//export onGameModeExit
func onGameModeExit() bool {
	evt, ok := Events["goModeExit"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func() bool)
	if !ok {
		return false
	}
	fn()
	return true
}

//export onPlayerConnect
func onPlayerConnect(playerid C.int) bool {
	evt, ok := Events["playerConnect"]
	if !ok {
		return false
	}

	natives.SendClientMessage(int(playerid), -1,
		"{ffffff}[{ff5ef2}RAKSTAR{ffffff}]{db34eb} desenvolvimento de servidores samp modernos com o estilo GoSeventh",
	)
	natives.SendClientMessage(int(playerid), -1,
		"{ffffff}[{ff5ef2}RAKSTAR{ffffff}]{ffac12} github.com/goseventh/rakstar",
	)
	fn, ok := evt.Handler.(func(natives.Player) bool)
	if !ok {
		return false
	}

	return fn(natives.Player{ID: int(playerid)})
}

//export onPlayerDisconnect
func onPlayerDisconnect(playerid C.int, reason C.int) bool {
	evt, ok := Events["playerDisconnect"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(natives.Player, int) bool)
	if !ok {
		return false
	}
	return fn(natives.Player{ID: int(playerid)}, int(reason))
}

//export onPlayerSpawn
func onPlayerSpawn(playerid C.int) bool {
	evt, ok := Events["playerSpawn"]
	if !ok {
		return true
	}

	fn, ok := evt.Handler.(func(natives.Player) bool)
	if !ok {
		return true
	}
	return fn(natives.Player{ID: int(playerid)})
}

//export onPlayerDeath
func onPlayerDeath(playerid C.int, killerid C.int, reason C.int) bool {
	evt, ok := Events["playerDeath"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(natives.Player, natives.Player, int) bool)
	if !ok {
		return false
	}
	return fn(natives.Player{ID: int(playerid)}, natives.Player{ID: int(killerid)}, int(reason))
}

//export onVehicleSpawn
func onVehicleSpawn(vehicleid C.int) bool {
	evt, ok := Events["vehicleSpawn"]
	if !ok {
		return true
	}

	fn, ok := evt.Handler.(func(int) bool)
	if !ok {
		return true
	}
	return fn(int(vehicleid))
}

//export onVehicleDeath
func onVehicleDeath(vehicleid C.int, killerid C.int) bool {
	evt, ok := Events["vehicleDeath"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(int, natives.Player) bool)
	if !ok {
		return false
	}
	fn(int(vehicleid), natives.Player{ID: int(killerid)})
	return true
}

//export onPlayerText
func onPlayerText(playerid C.int, text *C.char_t) bool {
	evt, ok := Events["playerText"]
	if !ok {
		return true
	}

	fn, ok := evt.Handler.(func(natives.Player, string) bool)
	if !ok {
		return true
	}
	return fn(natives.Player{ID: int(playerid)}, C.GoString(C.constToNonConst(text)))
}

//export onPlayerCommandText
func onPlayerCommandText(playerid C.int, cmdtext *C.char_t) bool {
	evt, ok := Events["playerCommandText"]
	if !ok {
		return true
	}

	fn, ok := evt.Handler.(func(natives.Player, string) bool)
	if !ok {
		return true
	}
	return fn(natives.Player{ID: int(playerid)}, C.GoString(C.constToNonConst(cmdtext)))
}

//export onPlayerRequestClass
func onPlayerRequestClass(playerid C.int, classid C.int) bool {
	evt, ok := Events["playerRequestClass"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(natives.Player, int) bool)
	if !ok {
		return false
	}
	fn(natives.Player{ID: int(playerid)}, int(classid))
	return true
}

//export onPlayerEnterVehicle
func onPlayerEnterVehicle(playerid C.int, vehicleid C.int, ispassenger C.bool) bool {
	evt, ok := Events["playerEnterVehicle"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(natives.Player, int, bool) bool)
	if !ok {
		return false
	}
	fn(natives.Player{ID: int(playerid)}, int(vehicleid), bool(ispassenger))
	return true
}

//export onPlayerExitVehicle
func onPlayerExitVehicle(playerid C.int, vehicleid C.int) bool {
	evt, ok := Events["playerExitVehicle"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(natives.Player, int) bool)
	if !ok {
		return false
	}
	fn(natives.Player{ID: int(playerid)}, int(vehicleid))
	return true
}

//export onPlayerStateChange
func onPlayerStateChange(playerid C.int, newstate C.int, oldstate C.int) bool {
	evt, ok := Events["playerStateChange"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(natives.Player, int, int) bool)
	if !ok {
		return false
	}
	fn(natives.Player{ID: int(playerid)}, int(newstate), int(oldstate))
	return true
}

//export onPlayerEnterCheckpoint
func onPlayerEnterCheckpoint(playerid C.int) bool {
	evt, ok := Events["playerEnterCheckpoint"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(natives.Player) bool)
	if !ok {
		return false
	}
	fn(natives.Player{ID: int(playerid)})
	return true
}

//export onPlayerLeaveCheckpoint
func onPlayerLeaveCheckpoint(playerid C.int) bool {
	evt, ok := Events["playerLeaveCheckpoint"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(natives.Player) bool)
	if !ok {
		return false
	}
	fn(natives.Player{ID: int(playerid)})
	return true
}

//export onPlayerEnterRaceCheckpoint
func onPlayerEnterRaceCheckpoint(playerid C.int) bool {
	evt, ok := Events["playerEnterRaceCheckpoint"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(natives.Player) bool)
	if !ok {
		return false
	}
	fn(natives.Player{ID: int(playerid)})
	return true
}

//export onPlayerLeaveRaceCheckpoint
func onPlayerLeaveRaceCheckpoint(playerid C.int) bool {
	evt, ok := Events["playerLeaveRaceCheckpoint"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(natives.Player) bool)
	if !ok {
		return false
	}
	fn(natives.Player{ID: int(playerid)})
	return true
}

//export onRconCommand
func onRconCommand(cmd *C.char_t) bool {
	evt, ok := Events["rconCommand"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(string) bool)
	if !ok {
		return false
	}
	return fn(C.GoString(C.constToNonConst(cmd)))
}

//export onPlayerRequestSpawn
func onPlayerRequestSpawn(playerid C.int) bool {
	evt, ok := Events["playerRequestSpawn"]
	if !ok {
		return true
	}

	fn, ok := evt.Handler.(func(natives.Player) bool)
	if !ok {
		return true
	}
	return fn(natives.Player{ID: int(playerid)})
}

//export onObjectMoved
func onObjectMoved(objectid C.int) bool {
	evt, ok := Events["objectMoved"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(int) bool)
	if !ok {
		return false
	}
	fn(int(objectid))
	return true
}

//export onPlayerObjectMoved
func onPlayerObjectMoved(playerid C.int, objectid C.int) bool {
	evt, ok := Events["playerObjectMoved"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(natives.Player, int) bool)
	if !ok {
		return false
	}
	fn(natives.Player{ID: int(playerid)}, int(objectid))
	return true
}

//export onPlayerPickUpPickup
func onPlayerPickUpPickup(playerid C.int, pickupid C.int) bool {
	evt, ok := Events["playerPickUpPickup"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(natives.Player, int) bool)
	if !ok {
		return false
	}
	fn(natives.Player{ID: int(playerid)}, int(pickupid))
	return true
}

//export onVehicleMod
func onVehicleMod(playerid C.int, vehicleid C.int, componentid C.int) bool {
	evt, ok := Events["vehicleMod"]
	if !ok {
		return true
	}

	fn, ok := evt.Handler.(func(natives.Player, int, int) bool)
	if !ok {
		return true
	}
	return fn(natives.Player{ID: int(playerid)}, int(vehicleid), int(componentid))
}

//export onEnterExitModShop
func onEnterExitModShop(playerid C.int, enterexit C.bool, interiorid C.int) bool {
	evt, ok := Events["enterExitModShop"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(natives.Player, bool, int) bool)
	if !ok {
		return false
	}
	fn(natives.Player{ID: int(playerid)}, bool(enterexit), int(interiorid))
	return true
}

//export onVehiclePaintjob
func onVehiclePaintjob(playerid C.int, vehicleid C.int, paintjobid C.int) bool {
	evt, ok := Events["vehiclePaintjob"]
	if !ok {
		return true
	}

	fn, ok := evt.Handler.(func(natives.Player, int, int) bool)
	if !ok {
		return true
	}
	return fn(natives.Player{ID: int(playerid)}, int(vehicleid), int(paintjobid))
}

//export onVehicleRespray
func onVehicleRespray(playerid C.int, vehicleid C.int, color1 C.int, color2 C.int) bool {
	evt, ok := Events["vehicleRespray"]
	if !ok {
		return true
	}

	fn, ok := evt.Handler.(func(natives.Player, int, int, int) bool)
	if !ok {
		return true
	}
	return fn(natives.Player{ID: int(playerid)}, int(vehicleid), int(color1), int(color2))
}

//export onVehicleDamageStatusUpdate
func onVehicleDamageStatusUpdate(vehicleid C.int, playerid C.int) bool {
	evt, ok := Events["vehicleDamageStatusUpdate"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(int, natives.Player) bool)
	if !ok {
		return false
	}
	return fn(int(vehicleid), natives.Player{ID: int(playerid)})
}

//export onUnoccupiedVehicleUpdate
func onUnoccupiedVehicleUpdate(vehicleid C.int, playerid C.int, passenger_seat C.int, new_x C.float, new_y C.float, new_z C.float, vel_x C.float, vel_y C.float, vel_z C.float) bool {
	evt, ok := Events["unoccupiedVehicleUpdate"]
	if !ok {
		return true
	}

	fn, ok := evt.Handler.(func(int, natives.Player, int, float32, float32, float32, float32, float32, float32) bool)
	if !ok {
		return true
	}
	return fn(int(vehicleid), natives.Player{ID: int(playerid)}, int(passenger_seat), float32(new_x), float32(new_y), float32(new_z), float32(vel_x), float32(vel_y), float32(vel_z))
}

//export onPlayerSelectedMenuRow
func onPlayerSelectedMenuRow(playerid C.int, row C.int) bool {
	evt, ok := Events["playerSelectedMenuRow"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(natives.Player, int) bool)
	if !ok {
		return false
	}
	fn(natives.Player{ID: int(playerid)}, int(row))
	return true
}

//export onPlayerExitedMenu
func onPlayerExitedMenu(playerid C.int) bool {
	evt, ok := Events["playerExitedMenu"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(natives.Player) bool)
	if !ok {
		return false
	}
	fn(natives.Player{ID: int(playerid)})
	return true
}

//export onPlayerInteriorChange
func onPlayerInteriorChange(playerid C.int, newinteriorid C.int, oldinteriorid C.int) bool {
	evt, ok := Events["playerInteriorChange"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(natives.Player, int, int) bool)
	if !ok {
		return false
	}
	fn(natives.Player{ID: int(playerid)}, int(newinteriorid), int(oldinteriorid))
	return true
}

//export onPlayerKeyStateChange
func onPlayerKeyStateChange(playerid C.int, newkeys C.int, oldkeys C.int) bool {
	evt, ok := Events["playerKeyStateChange"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(natives.Player, int, int) bool)
	if !ok {
		return false
	}
	fn(natives.Player{ID: int(playerid)}, int(newkeys), int(oldkeys))
	return true
}

//export onRconLoginAttempt
func onRconLoginAttempt(ip *C.char_t, password *C.char_t, success C.bool) bool {
	evt, ok := Events["rconLoginAttempt"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(string, string, bool) bool)
	if !ok {
		return false
	}
	fn(C.GoString(C.constToNonConst(ip)), C.GoString(C.constToNonConst(password)), bool(success))
	return true
}

//export onPlayerUpdate
func onPlayerUpdate(playerid C.int) bool {
	evt, ok := Events["playerUpdate"]
	if !ok {
		return true
	}

	fn, ok := evt.Handler.(func(natives.Player) bool)
	if !ok {
		return true
	}
	return fn(natives.Player{ID: int(playerid)})
}

//export onPlayerStreamIn
func onPlayerStreamIn(playerid C.int, forplayerid C.int) bool {
	evt, ok := Events["playerStreamIn"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(natives.Player, int) bool)
	if !ok {
		return false
	}
	fn(natives.Player{ID: int(playerid)}, int(forplayerid))
	return true
}

//export onPlayerStreamOut
func onPlayerStreamOut(playerid C.int, forplayerid C.int) bool {
	evt, ok := Events["playerStreamOut"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(natives.Player, int) bool)
	if !ok {
		return false
	}
	fn(natives.Player{ID: int(playerid)}, int(forplayerid))
	return true
}

//export onVehicleStreamIn
func onVehicleStreamIn(vehicleid C.int, forplayerid C.int) bool {
	evt, ok := Events["vehicleStreamIn"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(int, int) bool)
	if !ok {
		return false
	}
	fn(int(vehicleid), int(forplayerid))
	return true
}

//export onVehicleStreamOut
func onVehicleStreamOut(vehicleid C.int, forplayerid C.int) bool {
	evt, ok := Events["vehicleStreamOut"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(int, int) bool)
	if !ok {
		return false
	}
	fn(int(vehicleid), int(forplayerid))
	return true
}

//export onActorStreamIn
func onActorStreamIn(actorid C.int, forplayerid C.int) bool {
	evt, ok := Events["actorStreamIn"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(int, int) bool)
	if !ok {
		return false
	}
	fn(int(actorid), int(forplayerid))
	return true
}

//export onActorStreamOut
func onActorStreamOut(actorid C.int, forplayerid C.int) bool {
	evt, ok := Events["actorStreamOut"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(int, int) bool)
	if !ok {
		return false
	}
	fn(int(actorid), int(forplayerid))
	return true
}

//export onDialogResponse
func onDialogResponse(playerid C.int, dialogid C.int, response C.int, listitem C.int, inputtext *C.char_t) bool {
	evt, ok := Events["dialogResponse"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(natives.Player, int, int, int, string) bool)
	if !ok {
		return false
	}
	return fn(natives.Player{ID: int(playerid)}, int(dialogid), int(response), int(listitem), C.GoString(C.constToNonConst(inputtext)))
}

//export onPlayerTakeDamage
func onPlayerTakeDamage(playerid C.int, issuerid C.int, amount C.float, weaponid C.int, bodypart C.int) bool {
	evt, ok := Events["playerTakeDamage"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(natives.Player, natives.Player, float32, int, int) bool)
	if !ok {
		return false
	}
	return fn(natives.Player{ID: int(playerid)}, natives.Player{ID: int(issuerid)}, float32(amount), int(weaponid), int(bodypart))
}

//export onPlayerGiveDamage
func onPlayerGiveDamage(playerid C.int, damagedid C.int, amount C.float, weaponid C.int, bodypart C.int) bool {
	evt, ok := Events["playerGiveDamage"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(natives.Player, int, float32, int, int) bool)
	if !ok {
		return false
	}
	return fn(natives.Player{ID: int(playerid)}, int(damagedid), float32(amount), int(weaponid), int(bodypart))
}

//export onPlayerGiveDamageActor
func onPlayerGiveDamageActor(playerid C.int, damaged_actorid C.int, amount C.float, weaponid C.int, bodypart C.int) bool {
	evt, ok := Events["playerGiveDamageActor"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(natives.Player, int, float32, int, int) bool)
	if !ok {
		return false
	}
	return fn(natives.Player{ID: int(playerid)}, int(damaged_actorid), float32(amount), int(weaponid), int(bodypart))
}

//export onPlayerClickMap
func onPlayerClickMap(playerid C.int, fX C.float, fY C.float, fZ C.float) bool {
	evt, ok := Events["playerClickMap"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(natives.Player, float32, float32, float32) bool)
	if !ok {
		return false
	}
	return fn(natives.Player{ID: int(playerid)}, float32(fX), float32(fY), float32(fZ))
}

//export onPlayerClickTextDraw
func onPlayerClickTextDraw(playerid C.int, clickedid C.int) bool {
	evt, ok := Events["playerClickTextDraw"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(natives.Player, int) bool)
	if !ok {
		return false
	}
	return fn(natives.Player{ID: int(playerid)}, int(clickedid))
}

//export onPlayerClickPlayerTextDraw
func onPlayerClickPlayerTextDraw(playerid C.int, playertextid C.int) bool {
	evt, ok := Events["playerClickPlayerTextDraw"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(natives.Player, int) bool)
	if !ok {
		return false
	}
	return fn(natives.Player{ID: int(playerid)}, int(playertextid))
}

//export onIncomingConnection
func onIncomingConnection(playerid C.int, ip_address *C.char_t, port C.int) bool {
	evt, ok := Events["incomingConnection"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(natives.Player, string, int) bool)
	if !ok {
		return false
	}
	return fn(natives.Player{ID: int(playerid)}, C.GoString(C.constToNonConst(ip_address)), int(port))
}

//export onTrailerUpdate
func onTrailerUpdate(playerid C.int, vehicleid C.int) bool {
	evt, ok := Events["trailerUpdate"]
	if !ok {
		return true
	}

	fn, ok := evt.Handler.(func(natives.Player, int) bool)
	if !ok {
		return true
	}
	return fn(natives.Player{ID: int(playerid)}, int(vehicleid))
}

//export onVehicleSirenStateChange
func onVehicleSirenStateChange(playerid C.int, vehicleid C.int, newstate C.int) bool {
	evt, ok := Events["vehicleSirenStateChange"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(natives.Player, int, int) bool)
	if !ok {
		return false
	}
	return fn(natives.Player{ID: int(playerid)}, int(vehicleid), int(newstate))
}

//export onPlayerClickPlayer
func onPlayerClickPlayer(playerid C.int, clickedplayerid C.int, source C.int) bool {
	evt, ok := Events["playerClickPlayer"]
	if !ok {
		return true
	}

	fn, ok := evt.Handler.(func(natives.Player, int, int) bool)
	if !ok {
		return true
	}
	fn(natives.Player{ID: int(playerid)}, int(clickedplayerid), int(source))
	return false
}

//export onPlayerEditObject
func onPlayerEditObject(playerid C.int, playerobject C.bool, objectid C.int, response C.int, fX C.float, fY C.float, fZ C.float, fRotX C.float, fRotY C.float, fRotZ C.float) bool {
	evt, ok := Events["playerEditObject"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(natives.Player, bool, int, int, float32, float32, float32, float32, float32, float32) bool)
	if !ok {
		return false
	}
	return fn(natives.Player{ID: int(playerid)}, bool(playerobject), int(objectid), int(response), float32(fX), float32(fY), float32(fZ), float32(fRotX), float32(fRotY), float32(fRotZ))
}

//export onPlayerEditAttachedObject
func onPlayerEditAttachedObject(playerid C.int, response C.int, index C.int, modelid C.int, boneid C.int, fOffsetX C.float, fOffsetY C.float, fOffsetZ C.float, fRotX C.float, fRotY C.float, fRotZ C.float, fScaleX C.float, fScaleY C.float, fScaleZ C.float) bool {
	evt, ok := Events["playerEditAttachedObject"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(natives.Player, int, int, int, int, float32, float32, float32, float32, float32, float32, float32, float32, float32) bool)
	if !ok {
		return false
	}
	return fn(natives.Player{ID: int(playerid)}, int(response), int(index), int(modelid), int(boneid), float32(fOffsetX), float32(fOffsetY), float32(fOffsetZ), float32(fRotX), float32(fRotY), float32(fRotZ), float32(fScaleX), float32(fScaleY), float32(fScaleZ))
}

//export onPlayerSelectObject
func onPlayerSelectObject(playerid C.int, type_ C.int, objectid C.int, modelid C.int, fX C.float, fY C.float, fZ C.float) bool {
	evt, ok := Events["playerSelectObject"]
	if !ok {
		return false
	}

	fn, ok := evt.Handler.(func(natives.Player, int, int, int, float32, float32, float32) bool)
	if !ok {
		return false
	}
	return fn(natives.Player{ID: int(playerid)}, int(type_), int(objectid), int(modelid), float32(fX), float32(fY), float32(fZ))
}

//export onPlayerWeaponShot
func onPlayerWeaponShot(playerid C.int, weaponid C.int, hittype C.int, hitid C.int, fX C.float, fY C.float, fZ C.float) bool {
	evt, ok := Events["playerWeaponShot"]
	if !ok {
		return true
	}

	fn, ok := evt.Handler.(func(natives.Player, int, int, int, float32, float32, float32) bool)
	if !ok {
		return true
	}
	return fn(natives.Player{ID: int(playerid)}, int(weaponid), int(hittype), int(hitid), float32(fX), float32(fY), float32(fZ))
}

//export onPlayerRequestDownload
func onPlayerRequestDownload(playerid C.int, type_ C.int, crc C.int) bool {
	evt, ok := Events["playerRequestDownload"]
	if !ok {
		return true
	}

	fn, ok := evt.Handler.(func(natives.Player, int, int) bool)
	if !ok {
		return true
	}
	return fn(natives.Player{ID: int(playerid)}, int(type_), int(crc))
}
