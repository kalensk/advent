package days


// write the simplest and smallest recursive function:

// 1) what does the algorithm do?
// 1A) See answers to 2-4 below.

// 2) what does my algorithm produce?
// 2A) Nothing. It returns nothing.

// 3) What is the base case? [when to stop recusing]
// 3A) when the input is false

// 4) What does it recurse on [the condition when to keep recursing (ie: call it self)]?
// 4A) The algorithm recurses when the input is true
func simplestRecursiveFunc(input bool) {
	if input {
		simplestRecursiveFunc(!input)
	}
}
// 1) what does the algorithm do?
// 1A)

// 2) what does my algorithm produce?
// 2A) Nothing. It returns nothing.

// 3) What is the base case? [when to stop recusing]
// 3A) when the input is zero it stops recursing

// 4) What does it recurse on [the condition when to keep recursing (ie: call it self)]?
// 4A) when the input is not-zero
func final(input int) {
	if input != 0 {
		final(input - 1)
	}
}
