/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package kie

import (
	"errors"

	"github.com/arielsrv/go-archaius/source/remote"
)

// GenerateLabels generate labels to an specific dimensions.
func GenerateLabels(dimension DimensionName, optionsLabels map[string]string) (map[string]string, error) {
	if optionsLabels == nil {
		return nil, remote.ErrLabelsNil
	}
	dimensionLabels := make(map[string]string)
	if optionsLabels[remote.LabelApp] == "" {
		return nil, remote.ErrAppEmpty
	}
	dimensionLabels[remote.LabelApp] = optionsLabels[remote.LabelApp]
	dimensionLabels[remote.LabelEnvironment] = optionsLabels[remote.LabelEnvironment]
	if dimension == DimensionApp {
		return dimensionLabels, nil
	}
	dimensionLabels[remote.LabelService] = optionsLabels[remote.LabelService]
	if dimension == DimensionService {
		return dimensionLabels, nil
	}
	return nil, errors.New("do not support dimension " + string(dimension))
}
