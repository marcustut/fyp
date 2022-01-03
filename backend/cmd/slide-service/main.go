package main

import (
	"fmt"
	"github.com/marcustut/fyp/backend/config"
	"github.com/marcustut/fyp/backend/internal/infrastructure/router"
	"log"
	"net/http"
)

func main() {
	config.ReadConfig(config.ReadConfigOption{})

	r := router.NewSlideRouter()

	log.Printf("slide-service running on port %d\n", config.C.Services.Slide.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.C.Services.Slide.Port), r))
}
