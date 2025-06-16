package size

import (
	"fmt"
	"math/big"
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

func (u unit) String() string {
	switch u {
	case B:
		return "B"
	case KB:
		return "KB"
	case MB:
		return "MB"
	case GB:
		return "GB"
	case TB:
		return "TB"
	case PB:
		return "PB"
	case EB:
		return "EB"
	case KiB:
		return "KiB"
	case MiB:
		return "MiB"
	case GiB:
		return "GiB"
	case TiB:
		return "TiB"
	case PiB:
		return "PiB"
	case EiB:
		return "EiB"
	default:
		return ""
	}
}

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
		return nil
	}
}

func (s Size) Int() *big.Int {
	factor := s.Unit.DecimalFactor()
	if factor == nil {
		return nil
	}
	result := new(big.Float).Mul(s.Quantity, factor)
	i := new(big.Int)
	result.Int(i)
	return i
}

func (s Size) String() string {
	bytes := new(big.Float).Mul(s.Quantity, s.Unit.DecimalFactor())

	units := []struct{ unit unit }{{B}, {KB}, {MB}, {GB}, {TB}, {PB}, {EB}}

	for _, u := range units {
		factor := u.unit.DecimalFactor()
		normalized := new(big.Float).Quo(bytes, factor)

		if f, _ := normalized.Float64(); f < 1000 {
			return fmt.Sprintf("%.3g%s", normalized, u.unit.String())
		}
	}

	return fmt.Sprintf("%sB", bytes.Text('f', 0))
}
