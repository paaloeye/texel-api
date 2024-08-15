/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package construction

import (
	"errors"
	"fmt"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/clip"
	"github.com/paulmach/orb/geojson"
	"github.com/paulmach/orb/planar"
)

//go:generate stringer -type=DesignRuleViolation -output=zz_generated_dre.stringer.go
const (
	DesignRuleViolationPOfBound DesignRuleViolation = iota
	DesignRuleViolationPConflict
	DesignRuleViolationPCoverage // Warning

	DesignRuleViolationOverlapped
)

var (
	rules map[DesignRuleViolation]DesignRuleFuncOne = map[DesignRuleViolation]DesignRuleFuncOne{}
)

func init() {
	designRuleRegister(DesignRuleViolationOverlapped,
		func(featureCollection *geojson.FeatureCollection) (ok bool) {
			polygons := []orb.Polygon{}

			for _, f := range featureCollection.Features {
				polygons = append(polygons, f.Geometry.(orb.Polygon))
			}

			for i := 0; i < len(polygons); i++ {
				for j := i + 1; j < len(polygons); j++ {
					if overlapped := polygonsOverlapped(polygons[i], polygons[j]); overlapped {
						return false
					}
				}
			}

			return true
		})

}

type DesignRuleViolation int

type DesignRuleEngine struct {
}

// type DesignRuleViolation struct {
// }

func NewDesignRuleEngine() *DesignRuleEngine {
	return &DesignRuleEngine{}
}

type DesignRuleFuncOne func(featureCollection *geojson.FeatureCollection) (ok bool)
type DesignRuleFuncMany func(featureCollectionA, featureCollectionB *geojson.FeatureCollection) (ok bool)

func (dre *DesignRuleEngine) ValidateCollection(featureCollection *geojson.FeatureCollection) (ok bool, violations []error) {
	for rule, ruleFunc := range rules {
		if ok := ruleFunc(featureCollection); !ok {
			violations = append(violations, errors.New(rule.String()))
		}
	}

	if len(violations) != 0 {
		ok = false
		return
	}

	return true, nil
}

// MARK: Private API

func designRuleRegister(rule DesignRuleViolation, ruleFunc DesignRuleFuncOne) {
	if _, ok := rules[rule]; ok {
		panic(fmt.Errorf("%+v is already registered", rule))
	}
	rules[rule] = ruleFunc
}

func polygonsOverlapped(polygonA, polygonB orb.Polygon) bool {
	// Check if any part of polygonA intersects with polygonB
	clipped := clip.Polygon(polygonB.Bound(), polygonA.Clone())
	if len(clipped) > 0 {
		return true
	}

	// Check if polygonB is entirely within polygonA
	if planar.PolygonContains(polygonA, polygonB[0][0]) {
		return true
	}

	// Check if polygonA is entirely within polygonB
	if planar.PolygonContains(polygonB, polygonA[0][0]) {
		return true
	}

	return false
}

// func (dre *DesignRuleEngine) Validate(
// }
