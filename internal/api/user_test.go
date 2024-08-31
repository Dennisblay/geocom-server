package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	mockdb "github.com/Dennisblay/ordering-app-server/internal/database/mock"
	"github.com/Dennisblay/ordering-app-server/internal/database/sqlc"
	"github.com/Dennisblay/ordering-app-server/util"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetUserByIDAPI(t *testing.T) {
	user := RandomUser()

	testCases := []struct {
		name          string
		userId        int32
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:   "Ok",
			userId: user.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUserById(gomock.Any(), gomock.Eq(user.ID)).
					Times(1).
					Return(user, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchUser(t, recorder.Body, user)
			},
		},
		{
			name:   "NotFound",
			userId: user.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUserById(gomock.Any(), gomock.Eq(user.ID)).
					Times(1).
					Return(db.GetUserByIdRow{}, pgx.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:   "InternalError",
			userId: user.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUserById(gomock.Any(), gomock.Eq(user.ID)).
					Times(1).
					Return(db.GetUserByIdRow{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:   "InvalidID",
			userId: 0,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUserById(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},

		// Todo: Add more cases later
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			// Start test server and send request
			server, err := NewServer(store)
			require.NoError(t, err)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/users/%d", tc.userId)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			// Check response
			tc.checkResponse(t, recorder)
		})
	}
}

func TestGetUsersAPI(t *testing.T) {
	var n int32 = 5
	users := RandomUsers(n)

	type Query struct {
		PageID   int32
		PageSize int32
	}

	testCases := []struct {
		name          string
		query         Query
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "Ok",
			query: Query{
				PageID:   1,
				PageSize: n,
			},
			buildStubs: func(store *mockdb.MockStore) {

				arg := db.GetUsersParams{
					Limit:  n,
					Offset: 0,
				}
				store.EXPECT().
					GetUsers(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(users, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchUsers(t, recorder.Body, users)
			},
		},
		{
			name: "InternalError",
			query: Query{
				PageID:   1,
				PageSize: n,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUsers(gomock.Any(), gomock.Any()).
					Times(1).
					Return([]db.GetUsersRow{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},

		// Todo: Add more cases later
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			// Start test server and send request
			server, err := NewServer(store)
			require.NoError(t, err)
			recorder := httptest.NewRecorder()

			url := "/users/"
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			// Add query parameters to request URL
			q := request.URL.Query()
			q.Add("page_id", fmt.Sprintf("%d", tc.query.PageID))
			q.Add("page_size", fmt.Sprintf("%d", tc.query.PageSize))
			request.URL.RawQuery = q.Encode()

			server.router.ServeHTTP(recorder, request)
			// Check response
			tc.checkResponse(t, recorder)
		})
	}
}

func requireBodyMatchUser(t *testing.T, body *bytes.Buffer, user db.GetUserByIdRow) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotUser db.GetUserByIdRow
	err = json.Unmarshal(data, &gotUser)
	require.NoError(t, err)

	require.Equal(t, user, gotUser)

}

func requireBodyMatchUsers(t *testing.T, body *bytes.Buffer, users []db.GetUsersRow) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotUsers []db.GetUsersRow
	err = json.Unmarshal(data, &gotUsers)
	require.NoError(t, err)

	require.Equal(t, users, gotUsers)

}

func RandomUser() db.GetUserByIdRow {

	return db.GetUserByIdRow{
		ID:        int32(util.RandomInt(1, 1000)),
		FirstName: util.RandomUser(),
		LastName:  util.RandomUser(),
		Email:     util.RandomEmail(),
		Phone:     util.RandomPhone(),
		Address:   util.RandomAddress(),
	}
}
func RandomUsers(n int32) []db.GetUsersRow {
	var users []db.GetUsersRow
	for _ = range n {
		users = append(users,
			db.GetUsersRow{
				ID:        int32(util.RandomInt(1, 1000)),
				FirstName: util.RandomUser(),
				LastName:  util.RandomUser(),
				Email:     util.RandomEmail(),
				Phone:     util.RandomPhone(),
				Address:   util.RandomAddress(),
				CreatedAt: time.Now().UTC(),
				UpdatedAt: time.Now().UTC(),
			},
		)
	}
	return users
}
