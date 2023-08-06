package vehicle

import (
	// "fmt"
	"math"
	"math/rand"
	"testing"
	// "time"
)

func TestFuelMinMax(t *testing.T) {
	veh := Builder()
	for i := 0; i < 483_647; i++ {
		veh.Engine().
			Fuel(float32(rand.Intn(math.MaxInt16 - (math.MinInt16) + math.MaxInt16)))

		if veh.engine.fuel > 100 {
			t.Errorf("expected: 100; got: %v", veh.engine.fuel)
		}
		if veh.engine.fuel < 0 {
			t.Errorf("expected: 0; got: %v", veh.engine.fuel)
		}
	}
}

func TestFuelEconomyMinMax(t *testing.T) {
	veh := Builder()
	for i := 0; i < 483_647; i++ {
		veh.Engine().
			FuelEconomy(float32(rand.Intn(math.MaxInt16 - (math.MinInt16) + math.MaxInt16)))

		if veh.engine.fuelEconomy > 100 {
			t.Errorf("expected: 100; got: %v", veh.engine.fuelEconomy)
		}
		if veh.engine.fuelEconomy < 0 {
			t.Errorf("expected: 0; got: %v", veh.engine.fuelEconomy)
		}
	}
}

func TestSetGetFuel(t *testing.T) {
	veh := Builder()
	for f := 0; f < 100; f++ {
		veh.Engine().Fuel(float32(f))
		if veh.Engine().GetFuel() == float32(f) {
			continue
		}
		t.Errorf("expected: %v; got: %v", f, veh.engine.fuel)
	}

}

func TestSetGetFuelEconomy(t *testing.T) {
	veh := Builder()
	var f float32
	for f = 0; f < 100; f++ {
		veh.Engine().FuelEconomy(f)
		if veh.Engine().GetFuelEconomy() == f {
			continue
		}
		t.Errorf("expected: %v; got: %v", f, veh.engine.fuelEconomy)
	}

}

func TestSetGetBatteryCharger(t *testing.T) {
	veh := Builder()
	veh.Eletrics().BatteryCharger(7)

	for f := 0; f < 100; f++ {
		veh.Eletrics().BatteryCharger(float32(f))

		if veh.Eletrics().GetBatteryCharger() == float32(f) {
			continue
		}
		t.Errorf("expected: %v; got: %v", f, veh.eletrics.batteryCharger)

	}

}

func TestSortIgnite(t *testing.T) {
	rounds := 7000
	for charger := float32(0); charger <= 100; charger+=5{
		for fuel := float32(0); fuel <= 100; fuel+=5{
			veh := Builder()
			starts := 0
			for i := 0; i < rounds; i++ {
				veh.Eletrics().BatteryCharger(charger)
				veh.Engine().Fuel(fuel)
				stared := veh.Engine().canIgniteEngine()
				if !stared {
					continue
				}
				starts += 1
			}

			switch {
			case veh.Engine().GetFuel() <= 90 &&
				veh.Engine().GetFuel() >= 85 ||
				veh.Eletrics().GetBatteryCharger() <= 100 &&
					veh.Eletrics().GetBatteryCharger() >= 80:

				if (float64(starts) / float64(rounds) * 100) <= 100 &&
        (float64(starts)/float64(rounds)*100 >= 85){
					continue
				}
				t.Errorf("expected: ≈90%%; got: %0.1f%%", (float64(starts) / float64(rounds) * 100))

			case veh.Engine().GetFuel() <= 40 &&
				veh.Engine().GetFuel() >= 30 ||
				veh.Eletrics().GetBatteryCharger() <= 50 &&
					veh.Eletrics().GetBatteryCharger() >= 35:

				if (float64(starts) / float64(rounds) * 100) <= 65&&
        (float64(starts)/float64(rounds)*100) >= 55 {
					continue
				}

				t.Errorf("expected: ≈60%%; got: %0.1f%%", (float64(starts) / float64(rounds) * 100))

			case veh.Engine().GetFuel() <= 20 &&
				veh.Engine().GetFuel() >= 15 ||
				veh.Eletrics().GetBatteryCharger() <= 30 &&
					veh.Eletrics().GetBatteryCharger() >= 25:

				if (float64(starts) / float64(rounds) * 100) <= 35 &&
        (float64(starts)/(float64(rounds)*100) >= 25){
					continue
				}

				t.Errorf("expected: ≈30%%; got: %0.1f%%", (float64(starts) / float64(rounds) * 100))

			case veh.Engine().GetFuel() <= 10 &&
				veh.Engine().GetFuel() >= 1 ||
				veh.Eletrics().GetBatteryCharger() <= 20 &&
					veh.Eletrics().GetBatteryCharger() >= 1:

				if float64(starts)/float64(rounds)*100 <= 20 &&
        float64(starts)/float64(rounds)*100 >= 10{
					continue
				}
				t.Errorf("expected: ≈15%%; got: %0.1f%%", (float64(starts) / float64(rounds) * 100))
			}
		}
	}
}
