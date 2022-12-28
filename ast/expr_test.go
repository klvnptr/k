package ast_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ExprTestSuite struct {
	suite.Suite
}

func (suite *ExprTestSuite) TestRange() {
	_, err := strconv.ParseInt("100000000000000000000", 10, 32)
	suite.ErrorIs(err, strconv.ErrRange)

	_, err = strconv.ParseInt("100000000000000000000100000000000000000000100000000000000000000", 10, 64)
	suite.ErrorIs(err, strconv.ErrRange)
}

func TestExprTestSuite(t *testing.T) {
	suite.Run(t, new(ExprTestSuite))
}
