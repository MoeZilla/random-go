/*
 * File: choice.go
 * Created on Wed Sep 15 2021
 *
 * The MIT License (MIT)
 * Copyright (c) 2021 Veer (anonyindian)
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software
 * and associated documentation files (the "Software"), to deal in the Software without restriction,
 * including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense,
 * and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so,
 * subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all copies or substantial
 * portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED
 * TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL
 * THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
 * TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */
package random

import (
	"math/rand"
)

// Choice indexes its parameters and pick a random choice from them.
// it can be of any type: string, integer, floats, slices, bool, etc.
// It returns the randomly chosen value of type interface{}.
// example: random.Choice('a', 1, true), returns any one value from 'a' (type string), 1 (type int) and true (type bool) randomly.
func Choice(a ...interface{}) interface{} {
	return a[rand.Intn(len(a))]
}

// ChoiceString indexes the string slice and pick a random choice from it.
// Its parameter 'a' must be of type []string.
// It returns the randomly chosen value of type string.
// example: random.ChoiceString('a', "b", "c"), returns any one string from 'a', "b" and "c" randomly.
func ChoiceString(a []string) string {
	return a[rand.Intn(len(a))]
}

// ChoiceInt indexes the int slice and pick a random choice from it.
// Its parameter 'a' must be of type []int.
// It returns the randomly chosen value of type int.
// example: random.ChoiceInt(1, 2, 3), returns any one integer from 1, 2 and 3 randomly.
func ChoiceInt(a []int) int {
	return a[rand.Intn(len(a))]
}

// ChoiceInt8 indexes the int8 slice and pick a random choice from it.
// Its parameter 'a' must be of type []int8.
// It returns the randomly chosen value of type int8.
func ChoiceInt8(a []int8) int8 {
	return a[rand.Intn(len(a))]
}

// ChoiceInt16 indexes the int16 slice and pick a random choice from it.
// Its parameter 'a' must be of type []int16.
// It returns the randomly chosen value of type int16.
func ChoiceInt16(a []int16) int16 {
	return a[rand.Intn(len(a))]
}

// ChoiceInt32 indexes the int32 slice and pick a random choice from it.
// Its parameter 'a' must be of type []int32.
// It returns the randomly chosen value of type int32.
func ChoiceInt32(a []int32) int32 {
	return a[rand.Intn(len(a))]
}

// ChoiceInt64 indexes the int64 slice and pick a random choice from it.
// Its parameter 'a' must be of type []int64.
// It returns the randomly chosen value of type int64.
func ChoiceInt64(a []int64) int64 {
	return a[rand.Intn(len(a))]
}

// ChoiceUint indexes the uint slice and pick a random choice from it.
// Its parameter 'a' must be of type []uint.
// It returns the randomly chosen value of type uint.
func ChoiceUint(a []uint) uint {
	return a[rand.Intn(len(a))]
}

// ChoiceUint8 indexes the uint8 slice and pick a random choice from it.
// Its parameter 'a' must be of type []uint8.
// It returns the randomly chosen value of type uint8.
func ChoiceUint8(a []uint8) uint8 {
	return a[rand.Intn(len(a))]
}

// ChoiceUint16 indexes the uint16 slice and pick a random choice from it.
// Its parameter 'a' must be of type []uint16.
// It returns the randomly chosen value of type uint16.
func ChoiceUint16(a []uint16) uint16 {
	return a[rand.Intn(len(a))]
}

// ChoiceUint indexes the uint32 slice and pick a random choice from it.
// Its parameter 'a' must be of type []uint32.
// It returns the randomly chosen value of type uint32.
func ChoiceUint32(a []uint32) uint32 {
	return a[rand.Intn(len(a))]
}

// ChoiceUint64 indexes the uint64 slice and pick a random choice from it.
// Its parameter 'a' must be of type []uint64.
// It returns the randomly chosen value of type uint64.
func ChoiceUint64(a []uint64) uint64 {
	return a[rand.Intn(len(a))]
}

// ChoiceFloat32 indexes the float32 slice and pick a random choice from it.
// Its parameter 'a' must be of type []float32.
// It returns the randomly chosen value of type float32.
// example: random.ChoiceFloat32(1.234, 2.38333, 5.3227), returns any one float value from 1.234, 2.38333, 5.3227 randomly.
func ChoiceFloat32(a []float32) float32 {
	return a[rand.Intn(len(a))]
}

// ChoiceFloat64 indexes the float32 slice and pick a random choice from it.
// Its parameter 'a' must be of type []float64.
// It returns the randomly chosen value of type float64.
func ChoiceFloat64(a []float64) float64 {
	return a[rand.Intn(len(a))]
}
