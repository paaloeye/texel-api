/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package construction

import "github.com/paulmach/orb/geojson"

const (
	DesignRuleViolationPOfBound = iota
	DesignRuleViolationPConflict
	DesignRuleViolationPCoverage // Warning
)

type DesignRuleEngine struct {
}

type DesignRuleViolation struct {
}

func NewDesignRuleEngine() *DesignRuleEngine {
	return &DesignRuleEngine{}
}

func (dre *DesignRuleEngine) ValidateCollection(featureCollection *geojson.FeatureCollection) (ok bool, violations []error) {
	return true, nil
}

// func (dre *DesignRuleEngine) Validate(
// 	ctx context.Context,
// 	featureCollectionL *geojson.FeatureCollection,
// 	featureCollectionP *geojson.FeatureCollection) (warnCollection []DesignRuleViolation, errCollection []DesignRuleViolation, err error) {

// 	pSlice := []orb.Bound{}
// 	lSlice := []orb.Bound{}

// 	for _, f := range featureCollectionP.Features {
// 		pSlice = append(pSlice, f.Geometry.Bound())
// 	}

// 	for _, f := range featureCollectionL.Features {
// 		lSlice = append(lSlice, f.Geometry.Bound())
// 	}

// 	// P
// 	intersected := false
// 	for i := 0; i < len(pSlice); i++ {
// 		for j := i + 1; j < len(pSlice); j++ {
// 			intersected = pSlice[i].Intersects(pSlice[j])
// 			if intersected {
// 				break
// 			}
// 		}

// 		if intersected {
// 			break
// 		}
// 	}

// 	return nil, nil, nil
// }
