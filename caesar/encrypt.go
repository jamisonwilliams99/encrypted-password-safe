package caesar

import (
	"strconv"
)

type keySegment struct {
	shiftDirection bool
	shiftAmount    int
}

func detShiftDir(shiftDirBin string, encryption bool) bool {
	var shiftDir bool
	if encryption {
		if shiftDirBin[len(shiftDirBin)-1:] == "0" {
			shiftDir = false
		} else {
			shiftDir = true
		}
	} else {
		if shiftDirBin[len(shiftDirBin)-1:] == "0" {
			shiftDir = true
		} else {
			shiftDir = false
		}
	}
	return shiftDir
}

// encryption = true  -> we are encrypting
// encryption = false -> we are decrypting
func decodeKey(key string, encryption bool) []keySegment {
	var keySegments []keySegment
	keyRunes := []rune(key) // convert key string to slice of runes
	for i := 0; i < len(key); i = i + 2 {
		shiftDirByte := keyRunes[i]
		shiftAmtByte := keyRunes[i+1]

		shiftDirInt := int(shiftDirByte)
		shiftAmtInt := int(shiftAmtByte)

		shiftDirBin := strconv.FormatInt(int64(shiftDirInt), 2) // binary string representation of the key segments shift direction number

		shiftDir := detShiftDir(shiftDirBin, encryption)

		ks := keySegment{
			shiftDirection: shiftDir,
			shiftAmount:    shiftAmtInt,
		}
		keySegments = append(keySegments, ks)
	}
	return keySegments
}

func shiftLetter(c rune, shiftAmount int, shiftDir bool) string {
	asciiNum := int(c)

	for i := 0; i < shiftAmount; i++ {
		if shiftDir {
			if asciiNum < 127 {
				asciiNum++
			} else {
				asciiNum = 0
			}
		} else {
			if asciiNum > 0 {
				asciiNum--
			} else {
				asciiNum = 127
			}
		}
	}

	return string(rune(asciiNum))
}

func shiftText(text string, keySegments []keySegment) string {
	shiftedText := ""
	j := 0
	counter := 0
	for _, c := range text {
		ks := keySegments[j]
		shiftedText += shiftLetter(c, ks.shiftAmount, ks.shiftDirection)
		counter++
		if counter == 4 {
			j++
			counter = 0
		}
	}
	return shiftedText
}

func Encrypt(text string, key string) string {
	keySegments := decodeKey(key, true)
	encryptedText := shiftText(text, keySegments)
	return encryptedText
}

func Decrypt(text string, key string) string {
	keySegments := decodeKey(key, false)
	decryptedText := shiftText(text, keySegments)
	return decryptedText
}
