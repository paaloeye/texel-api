/*
* This Source Code Form is subject to the terms of the Mozilla Public
* License, v. 2.0. If a copy of the MPL was not distributed with this
* file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package app

import (
	"github.com/gin-gonic/gin"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/pbrit/texel-api/pkg/api/prometheus"
	"github.com/pbrit/texel-api/pkg/api/status"
	projectControllerV1 "github.com/pbrit/texel-api/pkg/controller/v1/project"
	"github.com/pbrit/texel-api/pkg/middleware"
	"github.com/pbrit/texel-api/pkg/mnemosyne"
)

type App struct {
	gin *gin.Engine
	zap *zap.Logger

	Mnemosyne *mnemosyne.Mnemosyne
}

// Configure the app and run it
// NB: It blocks current coroutine
func ConfigureAppAndRun() {
	app := &App{}
	app.gin = gin.New()

	// Configure logging
	zapConfig := zap.NewDevelopmentConfig()
	zapConfig.Level = zap.NewAtomicLevelAt(zapcore.Level(-3)) // Show everything higher V=3

	zap, err := zapConfig.Build()
	if err != nil {
		panic(err)
	}
	app.zap = zap

	log := zapr.NewLogger(app.zap)

	// Configure persistance layer
	app.Mnemosyne = mnemosyne.New(log)
	defer app.Mnemosyne.Drop()

	// Configure all required middlewares
	app.gin.Use(middleware.Logging(app.zap))
	app.gin.Use(gin.Recovery())
	app.gin.Use(modelMiddleware(app))

	projectControllerV1.Register(app.gin.Group("/v1"))
	// projectControllerV2.Register(app.gin.Group("/v2"))

	prometheus.Bind(app.gin.Group("/metrics"))
	status.Bind(app.gin.Group("/status"))

	if err := app.gin.Run(); err != nil {
		panic(err)
	}

	// Finish the app peacefully
}

func modelMiddleware(app *App) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		gctx.Set("model", app.Mnemosyne)
		gctx.Next()
	}
}
