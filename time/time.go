package time

import (
	"fmt"
	"strconv"
	//"time"
)

// DehumanizeDuration converts string in format 1w1d1h1m into number of seconds
func DehumanizeDuration(dura string) (int64, error) {
	var num string
	var base int64
	var duration int64

	for i := 0; i < len(dura); i++ {
		// Check if current char is a number or string
		if _, err := strconv.Atoi(string(dura[i])); err == nil {
			num += string(dura[i])
		} else {
			// Convert existing num to number as we got our first char
			if d, err := strconv.Atoi(num); err == nil {
				base = int64(d)
				num = ""
			}

			switch string(dura[i]) {
			case "w":
				duration += (base * 604800)
			case "d":
				duration += (base * 86400)
			case "h":
				duration += (base * 3600)
			case "m":
				duration += (base * 60)
			default:
				return 0, fmt.Errorf("wrong char only w,d,h,m are accepted: %s", string(dura[i]))
			}

			base = 0
		}
	}

	return duration, nil
}
