/**
 * Copyright 2023 IBM Corp.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package driver

import (
	"fmt"
	"github.com/container-storage-interface/spec/lib/go/csi"
)

func ReplaceAndReturnCopy(req interface{}, newAccessKey, newSecretKey string) (interface{}, error) {
	switch r := req.(type) {
	case *csi.CreateVolumeRequest:
		// Create a new CreateVolumeRequest and copy the original values
		newReq := &csi.CreateVolumeRequest{}
		*newReq = *r

		// Modify the Secrets map in the new request
		newReq.Secrets = map[string]string{
			"accesKey":  newAccessKey,
			"secretKey": newSecretKey,
		}

		return newReq, nil
	case *csi.DeleteVolumeRequest:
		// Create a new DeleteVolumeRequest and copy the original values
		newReq := &csi.DeleteVolumeRequest{}
		*newReq = *r

		// Modify the Secrets map in the new request
		newReq.Secrets = map[string]string{
			"accesKey":  newAccessKey,
			"secretKey": newSecretKey,
		}

		return newReq, nil
	case *csi.NodePublishVolumeRequest:
		// Create a new DeleteVolumeRequest and copy the original values
		newReq := &csi.NodePublishVolumeRequest{}
		*newReq = *r

		// Modify the Secrets map in the new request
		newReq.Secrets = map[string]string{
			"accesKey":  newAccessKey,
			"secretKey": newSecretKey,
		}

		return newReq, nil

	default:
		return req, fmt.Errorf("unsupported request type")
	}
}
