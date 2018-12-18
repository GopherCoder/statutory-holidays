package holiday

import (
	"fmt"
	"testing"
)

func TestValidYear(tests *testing.T) {
	tt := []struct {
		Field GetHolidaysParams
	}{
		{
			Field: GetHolidaysParams{
				Year:   "2020",
				ChName: "中秋靠",
				EnName: "Na",
			},
		},
		{
			Field: GetHolidaysParams{
				Year:   "2018 ",
				ChName: "中",
				EnName: "Mid",
			},
		},
	}

	for _, t := range tt {
		fmt.Println(t.Field.validYear())
		fmt.Println(t.Field.validChName())
		fmt.Println(t.Field.validEnName())
	}
}
