func padding(str string, padding_char string) string {

	var strlen int

	var buf bytes.Buffer

	strlen = len(str)

	if strlen%16 == 0 {
		return str
	} else if strlen < 16 {

		buf.WriteString(str)

		strlen = 16 - strlen

		for i := 0; i < strlen; i++ {
			buf.WriteString(padding_char)
		}

		return buf.String()
	} else {
		//strlen > 16
		buf.WriteString(str)

		buf.WriteString(padding_char)

		padding(buf.String(), padding_char)
	}

	return str
}
