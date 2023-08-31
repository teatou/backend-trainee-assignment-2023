package myapi

import (
	"backend-trainee-assignment-2023/internal/mylogic"
	"backend-trainee-assignment-2023/pkg/mylogger"
)

type Api struct {
	logic  *mylogic.Logic
	logger mylogger.Logger
}

func NewApi(logic *mylogic.Logic, logger mylogger.Logger) *Api {
	return &Api{
		logger: logger,
		logic:  logic,
	}
}
