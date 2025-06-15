package size

import (
	"errors"
	"math/big"
	"regexp"
	"strings"
)

type unit int

type Size struct {
	Quantity *big.Float
	Unit     unit
}

const (
	B unit = iota
	KB
	MB
	GB
	TB
	PB
	EB
	KiB
	MiB
	GiB
	TiB
	PiB
	EiB
)

var (
	ErrInvalidSizeString error = errors.New("invalid size string")
)

var sizeRegexp = regexp.MustCompile(`(?i)^([0-9]+(?:\.[0-9]+)?(?:[eE][+-]?[0-9]+)?)([KMGTPE]i?B|B)$`)

func (u unit) DecimalFactor() *big.Float {
	switch u {
	case B:
		return big.NewFloat(1)
	case KB:
		return big.NewFloat(1e3)
	case MB:
		return big.NewFloat(1e6)
	case GB:
		return big.NewFloat(1e9)
	case TB:
		return big.NewFloat(1e12)
	case PB:
		return big.NewFloat(1e15)
	case EB:
		return big.NewFloat(1e18)
	case KiB:
		return new(big.Float).SetUint64(1 << 10)
	case MiB:
		return new(big.Float).SetUint64(1 << 20)
	case GiB:
		return new(big.Float).SetUint64(1 << 30)
	case TiB:
		return new(big.Float).SetUint64(1 << 40)
	case PiB:
		return new(big.Float).SetUint64(1 << 50)
	case EiB:
		return new(big.Float).SetUint64(1 << 60)
	default:
		return big.NewFloat(0)
	}
}

func (s Size) ToInt() *big.Int {
	factor := s.Unit.DecimalFactor()
	result := new(big.Float).Mul(s.Quantity, factor)
	i := new(big.Int)
	result.Int(i)
	return i
}

func ParseSizeFromString(size string) (*big.Int, error) {
	size = strings.TrimSpace(size)

	match := sizeRegexp.FindStringSubmatch(size)
	if match == nil {
		return new(big.Int), ErrInvalidSizeString
	}

	value, _, err := big.ParseFloat(match[1], 10, 256, big.ToZero)
	if err != nil {
		return new(big.Int), ErrInvalidSizeString
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
		return new(big.Int), ErrInvalidSizeString
	}

	return Size{Quantity: value, Unit: u}.ToInt(), nil
}
