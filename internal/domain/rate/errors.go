package rate

import "errors"

var (
	ErrBaseNotSupported  = errors.New("base currency not supported")
	ErrThirdPartyRequest = errors.New("rate: can't reach third party service")
)
