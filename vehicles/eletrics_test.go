package vehicle
import (
	"math"
  "math/rand"
	"testing"
)

func TestBatteryMinMax(t *testing.T) {
  veh := Builder()
  for i:=0; i < 483_647; i++ {
		veh.Eletrics().
			BatteryCharger(float32(rand.Intn(math.MaxInt16- (math.MinInt16) + math.MaxInt16)))
		if veh.eletrics.batteryCharger > 100 {
			t.Errorf("expected: 100; got: %v", veh.eletrics.batteryCharger)
		}
		if veh.eletrics.batteryCharger < 0 {
			t.Errorf("expected: 0; got: %v", veh.eletrics.batteryCharger)
		}
	}
}


