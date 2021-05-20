package gofiber

import (
	"net/http"
	"testing"

	"github.com/HongJaison/go-admin4/tests/common"
	"github.com/gavv/httpexpect"
)

func TestGofiber(t *testing.T) {
	common.ExtraTest(httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewFastBinder(newHandler()),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
	}))
}