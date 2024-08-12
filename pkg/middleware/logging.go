/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-logr/zapr"
	"github.com/pbrit/texel-api/pkg/errors"
	"go.uber.org/zap"
)

func Logging(zap *zap.Logger) gin.HandlerFunc {

	return func(gctx *gin.Context) {
		log := zapr.NewLogger(zap)
		gctx.Set("log", log)

		start := time.Now()

		gctx.Next()

		elapsedDuration := time.Since(start)

		keysAndValues := []any{
			"method", gctx.Request.Method,
			"path", gctx.Request.RequestURI,
			"status", gctx.Writer.Status(),
			"referrer", gctx.Request.Referer(),
			"duration", elapsedDuration,
		}

		// 	"client_ip":  util.GetClientIP(c),
		// 	"user_id":    util.GetUserID(c),
		// 	"api_version": util.ApiVersion,
		// 	"request_id": c.Writer.Header().Get("Request-Id"),

		// Make sure errors are being taken care of and return
		if gctx.Writer.Status() >= http.StatusInternalServerError {
			log.Error(errors.ErrInternalServer, "", keysAndValues...)
			return
		}

		log.Info("", keysAndValues...)
	}
}
