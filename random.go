/*
 * File: random.go
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
	"time"
)

// Bool function is used to return a random bool value from true and false (type bool).
// example: random.Bool(), returns any one from true, false randomly.
func Bool() bool {
	a := []bool{false, true}
	return a[rand.Intn(len(a))]
}

// Integer function is used to get a random integer between a range [startNum, endNum].
// startNum (type int) is an integer from where a random value will be chosen.
// endNum (type int) is an integer upto which, a random value will be chosen.
// It returns the randomly chosen value of type int and and any write error encountered.
// example: random.Integer(10, 20), returns any one integer from the range [10, 20].
func Integer(startNum int, endNum int) (int, error) {
	const fn = "Integer"
	if (startNum > endNum) || (startNum == endNum) {
		return 0, &Error{fn, ErrEndNumSmaller}
	}
	intslice := make([]int, 0)
	for i := startNum; i <= endNum; i++ {
		intslice = append(intslice, i)
	}
	return ChoiceInt(intslice), nil
}

// Float32 function is used to get a random float32 value between a range [startNum, endNum].
// startNum (type float32) is a float value from where a random value will be chosen.
// endNum (type float32) is a float value upto which, a random value will be chosen.
// It returns the randomly chosen value of type float32 and and any write error encountered.
// example: random.Float32(1.292, 1.388), returns any one float32 value from the range [1.292, 1.388].
func Float32(startNum float32, endNum float32) (float32, error) {
	const fn = "Float32"
	if (startNum > endNum) || (startNum == endNum) {
		return 0, &Error{fn, ErrEndNumSmaller}
	}
	num, err := Float32N(startNum, endNum, 1)
	return num[0], err
}

// Float64 function is used to get a random float64 value between a range [startNum, endNum].
// startNum (type float64) is a float value from where a random value will be chosen.
// endNum (type float64) is a float value upto which, a random value will be chosen.
// It returns the randomly chosen value of type float64 and and any write error encountered.
// example: random.Float64(1.292, 1.388), returns any one float64 value from the range [1.292, 1.388].
func Float64(startNum float64, endNum float64) (float64, error) {
	const fn = "Float64"
	if (startNum > endNum) || (startNum == endNum) {
		return 0, &Error{fn, ErrEndNumSmaller}
	}
	num, err := Float64N(startNum, endNum, 1)
	return num[0], err
}

// IntegerN function is used to get an array of type []int containing integers between a range [startNum, endNum].
// startNum (type int) is an integer from where a random value will be chosen.
// endNum (type int) is an integer upto which, a random value will be chosen.
// n (type int) is the number of values to be randomly chosen.
// It returns the randomly chosen values of type int in an array of type []int and any write error encountered.
// example: random.IntegerN(10, 20, 2), returns an array containing 2 integers from the range [10, 20].
func IntegerN(startNum int, endNum int, n int) ([]int, error) {
	const fn = "IntegerN"
	if (startNum > endNum) || (startNum == endNum) {
		return nil, &Error{fn, ErrEndNumSmaller}
	}
	r := make([]int, n)
	for i := range r {
		in, _ := Integer(startNum, endNum)
		r[i] = in
	}
	return r, nil
}

// Float32N function is used to get an array of type []float32 containing float32 values between a range [startNum, endNum].
// startNum (type float32) is an float32 value from where a random value will be chosen.
// endNum (type float32) is an float32 value upto which, a random value will be chosen.
// n (type int) is the number of values to be randomly chosen.
// It returns the randomly chosen values of type float32 in an array of type []float32 and any write error encountered.
// example: random.Float32N(1.11, 2.22, 2), returns an array containing 2 float32 values from the range [1.11, 2.22].
func Float32N(startNum float32, endNum float32, n int) ([]float32, error) {
	const fn = "Float32N"
	if (startNum > endNum) || (startNum == endNum) {
		return nil, &Error{fn, ErrEndNumSmaller}
	}
	r := make([]float32, n)
	for i := range r {
		r[i] = startNum + rand.Float32()*(endNum-startNum)
	}
	return r, nil
}

// Float64N function is used to get an array of type []float64 containing float64 values between a range [startNum, endNum].
// startNum (type float64) is an float64 value from where a random value will be chosen.
// endNum (type float64) is an float64 value upto which, a random value will be chosen.
// n (type int) is the number of values to be randomly chosen.
// It returns the randomly chosen values of type float64 in an array of type []float64 and any write error encountered.
// example: random.Float64N(1.11, 2.22, 2), returns an array containing 2 float64 values from the range [1.11, 2.22].
func Float64N(startNum float64, endNum float64, n int) ([]float64, error) {
	const fn = "Float64N"
	if (startNum > endNum) || (startNum == endNum) {
		return nil, &Error{fn, ErrEndNumSmaller}
	}
	r := make([]float64, n)
	for i := range r {
		r[i] = startNum + rand.Float64()*(endNum-startNum)
	}
	return r, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
