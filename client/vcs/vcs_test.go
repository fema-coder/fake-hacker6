package vcs

import (
	"github.com/innogames/slack-bot/bot/config"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitBranchWatcher(t *testing.T) {
	cfg := config.Config{}
	logger := logrus.New()

	abort := InitBranchWatcher(cfg, logger)
	abort <- true
}

func TestGetMatchingBranches(t *testing.T) {
	logger, _ = test.NewNullLogger()

	branches = []string{
		"master",
		"feature/PROJ-1234-do-something",
		"feature/PROJ-1234-do-something-hotfix",
		"bugfix/PROJ-1235-fixed",
		"release/3.12.23",
	}

	t.Run("Empty", func(t *testing.T) {
		actual, err := GetMatchingBranch("")
		assert.NotNil(t, err)
		assert.Equal(t, "", actual)
	})

	t.Run("Not found", func(t *testing.T) {
		actual, err := GetMatchingBranch("this-might-be-a-branch")
		assert.Equal(t, "this-might-be-a-branch", actual)
		assert.Nil(t, err)
	})

	t.Run("Not unique", func(t *testing.T) {
		actual, err := GetMatchingBranch("PROJ-1234")
		assert.Equal(t, "multiple branches found: feature/PROJ-1234-do-something, feature/PROJ-1234-do-something-hotfix", err.Error())
		assert.Equal(t, "", actual)
	})

	t.Run("Test unique branches", func(t *testing.T) {
		actual, err := GetMatchingBranch("master")
		assert.Equal(t, "master", actual)
		assert.Nil(t, err)

		actual, err = GetMatchingBranch("PROJ-1235")
		assert.Equal(t, "bugfix/PROJ-1235-fixed", actual)
		assert.Nil(t, err)

		actual, err = GetMatchingBranch("feature/PROJ-1234-do-something")
		assert.Equal(t, "feature/PROJ-1234-do-something", actual)
		assert.Nil(t, err)
	})
}
