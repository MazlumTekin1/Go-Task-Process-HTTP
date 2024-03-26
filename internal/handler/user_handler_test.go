package handler_test

// func TestUserHandler_Create(t *testing.T) {
// 	mockRepo := new(mock.MockUserRepository)

// 	user := domain.UserAddReq{FirstName: "Test User", LastName: "Test Last Name", Email: "test@test.com", Password: "Test Password", CreateUserId: 1}
// 	mockRepo.On("Create", user).Return(1, nil)

// 	handler := handler.NewUserHandler(mockRepo)

// 	body, _ := json.Marshal(user)
// 	req, err := http.NewRequest("POST", "/users/add", bytes.NewBuffer(body))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	rr := httptest.NewRecorder()

// 	handler.Create(rr, req)

// 	assert.Equal(t, http.StatusCreated, rr.Code)

// 	expectedBody := `{"id":1}`
// 	assert.JSONEq(t, expectedBody, rr.Body.String())
// }

// func TestUserHandler_Update(t *testing.T) {
// 	mockRepo := new(mock.MockUserRepository)

// 	user := domain.UserUpdateReq{Id: 1, FirstName: "Test User", LastName: "Test Last Name", Email: "test.user@test.com", Password: "Test Password", UpdateUserId: 1}
// 	mockRepo.On("Update", user).Return(1, nil)

// 	handler := handler.NewUserHandler(mockRepo)

// 	body, _ := json.Marshal(user)
// 	req, err := http.NewRequest("PUT", "/users/update", bytes.NewBuffer(body))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	rr := httptest.NewRecorder()

// 	handler.Update(rr, req)

// 	assert.Equal(t, http.StatusOK, rr.Code)

// 	expectedBody := `{"id":1}`
// 	assert.JSONEq(t, expectedBody, rr.Body.String())
// }

// func TestUserHandler_Delete(t *testing.T) {
// 	mockRepo := new(mock.MockUserRepository)

// 	user := domain.UserDeleteReq{Id: 1, UpdateUserId: 1}
// 	mockRepo.On("Delete", user).Return(1, nil)

// 	handler := handler.NewUserHandler(mockRepo)

// 	body, _ := json.Marshal(user)
// 	req, err := http.NewRequest("DELETE", "/users/delete", bytes.NewBuffer(body))
// 	if err != nil {
// 		t.Fatal("Hata olustu: ", err)
// 	}

// 	rr := httptest.NewRecorder()

// 	handler.Delete(rr, req)

// 	assert.Equal(t, http.StatusOK, rr.Code)

// 	expectedBody := `{"id":1}`
// 	assert.JSONEq(t, expectedBody, rr.Body.String())
// }

// func TestUserHandler_GetById(t *testing.T) {
// 	mockRepo := new(mock.MockUserRepository)

// 	userRes := domain.UserGetDataList{
// 		Id:        1,
// 		FirstName: "Mazlum",
// 		LastName:  "Tekin",
// 		Email:     "mazlumtekin.kariyer@gmail.com",
// 		CreatedAt: "2024-03-23 15:09:01",
// 	}
// 	user := domain.UserGetByIdReq{Id: 1}

// 	mockRepo.On("GetById", user).Return(userRes, nil)

// 	handler := handler.NewUserHandler(mockRepo)
// 	body, _ := json.Marshal(user)
// 	req, err := http.NewRequest("GET", "/users/getById", bytes.NewBuffer(body))
// 	if err != nil {
// 		t.Fatal("Error creating request:", err)
// 	}

// 	rr := httptest.NewRecorder()

// 	handler.GetById(rr, req)

// 	assert.Equal(t, http.StatusOK, rr.Code)

// 	var actual domain.UserGetDataList
// 	err = json.Unmarshal(rr.Body.Bytes(), &actual)
// 	if err != nil {
// 		t.Fatal("Error decoding response body:", err)
// 	}

// 	assert.Equal(t, userRes, actual)
// }

// func TestUserHandler_GetAll(t *testing.T) {
// 	mockRepo := new(mock.MockUserRepository)

// 	users := []domain.UserGetDataList{
// 		{
// 			Id:        1,
// 			FirstName: "Mazlum",
// 			LastName:  "Tekin",
// 			Email:     "mazlumtekin.kariyer@gmail.com",
// 			CreatedAt: "2024-03-23 15:09:01.000"},
// 	}
// 	mockRepo.On("GetAll").Return(users, nil)

// 	handler := handler.NewUserHandler(mockRepo)

// 	req, err := http.NewRequest("GET", "/users/getAll", nil)
// 	if err != nil {
// 		t.Fatal("Error creating request:", err)
// 	}

// 	rr := httptest.NewRecorder()

// 	handler.GetAll(rr, req)

// 	assert.Equal(t, http.StatusOK, rr.Code)

// 	var actual []domain.UserGetDataList
// 	err = json.Unmarshal(rr.Body.Bytes(), &actual)
// 	if err != nil {
// 		t.Fatal("Error decoding response body:", err)
// 	}

// 	assert.Equal(t, users, actual)
// }
