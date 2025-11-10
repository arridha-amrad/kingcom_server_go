package main

import (
	"kingcom_api/internal/modules"

	"go.uber.org/fx"
)

func main() {
	fx.New(modules.CommonModule).Run()
}
