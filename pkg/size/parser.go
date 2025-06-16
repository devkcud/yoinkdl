package size

import (
	"errors"
	"math/big"
	"strings"
)

var ErrInvalidSizeString error = errors.New("invalid size string")

func ParseSizeFromString(size string) (*Size, error) {
	size = strings.TrimSpace(size)

	match := sizeRegex.FindStringSubmatch(size)
	if match == nil {
		return nil, ErrInvalidSizeString
	}

	value, _, err := big.ParseFloat(match[1], 10, 256, big.ToZero)
	if err != nil {
		return nil, ErrInvalidSizeString
	}

	unitStr := strings.ToLower(match[2])
	var u unit

	switch unitStr {
	case "b":
		u = B
	case "kb":
		u = KB
	case "mb":
		u = MB
	case "gb":
		u = GB
	case "tb":
		u = TB
	case "pb":
		u = PB
	case "eb":
		u = EB
	case "kib":
		u = KiB
	case "mib":
		u = MiB
	case "gib":
		u = GiB
	case "tib":
		u = TiB
	case "pib":
		u = PiB
	case "eib":
		u = EiB
	default:
		return nil, ErrInvalidSizeString
	}

	return &Size{Quantity: value, Unit: u}, nil
}
