package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/CommuniTEAM/CommuniTEA/api"
	"github.com/pashagolub/pgxmock/v3"
)

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

func TestShouldGetUser(t *testing.T) {
	mock, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection pool", err)
	}
	defer mock.Close()

	mock.ExpectBegin()
	mock.ExpectExec("insert into users (id, role, username, first_name, last_name, email, password, location)
values (
  '423a8b41-eed4-423f-bf25-039e15bd471c',
  'admin',
  'admin',
  'brian',
  'la',
  'brianla@gmail.com',
  'password',
  '423a8b41-eed4-423f-bf25-039e15bd471b'
  )").WillReturnResult(pgxmock.NewResult("INSERT", 1))
	mock.ExpectExec("select 'id',
    'role',
    'username',
    'first_name',
    'last_name',
    'email',
    'location'
from users
where username = $1").
		WithArgs("admin").WillReturnError(fmt.Errorf("some error"))
	mock.ExpectCommit()

	// now we execute our method
	if err = api.GetUser(); err != nil {
		t.Errorf("error was not expected while updating: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// // create app with mocked db, request and response to test
	// app := &api.API{mock}
	// urlString := Sprintf("http://localhost:8000/users/%s", "username")
	// req, err := http.NewRequest("GET", urlString, nil)
	// if err != nil {
	// 	t.Fatalf("an error '%s' was not expected while creating request", err)
	// }
	// w := httptest.NewRecorder()

	// // before we actually execute our api function, we need to expect required DB actions
	// rows := mock.NewRows([]string{"id", "title", "body"}).
	// 	AddRow(1, "post 1", "hello").
	// 	AddRow(2, "post 2", "world")

	// mock.ExpectQuery("^SELECT (.+) FROM posts$").WillReturnRows(rows)

	// now we execute our request
	// app.GetUser()

	// if w.Code != 200 {
	// 	t.Fatalf("expected status code to be 200, but got: %d\nBody: %v", w.Code, w.Body)
	// }

	// data := struct {
	// 	Posts []*post
	// }{Posts: []*post{
	// 	{ID: 1, Title: "post 1", Body: "hello"},
	// 	{ID: 2, Title: "post 2", Body: "world"},
	// }}
	// app.assertJSON(w.Body.Bytes(), data, t)

	// // we make sure that all expectations were met
	// if err := mock.ExpectationsWereMet(); err != nil {
	// 	t.Errorf("there were unfulfilled expectations: %s", err)
	// }
}
