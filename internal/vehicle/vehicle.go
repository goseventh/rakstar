package vehicle

import (
	"fmt"
	"math"

	"github.com/goseventh/rakstar/internal/player"
	"github.com/goseventh/rakstar/internal/utils/constants/vehiclesConst"
	"github.com/goseventh/rakstar/internal/natives"
)

const (
	SeatDriver = iota
	SeatFront
	SeatBackLeft
	SeatBackRight
)

type Vehicle struct {
	ID int
}

type VehicleParams struct {
	Engine    int
	Lights    int
	Alarm     int
	Doors     int
	Bonnet    int
	Boot      int
	Objective int
}

func NewVehicle(modelid int, x, y, z, rotation float32, color1, color2 uint8, respawn_delay int, addsiren bool) (Vehicle, error) {
	var v Vehicle
	if !natives.IsValidVehicleModel(modelid) {
		return v, fmt.Errorf("invalid vehicle model")
	}
	v.ID = natives.CreateVehicle(modelid, x, y, z, rotation, int(color1), int(color2), respawn_delay, addsiren)
	if v.ID == vehiclesConst.InvalidVehicleId {
		return v, fmt.Errorf("couldn't create vehicle")
	}
	return v, nil
}

func (v *Vehicle) Destroy() error {

	if !natives.DestroyVehicle(v.ID) {
		return fmt.Errorf("vehicle doesn't exist")
	}
	return nil
}

func (v *Vehicle) GetSpeedFloat64() float64 {
	var x, y, z float32
	natives.GetVehicleVelocity(v.ID, &x, &y, &z)

	return math.Sqrt(float64((x*x)+(y*y)+(z*z))) * 136.666667
}

func (v *Vehicle) GetSpeedFloat32() float32 {
	return float32(v.GetSpeedFloat64())
}

func (v *Vehicle) GetSpeedInt() int {
	return int(math.Round(v.GetSpeedFloat64()))
}

func (v *Vehicle) PutPlayer(p *player.Player, seat int) error {
	if !natives.PutPlayerInVehicle(p.ID, v.ID, seat) {
		return fmt.Errorf("player or vehicle doesn't exist")
	}
	return nil
}

func (v *Vehicle) GetParams() VehicleParams {
	var params VehicleParams
	natives.GetVehicleParamsEx(v.ID, &params.Engine, &params.Lights, &params.Alarm, &params.Doors, &params.Bonnet, &params.Boot, &params.Objective)
	return params
}

func (v *Vehicle) SetParams(params VehicleParams) {
	natives.SetVehicleParamsEx(v.ID, params.Engine, params.Lights, params.Alarm, params.Doors, params.Bonnet, params.Boot, params.Objective)
}
