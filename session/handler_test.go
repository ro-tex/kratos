package session_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/x/urlx"

	"github.com/ory/viper"

	"github.com/ory/kratos/driver/configuration"
	"github.com/ory/kratos/internal"
	"github.com/ory/kratos/internal/testhelpers"
	. "github.com/ory/kratos/session"
	"github.com/ory/kratos/x"
)

func init() {
	internal.RegisterFakes()
}

func send(code int) httprouter.Handle {
	return func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		w.WriteHeader(code)
	}
}

func TestSessionWhoAmI(t *testing.T) {
	t.Run("public", func(t *testing.T) {
		_, reg := internal.NewFastRegistryWithMocks(t)
		r := x.NewRouterPublic()

		// set this intermediate because kratos needs some valid url for CRUDE operations
		viper.Set(configuration.ViperKeyURLsSelfPublic, "http://example.com")
		h, _ := testhelpers.MockSessionCreateHandler(t, reg)
		r.GET("/set", h)

		NewHandler(reg).RegisterPublicRoutes(r)
		ts := httptest.NewServer(r)
		defer ts.Close()

		viper.Set(configuration.ViperKeyURLsSelfPublic, ts.URL)

		client := testhelpers.MockCookieClient(t)

		// No cookie yet -> 401
		res, err := client.Get(ts.URL + SessionsWhoamiPath)
		require.NoError(t, err)
		assert.EqualValues(t, http.StatusUnauthorized, res.StatusCode)

		// Set cookie
		testhelpers.MockHydrateCookieClient(t, client, ts.URL+"/set")

		// Cookie set -> 200 (GET)
		for _, method := range []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
		} {
			t.Run("http_method="+method, func(t *testing.T) {
				req, err := http.NewRequest(method, ts.URL+SessionsWhoamiPath, nil)
				require.NoError(t, err)

				res, err = client.Do(req)
				require.NoError(t, err)
				assert.EqualValues(t, http.StatusOK, res.StatusCode)
			})
		}
	})
}

func TestIsNotAuthenticatedSecurecookie(t *testing.T) {
	_, reg := internal.NewFastRegistryWithMocks(t)
	r := x.NewRouterPublic()
	r.GET("/public/with-callback", reg.SessionHandler().IsNotAuthenticated(send(http.StatusOK), send(http.StatusBadRequest)))

	ts := httptest.NewServer(r)
	defer ts.Close()
	viper.Set(configuration.ViperKeyURLsSelfPublic, ts.URL)

	c := testhelpers.MockCookieClient(t)
	c.Jar.SetCookies(urlx.ParseOrPanic(ts.URL), []*http.Cookie{
		{
			Name: DefaultSessionCookieName,
			// This is an invalid cookie because it is generated by a very random secret
			Value:    "MTU3Mjg4Njg0MXxEdi1CQkFFQ180SUFBUkFCRUFBQU52LUNBQUVHYzNSeWFXNW5EQVVBQTNOcFpBWnpkSEpwYm1jTUd3QVpUWFZXVUhSQlZVeExXRWRUUmxkVVoyUkpUVXhzY201SFNBPT187kdI3dMP-ep389egDR2TajYXGG-6xqC2mAlgnBi0vsg=",
			HttpOnly: true,
			Path:     "/",
			Expires:  time.Now().Add(time.Hour),
		},
	})

	res, err := c.Get(ts.URL + "/public/with-callback")
	require.NoError(t, err)

	assert.EqualValues(t, http.StatusOK, res.StatusCode)
}

func TestIsNotAuthenticated(t *testing.T) {
	_, reg := internal.NewFastRegistryWithMocks(t)
	r := x.NewRouterPublic()
	// set this intermediate because kratos needs some valid url for CRUDE operations
	viper.Set(configuration.ViperKeyURLsSelfPublic, "http://example.com")

	reg.WithCSRFHandler(new(x.FakeCSRFHandler))
	h, _ := testhelpers.MockSessionCreateHandler(t, reg)
	r.GET("/set", h)
	r.GET("/public/with-callback", reg.SessionHandler().IsNotAuthenticated(send(http.StatusOK), send(http.StatusBadRequest)))
	r.GET("/public/without-callback", reg.SessionHandler().IsNotAuthenticated(send(http.StatusOK), nil))
	ts := httptest.NewServer(r)
	defer ts.Close()
	viper.Set(configuration.ViperKeyURLsSelfPublic, ts.URL)

	sessionClient := testhelpers.MockCookieClient(t)
	testhelpers.MockHydrateCookieClient(t, sessionClient, ts.URL+"/set")

	for k, tc := range []struct {
		c    *http.Client
		call string
		code int
	}{
		{
			c:    sessionClient,
			call: "/public/with-callback",
			code: http.StatusBadRequest,
		},
		{
			c:    http.DefaultClient,
			call: "/public/with-callback",
			code: http.StatusOK,
		},

		{
			c:    sessionClient,
			call: "/public/without-callback",
			code: http.StatusForbidden,
		},
		{
			c:    http.DefaultClient,
			call: "/public/without-callback",
			code: http.StatusOK,
		},
	} {
		t.Run(fmt.Sprintf("case=%d", k), func(t *testing.T) {
			res, err := tc.c.Get(ts.URL + tc.call)
			require.NoError(t, err)

			assert.EqualValues(t, tc.code, res.StatusCode)
		})
	}
}

func TestIsAuthenticated(t *testing.T) {
	_, reg := internal.NewFastRegistryWithMocks(t)
	reg.WithCSRFHandler(new(x.FakeCSRFHandler))
	r := x.NewRouterPublic()
	// set this intermediate because kratos needs some valid url for CRUDE operations
	viper.Set(configuration.ViperKeyURLsSelfPublic, "http://example.com")

	h, _ := testhelpers.MockSessionCreateHandler(t, reg)
	r.GET("/set", h)
	r.GET("/privileged/with-callback", reg.SessionHandler().IsAuthenticated(send(http.StatusOK), send(http.StatusBadRequest)))
	r.GET("/privileged/without-callback", reg.SessionHandler().IsAuthenticated(send(http.StatusOK), nil))
	ts := httptest.NewServer(r)
	defer ts.Close()
	viper.Set(configuration.ViperKeyURLsSelfPublic, ts.URL)

	sessionClient := testhelpers.MockCookieClient(t)
	testhelpers.MockHydrateCookieClient(t, sessionClient, ts.URL+"/set")

	for k, tc := range []struct {
		c    *http.Client
		call string
		code int
	}{
		{
			c:    sessionClient,
			call: "/privileged/with-callback",
			code: http.StatusOK,
		},
		{
			c:    http.DefaultClient,
			call: "/privileged/with-callback",
			code: http.StatusBadRequest,
		},

		{
			c:    sessionClient,
			call: "/privileged/without-callback",
			code: http.StatusOK,
		},
		{
			c:    http.DefaultClient,
			call: "/privileged/without-callback",
			code: http.StatusForbidden,
		},
	} {
		t.Run(fmt.Sprintf("case=%d", k), func(t *testing.T) {
			res, err := tc.c.Get(ts.URL + tc.call)
			require.NoError(t, err)

			assert.EqualValues(t, tc.code, res.StatusCode)
		})
	}
}
