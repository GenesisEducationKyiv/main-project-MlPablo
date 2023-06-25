package rate_domain

import "errors"

var ErrThirdPartyRequest = errors.New("rate: can't reach third party service")
