/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package project

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pbrit/texel-api/pkg/app"
	"github.com/pbrit/texel-api/pkg/logger"
)

func Register(ginRouter *gin.RouterGroup) {
	api := ginRouter.Group("/projects/:project_id")

	// Bind project ID
	api.Use(projectIDMiddleware)

	// * Object: building_limits
	//   Methods: PUT, GET
	api.GET("/building_limits", func(gctx *gin.Context) {
		log := logger.FromContext(gctx)
		project := gctx.MustGet("project").(Project)
		app := gctx.MustGet("app").(*app.App)

		buildingLimits, err := app.Mnemosyne.GetBuildingLimits()

		gctx.JSON(http.StatusOK, gin.H{"projectUUID": project.ID})

	})
	api.PUT("/building_limits", func(g *gin.Context) { g.JSON(http.StatusOK, "helloworld") })

	// * Object: height_plateaus
	//   Methods: PUT, GET
	api.GET("/height_plateaus", func(g *gin.Context) { g.JSON(http.StatusOK, "helloworld") })
	api.PUT("/height_plateaus", func(g *gin.Context) { g.JSON(http.StatusOK, "helloworld") })

	// * Object: split_building_limits
	//   Methods: GET
	api.GET("/split_building_limits", func(g *gin.Context) { g.JSON(http.StatusOK, "helloworld") })
}

func projectIDMiddleware(gctx *gin.Context) {
	var project Project

	log := logger.FromContext(gctx)

	err := gctx.ShouldBindUri(&project)
	if err != nil {
		gctx.JSON(400, gin.H{"msg": "bla"})
		return
	}

	gctx.Set("project", project)
	gctx.Set("log", log.WithValues("project-id", project.ID))

	gctx.Next()
}
