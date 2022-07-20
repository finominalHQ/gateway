package actions

import (
	"net/http"
	"os"
	"testing"

	"github.com/gobuffalo/suite/v4"
)

type ActionSuite struct {
	*suite.Action
}

func Test_ActionSuite(t *testing.T) {
	action, err := suite.NewActionWithFixtures(App(), os.DirFS("../fixtures"))
	if err != nil {
		t.Fatal(err)
	}

	as := &ActionSuite{
		Action: action,
	}
	suite.Run(t, as)
}

func (as *ActionSuite) Test_HomeHandler() {
	res := as.JSON("/").Get()

	as.Equal(http.StatusOK, res.Code)
	as.Contains(res.Body.String(), "Welcome to Buffalo")
}
