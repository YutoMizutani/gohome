package primitive

import "strconv"

type GormModelID struct {
	Value uint
}

func NewGormModelID(str string) *GormModelID {
	id64, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return nil
	}
	id := uint(id64)

	p := new(GormModelID)
	p.Value = id
	return p
}
