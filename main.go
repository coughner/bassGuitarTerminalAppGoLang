package main

import (
	"fmt"
	"math/rand"
	"time"
)

// this function gets a random int from a range of ints
func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

// this function applies the notes to the string
func newString(x [12]string, y [12]string, z int) [12]string {
	fretInt := 0
	noteInt := z
	for i := 0; i < 12; i++ {
		if noteInt == 12 {
			noteInt = 0
		}
		x[fretInt] = y[noteInt]
		fretInt++
		noteInt++
	}
	return x
}

// this function picks a random bass string from the four
func pickRandStr(e [12]string, a [12]string, d [12]string, g [12]string) [12]string {
	x := randInt(1, 4)
	var pickedStr [12]string
	if x == 1 {
		pickedStr = e
	} else if x == 2 {
		pickedStr = a
	} else if x == 3 {
		pickedStr = d
	} else if x == 4 {
		pickedStr = g
	} else {
		fmt.Println("the number did not equal one thru four")
	}
	return pickedStr
}

// this function picks a random fret
func pickRandFret() int {
	x := randInt(0, 11)
	return x
}

func addPoint(x int) int {
	x++
	return x
}

func losePoint(x int) int {
	x--
	return x
}

func main() {

	Score := 0

	// creates an array that stores all the notes in music :)
	var scale [12]string
	scale[0] = "c"
	scale[1] = "c#"
	scale[2] = "d"
	scale[3] = "d#"
	scale[4] = "e"
	scale[5] = "f"
	scale[6] = "f#"
	scale[7] = "g"
	scale[8] = "g#"
	scale[9] = "a"
	scale[10] = "a#"
	scale[11] = "b"

	// creating an empty array of each of the strings on the bass
	var eString [12]string
	var aString [12]string
	var dString [12]string
	var gString [12]string

	// creates a new array and applies the value
	var newEString [12]string = newString(eString, scale, 4)
	var newAString [12]string = newString(aString, scale, 9)
	var newDString [12]string = newString(dString, scale, 2)
	var newGString [12]string = newString(gString, scale, 7)

	// tests all the strings
	fmt.Println(newEString[0:11])
	fmt.Println(newAString[0:11])
	fmt.Println(newDString[0:11])
	fmt.Println(newGString[0:11])

	for {
		// i believe this sets the seed from the time
		rand.Seed(time.Now().UnixNano())

		// uses the pick string function to pick a string
		pickedString := pickRandStr(newEString, newAString, newDString, newGString)
		// picks a random fret number between 1 and 12
		fretNum := pickRandFret()

		// uses picked string and fret number to get a picked note
		var pickedNote string = pickedString[fretNum]

		// tell the user the picked string and fret
		fmt.Printf("What is %s%d", pickedString[0], fretNum)

		// get input from user
		var guess string

		fmt.Scanln(&guess)

		// check it with stored note
		if guess == pickedNote {
			fmt.Println("Correct Answer!!!")
			Score = addPoint(Score)
		} else {
			fmt.Println("Wrong Answer")
			Score = losePoint(Score)
		}
		fmt.Println(Score)
	}
}

/* ideas for future:
add a timer that you have to get it within
add a gui
add points
add lvls of difficulty
move it to mobile
*/
