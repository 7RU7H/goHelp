
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
