package functional

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestAll(t *testing.T) {
	suite.Run(t, new(Suite))
}
