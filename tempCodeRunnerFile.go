func int2DArrayToString(input [][]int) string {

    var res string
    
	for i := 0; i < len(input); i++ {

		row := input[i]
		var rowStrng string
		
        for j := 0; j < len(row); j++ {
			rowStrng += fmt.Sprintf(" %6d", row[j])
		}
	
        res += rowStrng + "\n"
	
    }

	return res
}