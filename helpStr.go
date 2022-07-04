
import (
	"fmt"
	"strings"
)

//collect argument is for some string pattern not regex to marshal strings in the list  
func replaceWS20(scrap, collect string, size int) string {
	rLargeWS := strings.NewReplacer(" ", ",", "   ", ",", "    ", ",", "     ", ",", "      ", ",", "       ", ",", "        ", ",", "         ", ",", "          ", ",", "           ", ",", "            ", ",", "             ", ",", "              ", ",", "               ", ",", "                ", ",", "                 ", ",", "                  ", ",", "                   ", ",", "                    ", ",")
	rSingleWS := strings.NewReplacer(" ", "")
	scrapComma := rLargeWS.Replace(scrap)
	scrapSlice := strings.Split(scrapComma, ",")
	data := make([]string, size)
	i := 0
	for _, elem := range scrapSlice {
		if strings.Contains(elem, collect) {
			if strings.Contains(elem, " ") {
				elem = rSingleWS.Replace(elem)
			}
			data[i] = elem
			i++
		}
	}
  result := strings.Join(data, "")
	return result
}

func reverseStrEndianness(s string) string {
	builder := strings.Builder{}
	var editStr string
	if strings.Contains(s, "0x") {
		editStr = strings.TrimPrefix(s, "0x")
	} else {
		editStr = s
	}
	sizeS := len(editStr) - 1
	pieces, piecesSize := explodeToPairs(editStr, sizeS)
	for i := piecesSize - 1; i >= 0; i-- {
		builder.WriteString(pieces[i])
	}
	return builder.String()
}


func covertStrToByte(s string) []byte {
  result := byte[](s)
	return result
}

func add0xStub(s string) string {
	addzeroX := "0x" + s
	return addzeroX
}

func remove0xStub(s string) string {
	stubless := strings.TrimPrefix(s, "0x")
	return stubless
}
