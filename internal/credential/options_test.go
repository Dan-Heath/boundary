package credential

import (
	"testing"

	"github.com/hashicorp/boundary/internal/iam"
	"github.com/hashicorp/boundary/internal/util/template"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_GetOpts(t *testing.T) {
	t.Parallel()
	t.Run("WithTemplateData", func(t *testing.T) {
		opts := getDefaultOptions()
		assert.Empty(t, opts.WithTemplateData)
		opts, err := GetOpts(WithTemplateData(template.Data{User: template.User{Id: "foo"}}))
		require.NoError(t, err)
		assert.Equal(t, "foo", opts.WithTemplateData.User.Id)
	})
	t.Run("WithIamRepoFn", func(t *testing.T) {
		opts := getDefaultOptions()
		assert.Empty(t, opts.WithIamRepoFn)
		opts, err := GetOpts(WithIamRepoFn(func() (*iam.Repository, error) { return nil, nil }))
		require.NoError(t, err)
		assert.NotEmpty(t, opts.WithIamRepoFn)
	})
}