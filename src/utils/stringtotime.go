package utils

import "time"


const layout string = "2006-01-02"

func StringToTime(str string) (time.Time, error) {
	t, err := time.Parse(layout, str)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}