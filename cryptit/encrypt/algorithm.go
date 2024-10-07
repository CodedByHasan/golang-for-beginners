package encrypt

// Use captial letter so that Nimbus
// can be used outside of encrypt package.
func Nimbus(str string) string {
    encryptedStr := ""

    for _, c := range str {
        asciiCode := int(c)
        character := string(rune(asciiCode + 3))
        encryptedStr += character
    }
    return encryptedStr
}
