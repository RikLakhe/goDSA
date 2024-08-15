// find shortest string
func shortestString(wordOne, wordTwo string) (string, string) {
	if len(wordOne) < len(wordTwo) {
		return wordOne, wordTwo
	} else if wordTwo < wordOne {
		return wordTwo, wordOne
	} else {
		return wordOne, wordTwo
	}
}

// Function to convert a string to an array of strings (in this example, placeholder)
func stringToArray(word string) []string {
	result := strings.Split(word, "")
	return result // Placeholder implementation
}

// Function to merge two arrays into a single string (in this example, placeholder)
func mergeArraysToSingleString(firstArray, secondArray []string) string {
    var mergedArray []string
    var firstLength = len(firstArray)
    var secondLength = len(secondArray)
    var total = firstLength + secondLength
    for count:=0 ; count<=total; count++ {
        if count+1 <= firstLength {
        mergedArray = append(mergedArray, firstArray[count])
        }
        if count+1 <= secondLength {
        mergedArray = append(mergedArray, secondArray[count])
        }
    }
	return strings.Join(mergedArray, "")
}

func mergeAlternately(word1 string, word2 string) string {
	/*
	   we get 2 words and out come should be a word which is the combination of these two words
	   1- split to array ?
	   2- length of word take the shortest first and then longest
	       loop in words and add single string in a var

	   if shortest is done move all remaining from longest to var
	*/
	//declare return array
	// var returnArray []string

	// shortest, longest := shortestString(word1, word2)
	fistArray := stringToArray(word1)
	secondArray := stringToArray(word2)

	
	returnString := mergeArraysToSingleString(fistArray, secondArray)
	// merge the return array
	return returnString
}