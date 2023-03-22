package cache

import "errors"

func validateKey(key string) error {
	if len(key) > 0 {
		return nil
	}

	return errors.New("key is empty")
}

func validateValue(value any) error {
	if value != nil {
		return nil
	}

	return errors.New("value is empty")
}
