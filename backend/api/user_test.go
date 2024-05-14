package api

// func TestGetUserAPI(t *testing.T) {
// 	User := randomUser()

// 	testCases := []struct {
// 		name          string
// 		UserID     int64
// 		buildStubs    func(store *mockdb.MockStore)
// 		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
//     }{
// 			{
// 				name: "OK",
// 				UserID: User.ID,
// 				buildStubs: func(store *mockdb.MockStore) {
// 						store.EXPECT().
// 								GetUser(gomock.Any(), gomock.Eq(User.Username)).
// 								Times(1).
// 								Return(User, nil)
// 				},
// 				checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 						require.Equal(t, http.StatusOK, recorder.Code)
// 						requireBodyMatchUser(t, recorder.Body, User)
// 				},
// 			},
// 			{
// 				name: "NotFound",
// 				UserID: User.ID,
// 				buildStubs: func(store *mockdb.MockStore) {
// 						store.EXPECT().
// 								GetUser(gomock.Any(), gomock.Eq(User.Username)).
// 								Times(1).
// 								Return(db.User{}, sql.ErrNoRows)
// 				},
// 				checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 						require.Equal(t, http.StatusNotFound, recorder.Code)
// 				},
// 			},
// 			{
// 				name: "InternalError",
// 				UserID: User.ID,
// 				buildStubs: func(store *mockdb.MockStore) {
// 						store.EXPECT().
// 								GetUser(gomock.Any(), gomock.Eq(User.Username)).
// 								Times(1).
// 								Return(db.User{}, sql.ErrConnDone)
// 				},
// 				checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 						require.Equal(t, http.StatusInternalServerError, recorder.Code)
// 				},
// 			},
// 			{
// 				name: "InvalidID",
// 				UserID: 0,
// 				buildStubs: func(store *mockdb.MockStore) {
// 						store.EXPECT().
// 								GetUser(gomock.Any(), gomock.Any()).
// 								Times(0)
// 				},
// 				checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 						require.Equal(t, http.StatusBadRequest, recorder.Code)
// 				},
// 			},
//     }
//     for i := range testCases {
//         tc := testCases[i]

//         t.Run(tc.name, func(t *testing.T) {
//             ctrl := gomock.NewController(t)
//             defer ctrl.Finish()

//             store := mockdb.NewMockStore(ctrl)
//             tc.buildStubs(store)

//             server := newTestServer(t,store)
//             recorder := httptest.NewRecorder()

//             url := fmt.Sprintf("/users/%d", tc.UserID)
//             request, err := http.NewRequest(http.MethodGet, url, nil)
//             require.NoError(t, err)

//             server.router.ServeHTTP(recorder, request)
//             tc.checkResponse(t, recorder)
//         })
//     }
// }

// func randomUser() db.User {
//     return db.User{
//         ID:       util.RandomInt(1, 1000),
//         Username: util.RandomUsername(),
// 				Password: util.RandomPassword(),
// 				Email: util.RandomEmail(),
//     }
// }

// func requireBodyMatchUser(t *testing.T, body io.Reader, account db.User) {
//     data, err := io.ReadAll(body)
//     require.NoError(t, err)

//     var gotUser db.User
//     err = json.Unmarshal(data, &gotUser)
//     require.NoError(t, err)
//     require.Equal(t, account, gotUser)
// }