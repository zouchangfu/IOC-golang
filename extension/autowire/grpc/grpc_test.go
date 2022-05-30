/*
 * Copyright (c) 2022, Alibaba Group;
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package grpc

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/alibaba/ioc-golang/autowire"
)

type mockInterface interface {
}

type mockImpl struct {
}

const mockInterfaceName = "github.com/alibaba/ioc-golang/extension/autowire/grpc.mockInterface"
const mockImplName = "github.com/alibaba/ioc-golang/extension/autowire/grpc.mockImpl"

func TestAutowire_RegisterAndGetAllStructDescriptors(t *testing.T) {
	t.Run("test config autowire register and get all struct descriptors", func(t *testing.T) {
		sd := &autowire.StructDescriptor{
			Interface: new(mockInterface),
			Factory: func() interface{} {
				return &mockImpl{}
			},
		}
		RegisterStructDescriptor(sd)
		a := &Autowire{}
		allStructDesc := a.GetAllStructDescriptors()
		assert.NotNil(t, allStructDesc)
		sdid := mockImplName + "#" + mockInterfaceName
		structDesc, ok := allStructDesc[sdid]
		assert.True(t, ok)
		assert.Equal(t, sdid, structDesc.ID())
	})
}

func TestAutowire_TagKey(t *testing.T) {
	t.Run("test grpc autowire tag", func(t *testing.T) {
		a := &Autowire{}
		assert.Equal(t, Name, a.TagKey())
	})
}
