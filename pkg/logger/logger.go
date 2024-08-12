/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/go-logr/logr"
)

// TODO: DocString
func FromContext(gctx *gin.Context) logr.Logger {
	return gctx.MustGet("log").(logr.Logger)
}
