// For every good kata idea there seem to be quite a few bad ones!

// In this kata you need to check the provided array (x) for good ideas 'good' and bad ideas 'bad'. If there are one or two good ideas, return 'Publish!', if there are more than 2 return 'I smell a series!'. If there are no good ideas, as is often the case, return 'Fail!'.
package kata

func Well(x []string) string {
	countGood := 0

	for _, v := range x {
		if v == "good" {
			countGood += 1
		}
	}

	if countGood > 2 {
		return "I smell a series!"
	}
	if countGood == 0 {
		return "Fail!"
	}
	return "Publish!"
}
