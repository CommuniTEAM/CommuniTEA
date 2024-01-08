package api_test

// import (
// 	"bytes"
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/google/uuid"
// 	"github.com/jackc/pgx/v5/pgtype"
// 	"github.com/pashagolub/pgxmock/v3"
// )

// type TestCase struct {
// 	input         struct{}
// 	expected      struct{}
// 	actualOutcome struct{}
// }

// func TestCreateUser(t *testing.T) {
// 	// testCase := TestCase{}
// 	mock, err := pgxmock.NewConn()
// 	if err != nil {
// 		t.Fail()
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer mock.Close(context.Background())

// 	// create app with mocked db, request and response to test

// 	// app := &api{mock}
// 	testUUID, err := uuid.NewRandom()
// 	if err != nil {
// 		log.Println(fmt.Errorf("could not generate new uuid: %w", err))
// 	}

// 	postUserBody := newUserInput{
// 		ID:            pgtype.UUID{Bytes: testUUID, Valid: true},
// 		Role:          "user",
// 		Username:      "test1",
// 		FirstName:     "test",
// 		LastName:      "one",
// 		Email:         "testone@gmail.com",
// 		Password:      "password",
// 		LocationCity:  "Seattle",
// 		LocationState: "WA",
// 	}

// 	var buf bytes.Buffer
// 	bodyErr := json.NewEncoder(&buf).Encode(postUserBody)
// 	if bodyErr != nil {
// 		log.Fatalf("unable to convert json to bytes due to error: '%s'", bodyErr)
// 	}

// 	req, err := http.NewRequest(http.MethodPost, "http://localhost:8000/users", &buf)
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected while creating request", err)
// 	}
// 	w := httptest.NewRecorder()

// 	// before we actually execute our api function, we need to expect required DB actions

// 	rows := mock.NewRows([]string{
// "id", "role", "username", "first_name", "last_name", "email", "password", "location"
// }).AddRow()

// 	// mock.ExpectBegin()
// 	// mock.ExpectExec("Insert into locations_cities").WillReturnResult(pgxmock.NewResult("Insert", 1))
// 	// mock.ExpectExec("Insert into users")
// }
