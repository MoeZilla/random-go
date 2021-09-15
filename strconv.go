/*
 * File: strconv.go
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
	"strconv"
)

// Atoi function allows you to get a random choice from a slice of type []string parsed into an integer with base 10.
// It returns the randomly chosen integer value from a slice containing elements (type string) and any write error encountered.
// example: random.Atoi([]string{"12", "28", "13"}), returns any one string from these three, parsed into an integer with base 10.
func Atoi(a []string) (int, error) {
	return strconv.Atoi(ChoiceString(a))
}

// Itoa function allows you to get a random choice from a slice of type []int parsed into a string.
// It returns the randomly chosen string from a slice containing elements (type int).
// example: random.Itoa([]int{383, 283, 282}), returns any one integer from these three, parsed into a string.
func Itoa(a []int) string {
	return strconv.Itoa(ChoiceInt(a))
}

// Quote function allows you to get a random choice from a slice of type []string as a double-quoted string.
// It returns the double-quoted randomly chosen string from a slice containing elements (type string).
// example: random.Quote([]string{"Hello", "Hi", "Nice"}), returns any one string from these three as a double-quoted string.
func Quote(a []string) string {
	return strconv.Quote(ChoiceString(a))
}

// Unquote function allows you to get a unquoted random choice from a slice of type []string containing double-quoted strings.
// It returns the unquoted randomly chosen string from a slice containing double-quoted elements (type string) and any write error encountered.
// example: random.Unquote([]string{`"Hello"`, `"Hi"`, `"Nice"`}), returns any one string from these three as an unquoted string.
func Unquote(a []string) (string, error) {
	const fn = "Unquote"
	s, err := strconv.Unquote(ChoiceString(a))
	if err != nil {
		err = &Error{fn, err}
	}
	return s, err
}
