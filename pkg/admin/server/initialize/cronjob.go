package initialize

import (
	"affiliate/internal/config"
	"affiliate/internal/middleware/schedule"
	"log"
	"net/http"
)

func runHttpGet() {
	cfg := config.GetENV()
	resp, err := http.Get(cfg.CRON_URL)
	if err != nil {
		log.Printf("Failed to send GET request: %v", err)
		return
	}
	defer resp.Body.Close()
	log.Printf("GET request to be affiliate succeeded with status code %d", resp.StatusCode)
}

func RunCronJob() {
	jobs := []*schedule.Job{
		{
			Spec: "0 */8 * * * *",
			Name: "GET request to be affiliate",
			Cmd: func() {
				runHttpGet()
			},
		},
	}
	s := schedule.New(jobs...)
	s.Start()
}
