package handlers

import "github.com/Rikypurnomo/warmup/internal/api/services"

type (
	handlesInit struct {
		service *services.ServicessInit
	}
)

func InitiateHandlersInterface() *handlesInit {
	return &handlesInit{
		service: services.InitiateServicessInterface(),
	}
}
