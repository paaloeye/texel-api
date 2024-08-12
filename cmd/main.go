/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pbrit/texel-api/pkg/api/prometheus"
	"github.com/pbrit/texel-api/pkg/api/status"
	v1 "github.com/pbrit/texel-api/pkg/api/v1"
)

func main() {
	ginRouter := gin.Default()

	v1.Bind(ginRouter.Group("/v1"))
	prometheus.Bind(ginRouter.Group("/metrics"))
	status.Bind(ginRouter.Group("/status"))

	ginRouter.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "pong"})
	})

	ginRouter.Run()
}
