package collector

import (
	"github.com/prometheus/client_golang/prometheus"
)


type PowerStats struct {
	CurrentPower float64
	YieldToday float64
	YieldMonth float64
	YieldYear float64
	YieldTotal float64
}


type powerCollector struct {
	CurrentPower *prometheus.Desc
	YieldToday *prometheus.Desc
	YieldMonth *prometheus.Desc
	YieldYear *prometheus.Desc
	YieldTotal *prometheus.Desc

	stats func() ([]PowerStats, error)
}

func NewPowerCollector(stats func() ([]PowerStats, error)) prometheus.Collector {
	return &powerCollector {
		CurrentPower: prometheus.NewDesc(
			"powerlevel_current_p",
			"Amount of wattage currently produced in Watt",
			[]string{},
			nil,
		),
		YieldToday: prometheus.NewDesc(
			"powerlevel_today_p",
			"Amount of wattage produced today in kWh",
			[]string{},
			nil,
		),
		YieldMonth: prometheus.NewDesc(
			"powerlevel_month_p",
			"Amount of wattage produced this month in kWh",
			[]string{},
			nil,
		),
		YieldYear: prometheus.NewDesc(
			"powerlevel_year_p",
			"Amount of wattage produced this year in kWh",
			[]string{},
			nil,
		),
		YieldTotal: prometheus.NewDesc(
			"powerlevel_total_e",
			"Amount of wattage produced in total in kWh",
			[]string{},
			nil,
		),
		stats: stats,
	}
}

func (p *powerCollector) Describe(ch chan<- *prometheus.Desc) {
	ds := []*prometheus.Desc{
		p.CurrentPower,
	}

	for _, d := range ds {
		ch <- d
	}
}


func (p *powerCollector) Collect(ch chan<- prometheus.Metric) {
	stats, err := p.stats()
	if err != nil {
		ch <- prometheus.NewInvalidMetric(p.CurrentPower, err)
		return
	}

	for _, s := range stats {
		ch <- prometheus.MustNewConstMetric(
			p.CurrentPower,
			prometheus.GaugeValue,
			s.CurrentPower,
		)

		ch <- prometheus.MustNewConstMetric(
			p.YieldToday,
			prometheus.GaugeValue,
			s.YieldToday,
		)

		
		ch <- prometheus.MustNewConstMetric(
			p.YieldMonth,
			prometheus.GaugeValue,
			s.YieldMonth,
		)
		
		ch <- prometheus.MustNewConstMetric(
			p.YieldYear,
			prometheus.GaugeValue,
			s.YieldYear,
		)

		ch <- prometheus.MustNewConstMetric(
			p.YieldTotal,
			prometheus.GaugeValue,
			s.YieldTotal,
		)
	}
}