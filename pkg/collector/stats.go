package collector

import (
	"math"
	"github.com/BenjaminGlusa/powerlevel/pkg/adapter"
)


func roundFloat(val float32, precision uint) float64 {
    ratio := math.Pow(10, float64(precision))
    return math.Round(float64(val)*ratio) / ratio
}


func FetchSolarStats(db adapter.DatabaseAdapter)([]PowerStats, error) {
	var stats []PowerStats
	var curStats PowerStats

	// fixme: Implement db.KwhLatest()
	curStats.CurrentPower = 0
	curStats.YieldTotal = roundFloat(db.KwhToday(),3)
	curStats.YieldMonth = roundFloat(db.KwhThisMonth(),3)
	curStats.YieldYear = roundFloat(db.KwhThisYear(),3)
	curStats.YieldTotal = roundFloat(db.KwhTotal(),3)

	stats = append(stats, curStats)

	return stats, nil
}