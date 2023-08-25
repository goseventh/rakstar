package vehicles

import (
	mongodb "github.com/goseventh/rakstar/database/mongo"
	"github.com/goseventh/rakstar/internal/natives"
	"go.mongodb.org/mongo-driver/bson"
)

const vehicleStatesColl = "vehicle_states"

type vehicleState struct {
	model                                                          int
	doorDriver, doorPassenger, doorBackLeft, doorBackRight         int
	windowDriver, windowPassenger, windowBackLeft, windowBackRight int
	damagePanels, damageDoors, damageLights, damageTires           int
	engine, lights, alarm, doors, bonnet, boot, objective          int
	vehicleTrailer                                                 int
	health                                                         float32
	x, y, z, r                                                     float32
	quatW, quatX, quatY, quatZ                                     float32
	fuel, batteryCharger                                           float32
	componentSlots                                                 []int
	colors                                                         []int
	plate                                                          string
}

func (v *vehicleBuilder) saveState() {
	v.state = new(vehicleState)

	v.state.model = v.model
	v.state.plate = v.plate
	v.state.colors[0] = v.colorPrimary
	v.state.colors[1] = v.colorSecondary

	natives.GetVehicleHealth(v.id, &v.state.health)
	natives.GetVehiclePos(v.id, &v.posX, &v.posY, &v.posZ)
	natives.GetVehicleRotationQuat(v.id, &v.state.quatW, &v.state.quatX, &v.state.quatY, &v.state.quatZ)
	natives.GetVehicleParamsCarDoors(v.id, &v.state.doorDriver, &v.state.doorPassenger, &v.state.doorBackLeft, &v.state.doorBackRight)
	natives.GetVehicleParamsCarWindows(v.id, &v.state.windowDriver, &v.state.windowPassenger, &v.state.windowBackLeft, &v.state.windowBackRight)
	natives.GetVehicleDamageStatus(v.id, &v.state.damagePanels, &v.state.damageDoors, &v.state.damageLights, &v.state.damageTires)

	for i := 0; i < 13; i++ {
		v.state.componentSlots = append(v.state.componentSlots, natives.GetVehicleComponentInSlot(v.id, i))
	}

	mongodb.
		Builder().
		UseDatabase("rk_state").
		UseCollection(vehicleStatesColl).
		CreateVehicleStateWithWorkers(v.state)
}

func (v *vehicleBuilder) loadState() {
	state := new(vehicleState)

	err := mongodb.
		Builder().
		UseDatabase("rk_state").
		UseCollection(vehicleStatesColl).
		GetVehicleState(state, bson.D{})

	v.state = state

	natives.SetVehicleHealth(v.id, v.state.health)
	natives.SetVehiclePos(v.id, v.posX, v.posY, v.posZ)
	natives.SetVehicleParamsCarDoors(v.id, v.state.doorDriver, v.state.doorPassenger, v.state.doorBackLeft, v.state.doorBackRight)
	natives.SetVehicleParamsCarWindows(v.id, v.state.windowDriver, v.state.windowPassenger, v.state.windowBackLeft, v.state.windowBackRight)
	natives.UpdateVehicleDamageStatus(v.id, v.state.damagePanels, v.state.damageDoors, v.state.damageLights, v.state.damageTires)

	for i := 0; i < 13; i++ {
		natives.AddVehicleComponent(v.id, v.state.componentSlots[i])
	}

	if err != nil {
		panic(err)
	}
}
