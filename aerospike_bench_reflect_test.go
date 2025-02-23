//go:build !as_performance

// Copyright 2014-2022 Aerospike, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package aerospike_test

import (
	"runtime"

	as "github.com/aerospike/aerospike-client-go/v7"

	"testing"
)

func benchGetObj(times int, client *as.Client, key *as.Key, obj interface{}) {
	for i := 0; i < times; i++ {
		if err = client.GetObject(nil, key, obj); err != nil {
			panic(err)
		}
	}
}

func benchPutObj(times int, client *as.Client, key *as.Key, wp *as.WritePolicy, obj interface{}) {
	for i := 0; i < times; i++ {
		if err = client.PutObject(wp, key, obj); err != nil {
			panic(err)
		}
	}
}

func Benchmark_GetObject(b *testing.B) {
	client, err := as.NewClientWithPolicy(clientPolicy, *host, *port)
	if err != nil {
		b.Fail()
	}

	key, _ := as.NewKey(*namespace, "databases", "Aerospike")

	obj := &OBJECT{198, "Jack Shaftoe and Company", []int64{1, 2, 3, 4, 5, 6}}
	client.PutObject(nil, key, obj)

	b.N = 1
	runtime.GC()
	b.ResetTimer()
	benchGetObj(b.N, client, key, obj)
}

func Benchmark_PutObject(b *testing.B) {
	client, err := as.NewClient(*host, *port)
	if err != nil {
		b.Fail()
	}

	// obj := &OBJECT{198, "Jack Shaftoe and Company", []byte(bytes.Repeat([]byte{32}, 1000))}
	obj := &OBJECT{198, "Jack Shaftoe and Company", []int64{1, 2, 3, 4, 5, 6}}
	key, _ := as.NewKey(*namespace, "databases", "Aerospike")
	writepolicy := as.NewWritePolicy(0, 0)

	b.N = 100
	runtime.GC()
	b.ResetTimer()
	benchPutObj(b.N, client, key, writepolicy, obj)
}
