/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package project

import (
	"context"
	"io"
	"net/http"

	ginAPI "github.com/gin-gonic/gin"
	"github.com/go-logr/logr"
	"github.com/pbrit/texel-api/pkg/construction"
	"github.com/pbrit/texel-api/pkg/logger"
	"github.com/pbrit/texel-api/pkg/mnemosyne"

	"github.com/paulmach/orb/geojson"
)

type ContextKey string

const (
	ctxKeyLogger           ContextKey = `looger`  // type: logr.Logger
	ctxKeyGin              ContextKey = `gin`     // type: *gin.Context
	ctxKeyDesignRuleEngine ContextKey = `dre`     // type: *construction.DesignRuleEngine
	ctxKeyProject          ContextKey = `project` // type: Project
	ctxKeyModel            ContextKey = `model`   // type: *mnemosyne.Mnemosyne
)

func Register(ginRouter *ginAPI.RouterGroup) {
	api := ginRouter.Group("/projects/:project_id")

	// Bind project ID
	api.Use(projectIDMiddleware)

	api.GET("/building_limits", func(gin *ginAPI.Context) {
		log := logger.FromContext(gin)
		project := gin.MustGet("project").(Project)
		model := gin.MustGet("model").(*mnemosyne.Mnemosyne)

		// Context business logic
		ctx := context.Background()
		ctx = context.WithValue(ctx, ctxKeyLogger, log)
		ctx = context.WithValue(ctx, ctxKeyGin, gin)

		buildingLimits, err := model.GetBuildingLimits(project.ID)

		if ok := handleNotFound(ctx, err); !ok {
			return
		}

		log.V(4).Info("building limits found", "object", buildingLimits)

		// Make sure it's a well-formatted GeoJSON Object
		geoJsonObj, err := geojson.UnmarshalFeatureCollection([]byte(buildingLimits))
		if ok := handleInternalServerError(ctx, err); !ok {
			return
		}

		gin.JSON(http.StatusOK, ginAPI.H{"data": *geoJsonObj})
	})
	api.PATCH("/building_limits", func(gin *ginAPI.Context) {
		ctx := makeUpdateContext(gin)
		project := ctx.Value(ctxKeyProject).(Project)
		model := ctx.Value(ctxKeyModel).(*mnemosyne.Mnemosyne)
		dre := ctx.Value(ctxKeyDesignRuleEngine).(*construction.DesignRuleEngine)

		body, err := io.ReadAll(gin.Request.Body)
		if ok := handleInternalServerError(ctx, err); !ok {
			return
		}

		// Make sure it's a well-formatted GeoJSON Object
		featureCollectionRequest, err := geojson.UnmarshalFeatureCollection(body)
		if ok := handleInternalServerError(ctx, err); !ok {
			return
		}

		// Check design rules
		warnCollection, errCollection, err := dre.Validate(ctx, featureCollectionRequest, featureCollectionComplementary)

		// Serialize errors and warnings

		geoJson, err := featureCollectionRequest.MarshalJSON()
		if ok := handleInternalServerError(ctx, err); !ok {
			return
		}

		err = model.UpdateBuildingLimits(project.ID, string(geoJson[:]))
		if ok := handleInternalServerError(ctx, err); !ok {
			return
		}

		gin.JSON(http.StatusOK, ginAPI.H{
			"data": *featureCollectionRequest,
		})
	})

	api.GET("/height_plateaus", func(gin *ginAPI.Context) {
		log := logger.FromContext(gin)
		project := gin.MustGet("project").(Project)
		model := gin.MustGet("model").(*mnemosyne.Mnemosyne)

		// Context business logic
		ctx := context.Background()
		ctx = context.WithValue(ctx, ctxKeyLogger, log)
		ctx = context.WithValue(ctx, ctxKeyGin, gin)

		heightPlateaux, err := model.GetHeightPlateaux(project.ID)

		if ok := handleNotFound(ctx, err); !ok {
			return
		}

		log.V(4).Info("height plateaux found", "object", heightPlateaux)

		// Make sure it's a well-formatted GeoJSON Object
		geoJsonObj, err := geojson.UnmarshalFeatureCollection([]byte(heightPlateaux))
		if ok := handleInternalServerError(ctx, err); !ok {
			return
		}

		gin.JSON(http.StatusOK, ginAPI.H{"data": *geoJsonObj})
	})

	api.PATCH("/height_plateaus", func(gin *ginAPI.Context) {
		ctx := makeUpdateContext(gin)
		project := ctx.Value(ctxKeyProject).(Project)
		model := ctx.Value(ctxKeyModel).(*mnemosyne.Mnemosyne)

		body, err := io.ReadAll(gin.Request.Body)
		if ok := handleInternalServerError(ctx, err); !ok {
			return
		}

		// Make sure it's a well-formatted GeoJSON Object
		geoJsonBody, err := geojson.UnmarshalFeatureCollection(body)
		if ok := handleInternalServerError(ctx, err); !ok {
			return
		}

		geoJson, err := geoJsonBody.MarshalJSON()
		if ok := handleInternalServerError(ctx, err); !ok {
			return
		}

		// TODO: Check design rules

		err = model.UpdateHeightPlateaux(project.ID, string(geoJson[:]))
		if ok := handleInternalServerError(ctx, err); !ok {
			return
		}

		gin.JSON(http.StatusOK, ginAPI.H{
			"data": *geoJsonBody,
		})
	})

	api.GET("/split_building_limits", func(gin *ginAPI.Context) {
		log := logger.FromContext(gin)
		project := gin.MustGet("project").(Project)
		model := gin.MustGet("model").(*mnemosyne.Mnemosyne)

		// Context business logic
		ctx := context.Background()
		ctx = context.WithValue(ctx, ctxKeyLogger, log)
		ctx = context.WithValue(ctx, ctxKeyGin, gin)

		split_building_limits, err := model.GetSplitBuildingLimits(project.ID)

		if ok := handleNotFound(ctx, err); !ok {
			return
		}

		log.V(4).Info("split building limits found", "object", split_building_limits)

		// Make sure it's a well-formatted GeoJSON Object
		geoJsonObj, err := geojson.UnmarshalFeatureCollection([]byte(split_building_limits))
		if ok := handleInternalServerError(ctx, err); !ok {
			return
		}

		gin.JSON(http.StatusOK, ginAPI.H{"data": *geoJsonObj})
	})
}

