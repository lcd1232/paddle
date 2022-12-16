package paddle

import "time"

type Time time.Time

func (t *Time) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	t1, err := time.Parse(`"2006-01-02 15:04:05.000000"`, string(data))
	if err != nil {
		return err
	}

	*t = Time(t1)
	return nil
}
