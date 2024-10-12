//
/*
 *
 * Copyright 2024 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package grpcrand implements math/rand functions in a concurrent-safe way
// with a global random source, independent of math/rand's global source.
package grpcrand

import "math/rand"

// This implementation will be used for Go version 1.21 or newer.
// For older versions, the original implementation with mutex will be used.

// Int implements rand.Int on the grpcrand global source.
func Int() int {
	return rand.Int()
}

// Int63n implements rand.Int63n on the grpcrand global source.
func Int63n(n int64) int64 {
	return rand.Int63n(n)
}

// Intn implements rand.Intn on the grpcrand global source.
func Intn(n int) int {
	return rand.Intn(n)
}

// Int31n implements rand.Int31n on the grpcrand global source.
func Int31n(n int32) int32 {
	return rand.Int31n(n)
}

// Float64 implements rand.Float64 on the grpcrand global source.
func Float64() float64 {
	return rand.Float64()
}

// Uint64 implements rand.Uint64 on the grpcrand global source.
func Uint64() uint64 {
	return rand.Uint64()
}

// Uint32 implements rand.Uint32 on the grpcrand global source.
func Uint32() uint32 {
	return rand.Uint32()
}

// ExpFloat64 implements rand.ExpFloat64 on the grpcrand global source.
func ExpFloat64() float64 {
	return rand.ExpFloat64()
}

// Shuffle implements rand.Shuffle on the grpcrand global source.
var Shuffle = func(n int, f func(int, int)) {
	rand.Shuffle(n, f)
}
