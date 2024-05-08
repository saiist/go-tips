package src

import "errors"

func join(a, b string, max int) (string, error) {
	if a == "" {
		return "", errors.New("a is empty")
	}

	if b == "" {
		return "", errors.New("b is empty")
	}

	concat, err := concatenate(a, b)
	if err != nil {
		return "", err
	}

	if len(concat) > max {
		return concat[:max], nil
	}

	return concat, nil
}

func concatenate(a, b string) (string, error) {
	return a + b, nil
}
