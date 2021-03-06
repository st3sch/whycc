package converter

import (
	"strings"
	"time"
)

func convertDateFrom(layout string, value string) (string, error) {
	t, err := time.Parse(layout, value)
	if err != nil {
		return "", err
	}
	return t.Format("02/01/2006"), nil
}

func convertThousandAndCommaSeparator(value string) string {
	value = strings.Replace(value, ".", "", -1)
	value = strings.Replace(value, ",", ".", -1)
	return value
}

func abs(value string) string {
	return strings.TrimLeft(value, "-")
}

func isNegative(value string) bool {
	return value[0:1] == "-"
}

func ImplodeLines(s string) string {
	replacer := strings.NewReplacer("\r", "", "\n", " ")
	return strings.TrimSpace(replacer.Replace(s))
}
