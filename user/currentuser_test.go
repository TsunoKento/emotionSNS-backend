package user

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCurrentUser(t *testing.T) {
	//setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()

	store := sessions.NewCookieStore()
	store.Options.Path = "/"
	sess, err := store.New(req, "session")
	if err != nil {
		t.Fatal("セッションを生成できませんでした", err)
	}
	sess.Values["userId"] = "kentots"
	sess.Values["auth"] = true
	sess.Save(req, rec)

	c := e.NewContext(req, rec)

	if assert.NoError(t, CurrentUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "{\"userId\":\"kentots\",\"name\":\"Kento\",\"picture\":\"https://images.unsplash.com/photo-1416339306562-f3d12fefd36f\"}\n", rec.Body.String())
	}

}
