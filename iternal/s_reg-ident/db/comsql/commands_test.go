package comsql_test

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"

	"pet/iternal/s_reg-ident/db/comsql"
)

func TestCheckUinquessLogin(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating sqlmock DB: %v", err)
	}
	defer db.Close()

	name := &[]struct {
		numtest     string
		dataDB      string
		dataclient  string
		correctresp error
	}{
		{numtest: "test_1", dataDB: "mark", dataclient: "mark", correctresp: fmt.Errorf("login already exists")},
		{numtest: "test_2", dataDB: "", dataclient: "bob", correctresp: nil},
		{numtest: "test_3", dataDB: "bob", dataclient: "bob", correctresp: fmt.Errorf("login already exists")},
		{numtest: "test_4", dataDB: "les", dataclient: "les", correctresp: fmt.Errorf("login already exists")},
		{numtest: "test_5", dataDB: "", dataclient: "tim", correctresp: nil},
	}
	for _, n := range *name {
		t.Run(n.numtest, func(t *testing.T) {
			rows := sqlmock.NewRows([]string{"logname"})
			if n.dataDB != "" {
				rows.AddRow(n.dataDB)
			}

			query := fmt.Sprintf(`SELECT log_name FROM account WHERE log_name = '%s';`, n.dataclient)
			mock.ExpectQuery(query).WillReturnRows(rows)

			req, err := http.NewRequest("POST", "/test", nil)
			if err != nil {
				t.Fatalf("failed to create request: %v", err)
			}

			req.Form = url.Values{}
			req.Form.Add("name", n.dataclient)

			resolt := comsql.CheckUinquenessLogin(db, req)

			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}

			if !reflect.DeepEqual(resolt, n.correctresp) {
				t.Errorf("\nbd log_name:	%s\nclient log_name:	%s\nresolt:	 %s", n.dataDB, n.dataclient, resolt)
			}
		})
	}
}

func TestCheckUinquessEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating sqlmock DB: %v", err)
	}
	defer db.Close()

	name := &[]struct {
		numtest     string
		dataDB      string
		dataclient  string
		correctresp error
	}{
		{numtest: "test_1", dataDB: "mark@test.com", dataclient: "mark@test.com", correctresp: fmt.Errorf("email already exists")},
		{numtest: "test_2", dataDB: "", dataclient: "bob@test.com", correctresp: nil},
		{numtest: "test_3", dataDB: "bob@test.com", dataclient: "bob@test.com", correctresp: fmt.Errorf("email already exists")},
		{numtest: "test_4", dataDB: "les@test.com", dataclient: "les@test.com", correctresp: fmt.Errorf("email already exists")},
		{numtest: "test_5", dataDB: "", dataclient: "tim@test.com", correctresp: nil},
	}
	for _, n := range *name {
		t.Run(n.numtest, func(t *testing.T) {
			rows := sqlmock.NewRows([]string{"email"})
			if n.dataDB != "" {
				rows.AddRow(n.dataDB)
			}

			query := fmt.Sprintf(`SELECT email FROM account WHERE email = '%s';`, n.dataclient)
			mock.ExpectQuery(query).WillReturnRows(rows)

			req, err := http.NewRequest("POST", "/test", nil)
			if err != nil {
				t.Fatalf("failed to create request: %v", err)
			}

			req.Form = url.Values{}
			req.Form.Add("email", n.dataclient)

			resolt := comsql.CheckUinquenessEmail(db, req)

			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}

			if !reflect.DeepEqual(resolt, n.correctresp) {
				t.Errorf("\nbd log_name:	%s\nclient log_name:	%s\nresolt:	 %s", n.dataDB, n.dataclient, resolt)
			}
		})
	}
}

// func TestSendRegdata(t *testing.T) {
// 	db, mock, err := sqlmock.New()
//     if err != nil {
//         t.Fatalf("error creating sqlmock DB: %v", err)
//     }
//     defer db.Close()

// }
