package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/BenjaminGlusa/powerlevel/pkg/adapter"
	"github.com/BenjaminGlusa/powerlevel/pkg/api"
	"github.com/BenjaminGlusa/powerlevel/pkg/collector"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func MakeAddMeasurementHandler(db adapter.DatabaseAdapter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		var addMeasurementRequest api.AddMeasurementRequest
		if err = json.Unmarshal(body, &addMeasurementRequest); err != nil {
      		w.WriteHeader(http.StatusBadRequest)
			log.Printf("could not add measurement: %v\n", err)
      		return
    	}

		err = db.AddMeasurement(addMeasurementRequest.Watt)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("could not add measurement: %v\n", err)
      		return
		}
		w.WriteHeader(http.StatusAccepted)
		
	}
}



func MakeMetricsHandler(db adapter.DatabaseAdapter) http.Handler {
	powerStats := func()([]collector.PowerStats, error) {
		return collector.FetchSolarStats(db)
	}

	sc := collector.NewPowerCollector(powerStats)
	reg := prometheus.NewRegistry()
	reg.MustRegister(sc)

	promHandler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{})
	return promHandler
}