package validate

import (
	"errors"
	"strconv"
)

func Int(value string, field string) (int, error) {
	i, err := strconv.Atoi(value)
	if err != nil {
		return 0, errors.New(field + " must be an integer")
	}
	return i, nil
}

func BlockNumber(value string) (int, error) {
	return Int(value, "block_number")
}

func Timestamp(value string) (int64, error) {
	t, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, errors.New("timestamp must be valid")
	}
	return t, nil
}

func TxCount(value string) (int, error) {
	tx, err := strconv.Atoi(value)
	if err != nil {
		return 0, errors.New("tx_count must be integer")
	}

	if tx < 0 {
		return 0, errors.New("tx_count cannot be negative")
	}

	return tx, nil
}
