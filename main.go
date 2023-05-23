package main 

import (
	"log"
	"net/http"
	"github.com/pbnjay/memory"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	memoriaLivreBytesGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "memoria_livre_bytes",
		Help: "Quantidade de memoria livre em bytes",
	})
	
	memoriaTotalBytesGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "memoria_total_bytes",
		Help: "Quantidade de memoria total em bytes",
	})

	memoriaTotalGigasGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "memoria_total_gigas",
		Help: "Quantidade de memoria total em gigas",
	})

	memoriaLivreMegasGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "memoria_livre_megas",
		Help: "Quantidade de memoria livre em megas",
	})
)

func memoriaLivre() float64 {
	memoria_livre := memory.FreeMemory()
	return float64(memoria_livre)
}

func totalMemoria() float64 {
	memoria_total := memory.TotalMemory()
	return float64(memoria_total)
}


func init() { // executing before main
		prometheus.MustRegister(memoriaLivreBytesGauge)
		prometheus.MustRegister(memoriaLivreMegasGauge)
		prometheus.MustRegister(memoriaTotalBytesGauge)
		prometheus.MustRegister(memoriaTotalGigasGauge)
}


func main () {
	memoriaLivreBytesGauge.Set(memoriaLivre())
	memoriaLivreMegasGauge.Set(memoriaLivre() / 1024 / 1024)
	memoriaTotalBytesGauge.Set(totalMemoria())
	memoriaTotalGigasGauge.Set(totalMemoria() / 1024 / 1024 / 1024)
	http.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServe(":7788", nil))

}

