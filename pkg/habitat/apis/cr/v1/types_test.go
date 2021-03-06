// Copyright (c) 2017 Chef Software Inc. and/or applicable contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1

import (
	"math/rand"
	"testing"

	"github.com/google/gofuzz"

	apitesting "k8s.io/apimachinery/pkg/api/testing"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
)

var _ runtime.Object = &ServiceGroup{}
var _ metav1.ObjectMetaAccessor = &ServiceGroup{}

var _ runtime.Object = &ServiceGroupList{}
var _ metav1.ListMetaAccessor = &ServiceGroupList{}

func serviceGroupFuzzerFuncs(t apitesting.TestingCommon) []interface{} {
	return []interface{}{
		func(obj *ServiceGroupList, c fuzz.Continue) {
			c.FuzzNoCustom(obj)
			obj.Items = make([]ServiceGroup, c.Intn(10))
			for i := range obj.Items {
				c.Fuzz(&obj.Items[i])
			}
		},
	}
}

// TestRoundTrip tests that the third-party kinds can be marshaled and unmarshaled correctly to/from JSON
// without the loss of information. Moreover, deep copy is tested.
func TestRoundTrip(t *testing.T) {
	scheme := runtime.NewScheme()
	codecs := serializer.NewCodecFactory(scheme)

	AddToScheme(scheme)

	seed := rand.Int63()
	fuzzerFuncs := apitesting.MergeFuzzerFuncs(t, apitesting.GenericFuzzerFuncs(t, codecs), serviceGroupFuzzerFuncs(t))
	fuzzer := apitesting.FuzzerFor(fuzzerFuncs, rand.NewSource(seed))

	apitesting.RoundTripSpecificKindWithoutProtobuf(t, SchemeGroupVersion.WithKind("ServiceGroup"), scheme, codecs, fuzzer, nil)
	apitesting.RoundTripSpecificKindWithoutProtobuf(t, SchemeGroupVersion.WithKind("ServiceGroupList"), scheme, codecs, fuzzer, nil)
}
