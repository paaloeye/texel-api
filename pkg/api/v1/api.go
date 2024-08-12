/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Project struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func Bind(ginRouter *gin.RouterGroup) {
	api := ginRouter.Group("/projects/:project_id")

	// * Object: building_limits
	//   Methods: PUT, GET
	api.GET("/building_limits", func(g *gin.Context) { g.JSON(http.StatusOK, "helloworld") })
	api.PUT("/building_limits", func(g *gin.Context) { g.JSON(http.StatusOK, "helloworld") })

	// * Object: height_plateaus
	//   Methods: PUT, GET
	api.GET("/height_plateaus", func(g *gin.Context) { g.JSON(http.StatusOK, "helloworld") })
	api.PUT("/height_plateaus", func(g *gin.Context) { g.JSON(http.StatusOK, "helloworld") })

	// * Object: split_building_limits
	//   Methods: GET
	api.GET("/split_building_limits", func(g *gin.Context) { g.JSON(http.StatusOK, "helloworld") })
}