// MARK: Private API

// Error handling

/*
 * @summary Handles internal server errors and logs the error to the logger.
 * @param ctx The context of the request.
 * @param err The error that occurred.
 * @return A boolean indicating whether the error was handled or not.
 */
func handleInternalServerError(ctx context.Context, err error) bool {
	log := ctx.Value(ctxKeyLogger).(logr.Logger)
	gin := ctx.Value(ctxKeyGin).(*ginAPI.Context)

	if err != nil {
		log.Error(err, "failed to get the model")
		gin.JSON(http.StatusInternalServerError, ginAPI.H{
			"error": ginAPI.H{
				"code":    http.StatusInternalServerError,
				"message": err.Error(),
			},
		})

		return false
	}

	return true
}

/**
 * @summary Handles not found errors and logs the error to the logger.
 * @param ctx The context of the request.
 * @param err The error that occurred.
 * @return A boolean indicating whether the error was handled or not.
 */
func handleNotFound(ctx context.Context, err error) bool {
	log := ctx.Value(ctxKeyLogger).(logr.Logger)
	gin := ctx.Value(ctxKeyGin).(*ginAPI.Context)

	if err == nil {
		return true
	}

	// Filter ErrNotFound first
	if err == mnemosyne.ErrNotFound {
		log.V(3).Info("object not found")

		gin.JSON(http.StatusOK, ginAPI.H{})
		return false
	}

	// Some other error occurred
	return handleInternalServerError(ctx, err)
}

// MARK: Middlewares

func projectIDMiddleware(gin *ginAPI.Context) {
	var project Project

	log := logger.FromContext(gin)

	err := gin.ShouldBindUri(&project)
	if err != nil {
		gin.JSON(400, ginAPI.H{"msg": "bla"})
		return
	}

	gin.Set("project", project)
	gin.Set("log", log.WithValues("project-id", project.ID))

	gin.Next()
}

func makeUpdateContext(gin *ginAPI.Context) context.Context {
	log := logger.FromContext(gin)
	project := gin.MustGet("project").(Project)
	model := gin.MustGet("model").(*mnemosyne.Mnemosyne)
	dre := construction.NewDesignRuleEngine()

	// Context business logic
	ctx := context.Background()
	ctx = context.WithValue(ctx, ctxKeyLogger, log)
	ctx = context.WithValue(ctx, ctxKeyGin, gin)
	ctx = context.WithValue(ctx, ctxKeyProject, project)
	ctx = context.WithValue(ctx, ctxKeyModel, model)
	ctx = context.WithValue(ctx, ctxKeyDesignRuleEngine, dre)

	return ctx
}
