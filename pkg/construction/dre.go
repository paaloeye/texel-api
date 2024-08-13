/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package construction

import (
	"context"

	"github.com/paulmach/orb/geojson"
)

type DesignRuleEngine struct {
}

type DesignRuleViolation struct {
}

func NewDesignRuleEngine() *DesignRuleEngine {
	return &DesignRuleEngine{}
}

func (dre *DesignRuleEngine) Validate(
	ctx context.Context,
	featureCollectionA *geojson.FeatureCollection,
	featureCollectionB *geojson.FeatureCollection) (warnCollection []DesignRuleViolation, errCollection []DesignRuleViolation, err error) {

	return nil, nil, nil
}
