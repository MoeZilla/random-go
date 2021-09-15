/*
 * File: examples.go
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
package main

import (
	"fmt"
	"strings"

	"github.com/anonyindian/random"
)

func main() {
	fmt.Println("Result of toss is", Toss())
	fmt.Println("I rolled a dice and the result was", RollDice())
	fmt.Println("Solve the jumbled sentence:", JumbleSentence("It's a sunny day"))
}

// This is an example function to explain the use of random.Bool()
func RandomBoolGeneration() string {
	b := random.Bool()
	if b {
		// will return the following value if b == true
		return "Yeah, you were right"
	}
	// will return the following value if b == false
	return "No, you were wrong"
}

// A function which returns any one value from heads and tails with the help random.Choice()
func Toss() string {
	r := random.Choice("heads", "tails")
	return fmt.Sprintf("%v", r)
}

// A function which return number from 1 to 6 like on a dice of 6 faces.
// random.Integer() was used for it.
func RollDice() int {
	i, err := random.Integer(1, 6)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return i
}

// A function to return a string with random person with a random profession in a string.
// String format: "{person} will become {"profession"}
// random.ChoiceString() was used for it.
func ChooseRandomPersonWithProfession() string {
	// Our data
	people := []string{"John", "Maureen", "Will", "Heroki", "Vijay", "Bill", "Penni", "Dawn", "Judi"}
	professions := []string{"doctor", "engineer", "programmer", "teacher", "nothing", "pilot", "scientist"}

	return random.ChoiceString(people) + " will become " + random.ChoiceString(professions)
}

// A function to let you know how does Shuffle functions work.
// We've used random.ShuffleStrings() function to jumble the sentence of type string.
// All the Shuffle functions are meant to be used in a similar way.
func JumbleSentence(sentence string) string {
	return strings.Join(random.ShuffleStrings(strings.Fields(sentence)), " ")
}
