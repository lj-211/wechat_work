package internal

import (
	"errors"
	"testing"

	"github.com/lj-211/wechat_work/internal/util"
	"github.com/stretchr/testify/assert"
)

var fakeUtripCode = "fake-utrip-code"
var fakeUtripKey = "fake-utrip-key"

var fakeInvalidUtripCode = ""
var fakeInvalidUtripKey = ""

func MockConfigGetEnvironment() string {
	return "test"
}

func TestPackUtripUrl(t *testing.T) {
	//t.Parallel()

	expect := "http://ubtrip.eatuo.com:9081/#/singleLogin?&user=apitest@ssharing.com&usertype=3&name=测试&corpcode=TestCorp&sign=a6d1aec46f958988cd3e8dfd4142d92a&type=home"

	fakeCorpId := "TestCorp"
	fakeKey := "6E26F0CA"

	tearDown := func() func() {
		_ = SetUtripKey(fakeCorpId, fakeKey)

		return func() {
			// restore copr id & key
		}
	}()
	defer tearDown()

	fakeUserPhone := "apitest@ssharing.com"
	fakeUserName := "测试"

	assert.Equal(t, expect, PackUtripUrl(fakeUserPhone, fakeUserName))
}

func TestSetUtripKey(t *testing.T) {
	//t.Parallel()

	invalid_parameter := [][]string{
		[]string{
			fakeUtripCode, fakeInvalidUtripKey,
		},
		[]string{
			fakeInvalidUtripCode, fakeUtripKey,
		},
	}

	for _, v := range invalid_parameter {
		err := SetUtripKey(v[0], v[1])
		assert.Equal(t, true,
			errors.Is(err, util.ParamError))
	}

	err := SetUtripKey(fakeUtripCode, fakeUtripKey)
	assert.Nil(t, err)
}
