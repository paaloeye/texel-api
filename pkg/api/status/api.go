/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package status

import (
	"context"
	"net/http"
	"time"

	ginAPI "github.com/gin-gonic/gin"
	"github.com/paaloeye/texel-api/pkg/mnemosyne"
)

func Register(ginRouter *ginAPI.RouterGroup) {

	ginRouter.GET("/healthz", func(gin *ginAPI.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		model := gin.MustGet("model").(*mnemosyne.Mnemosyne)

		err := model.PingContext(ctx)
		if err != nil {
			gin.JSON(http.StatusFailedDependency, ginAPI.H{})
			return
		}

		gin.JSON(http.StatusOK, ginAPI.H{})
	})

}
