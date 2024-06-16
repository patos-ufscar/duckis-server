package common_test

import (
	"testing"

	"github.com/patos-ufscar/duckis-server/common"
)

func TestConvertWildCardsToPattern(t *testing.T) {
	s := "name:*:info"
	ret := common.ConvertWildCardsToPattern(s)
	if ret != `name:\*:info` {
		t.Errorf("common.ConvertWildCardsToPattern(%s) should match %s but is: %s", s, `name:\*:info`, ret)
	}
}
