/*
Copyright 2019 Baidu, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package clustermessage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUtil(t *testing.T) {
	msg := &ClusterMessage{}
	// serialize
	data, err := msg.Serialize()
	assert.NotNil(t, data)
	assert.Nil(t, err)

	// nil deserialize
	err = msg.Deserialize(nil)
	assert.NotNil(t, err)
	// not nil deserialize
	err = msg.Deserialize([]byte{})
	assert.Nil(t, err)
	err = msg.Deserialize([]byte{'{'})
	assert.NotNil(t, err)
}
