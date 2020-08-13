package capitalize

import "strings"

// --- Directions
// Write a function that accepts a string.  The function should
// capitalize the first letter of each word in the string then
// return the capitalized string.
// --- Examples
//   capitalize('a short sentence') --> 'A Short Sentence'
//   capitalize('a lazy fox') --> 'A Lazy Fox'
//   capitalize('look, it is working!') --> 'Look, It Is Working!'

func capitalize(str string) string {
	arr := strings.Split(str, " ")
	var newArr []string
	for v := range arr {
		newArr = append(newArr, strings.ToUpper(arr[v][:1])+arr[v][1:])
	}
	return strings.Join(newArr, " ")
}
