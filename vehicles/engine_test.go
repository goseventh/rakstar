package vehicles

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
	rounds := 50_000
	for charger := float32(0); charger <= 100; charger += 5 {
		for fuel := float32(0); fuel <= 100; fuel += 5 {
			veh := Builder()
			igniteCount := 0

			for i := 0; i < rounds; i++ {
				veh.Eletrics().BatteryCharger(charger)
				veh.Engine().Fuel(fuel)

				canIgnite := veh.Engine().canIgniteEngine()

				if !canIgnite {
					continue
				}

				igniteCount += 1
			}

			successAverage := float32(igniteCount) / float32(rounds) * 100
			minFuelCharger := math.Min(
				float64(veh.Engine().GetFuel()),
				float64(veh.Eletrics().GetBatteryCharger()),
			)

			if minFuelCharger == 90 {
				if successAverage >= 85 && successAverage <= 95 {
					continue
				}

				t.Errorf("expected: ≈90%%; got: %0.1f%%", successAverage)
      }

			if minFuelCharger == 40 {
				if successAverage >= 35 && successAverage <= 45 {
					continue
				}

				t.Errorf("expected: ≈40%%; got: %0.1f%%", successAverage)
			}

			if minFuelCharger == 20 {
				if successAverage >= 15 && successAverage <= 25 {
					continue
				}

				t.Errorf("expected: ≈20%%; got: %0.1f%%", successAverage)
			}

			if minFuelCharger == 10 {
				if successAverage >= 5 && successAverage <= 15 {
					continue
				}

				t.Errorf("expected: ≈10%%; got: %0.1f%%", successAverage)
			}
		}
	}
}
