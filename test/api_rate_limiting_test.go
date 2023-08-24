/*
ArvanCloud CDN Services

Testing RateLimitingApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package arvancloud

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "git.arvancloud.ir/arvancloud/cdn-go-sdk"
)

func Test_arvancloud_RateLimitingApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RateLimitingApiService RateLimitingActionsReprioritize", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string

		resp, httpRes, err := apiClient.RateLimitingApi.RateLimitingActionsReprioritize(context.Background(), domain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RateLimitingApiService RateLimitingRulesCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string

		resp, httpRes, err := apiClient.RateLimitingApi.RateLimitingRulesCreate(context.Background(), domain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RateLimitingApiService RateLimitingRulesDestroy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string
		var id string

		resp, httpRes, err := apiClient.RateLimitingApi.RateLimitingRulesDestroy(context.Background(), domain, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RateLimitingApiService RateLimitingRulesIndex", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string

		resp, httpRes, err := apiClient.RateLimitingApi.RateLimitingRulesIndex(context.Background(), domain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RateLimitingApiService RateLimitingRulesShow", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string
		var id string

		resp, httpRes, err := apiClient.RateLimitingApi.RateLimitingRulesShow(context.Background(), domain, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RateLimitingApiService RateLimitingRulesUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string
		var id string

		resp, httpRes, err := apiClient.RateLimitingApi.RateLimitingRulesUpdate(context.Background(), domain, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RateLimitingApiService RateLimitingSettingsIndex", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string

		resp, httpRes, err := apiClient.RateLimitingApi.RateLimitingSettingsIndex(context.Background(), domain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RateLimitingApiService RateLimitingSettingsUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string

		resp, httpRes, err := apiClient.RateLimitingApi.RateLimitingSettingsUpdate(context.Background(), domain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
