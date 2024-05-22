package refract

import "fmt"

type Assertion struct{}

func (a *Assertion) Assertint(v any) (int, error) {
	out, ok := v.(int)
	if ok {
		return out, nil
	}
	return 0, fmt.Errorf("couldn't assert as int")
}

func (a *Assertion) Assertstring(v any) (string, error) {
	out, ok := v.(string)
	if ok {
		return out, nil
	}
	return "", fmt.Errorf("couldn't assert as int")
}
