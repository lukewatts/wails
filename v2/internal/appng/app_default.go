//go:build !windows && !dev

package appng

import (
	"github.com/wailsapp/wails/v2/internal/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
)

func PreflightChecks(_ *options.App, _ *logger.Logger) error {
	return nil
}

func (a *App) Init() {
}