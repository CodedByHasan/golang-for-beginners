package encrypt

func Nimbus(str string) string {
    encryptedStr := ""

    for _, c := range str {
        asciiCode := int(c)
        character := string(rune(asciiCode + 3))
        encryptedStr += character
    }
    return encryptedStr
}
