package internal

import "strconv"

type ResultCd int

const (
	ResultCdSuccess         ResultCd = 1000
	ResultCdValidationError ResultCd = 2000
	ResultCdSystemError     ResultCd = 9000
)

func (r ResultCd) String() string {
	return strconv.Itoa(int(r))
}
