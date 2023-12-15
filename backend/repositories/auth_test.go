package repositories_test

import (
	"database/sql"
	"testing"
	"umigame-api/models"
	"umigame-api/repositories"
)

func TestInsertUser_Basic_OK(t *testing.T) {
	tests := []struct {
		title    string
		user     models.User
		expected models.User
	}{
		{
			title:    "basic",
			user:     testdata.InsertUser_Basic[0],
			expected: testdata.InsertUser_Basic[0],
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			user, err := repositories.InsertUser(db, test.user)
			if err != nil {
				t.Fatal(err)
			}

			got, err := getUser(db, user.ID)
			if err != nil {
				t.Fatal(err)
			}

			if got.Name != test.expected.Name {
				t.Errorf("inserted user is expected %+v but got %+v\n", test.expected.Name, got.Name)
			}
			if got.Email != test.expected.Email {
				t.Errorf("inserted user is expected %+v but got %+v\n", test.expected.Email, got.Email)
			}
			if got.PasswordHash != test.expected.PasswordHash {
				t.Errorf("inserted user is expected %+v but got %+v\n", test.expected.PasswordHash, got.PasswordHash)
			}
			if got.UUID != test.expected.UUID {
				t.Errorf("inserted user is expected %+v but got %+v\n", test.expected.UUID, got.UUID)
			}
			if got.IsValid != test.expected.IsValid {
				t.Errorf("inserted user is expected %+v but got %+v\n", test.expected.IsValid, got.IsValid)
			}
		})
	}
}

func TestInsertUser_NoData_NG(t *testing.T) {
	tests := []struct {
		title    string
		user     models.User
		expected models.User
	}{
		{
			title:    "no name data",
			user:     testdata.InsertUser_NoData[0],
			expected: "cause insert error",
		},
		{
			title:    "no email data",
			user:     testdata.InsertUser_NoData[1],
			expected: "cause insert error",
		},
		{
			title:    "no password_hash data",
			user:     testdata.InsertUser_NoData[2],
			expected: "cause insert error",
		},
		{
			title:    "no uuid data",
			user:     testdata.InsertUser_NoData[3],
			expected: "cause insert error",
		},
		{
			title:    "no is_valid data",
			user:     testdata.InsertUser_NoData[4],
			expected: "cause insert error",
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			user, err := repositories.InsertUser(db, test.user)
			if err == nil {
				t.Errorf("expected %v but got %v", test.expected, test.user)
			}
		})
	}
}

func TestInsertUser_Char_OK(t *testing.T) {
	tests := []struct {
		title    string
		user     models.User
		expected models.User
	}{
		{
			title:    "katakana",
			user:     testdata.InsertUser_Char[0],
			expected: testdata.InsertUser_Char[0],
		},
		{
			title:    "hiragana",
			user:     testdata.InsertUser_Char[1],
			expected: testdata.InsertUser_Char[1],
		},
		{
			title:    "kanji",
			user:     testdata.InsertUser_Char[2],
			expected: testdata.InsertUser_Char[2],
		},
		{
			title:    "half-katakana",
			user:     testdata.InsertUser_Char[3],
			expected: testdata.InsertUser_Char[3],
		},
		{
			title:    "symbol",
			user:     testdata.InsertUser_Char[4],
			expected: testdata.InsertUser_Char[4],
		},
		{
			title:    "space",
			user:     testdata.InsertUser_Char[5],
			expected: testdata.InsertUser_Char[5],
		},
		{
			title:    "special char",
			user:     testdata.InsertUser_Char[6],
			expected: testdata.InsertUser_Char[6],
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			user, err := repositories.InsertUser(db, test.user)
			if err != nil {
				t.Fatal(err)
			}

			got, err := getUser(db, user.ID)
			if err != nil {
				t.Fatal(err)
			}

			if got.Name != test.expected.Name {
				t.Errorf("inserted user is expected %+v but got %+v\n", test.expected.Name, got.Name)
			}
			if got.Email != test.expected.Email {
				t.Errorf("inserted user is expected %+v but got %+v\n", test.expected.Email, got.Email)
			}
			if got.PasswordHash != test.expected.PasswordHash {
				t.Errorf("inserted user is expected %+v but got %+v\n", test.expected.PasswordHash, got.PasswordHash)
			}
			if got.UUID != test.expected.UUID {
				t.Errorf("inserted user is expected %+v but got %+v\n", test.expected.UUID, got.UUID)
			}
			if got.IsValid != test.expected.IsValid {
				t.Errorf("inserted user is expected %+v but got %+v\n", test.expected.IsValid, got.IsValid)
			}
		})
	}
}

func TestInsertUser_TooLong_NG(t *testing.T) {
	tests := []struct {
		title    string
		user     models.User
		expected models.User
	}{
		{
			title:    "too long name",
			user:     testdata.InsertUser_TooLong[0],
			expected: "cause insert error",
		},
		{
			title:    "too long email",
			user:     testdata.InsertUser_TooLong[1],
			expected: "cause insert error",
		},
		{
			title:    "too long password_hash",
			user:     testdata.InsertUser_TooLong[2],
			expected: "cause insert error",
		},
		{
			title:    "too long uuid",
			user:     testdata.InsertUser_TooLong[3],
			expected: "cause insert error",
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			user, err := repositories.InsertUser(db, test.user)
			if err == nil {
				t.Errorf("expected %v but got %v", test.expected, test.user)
			}
		})
	}
}

//////////////////////////////////////////////////////

func TestUpdateValid_Basic_OK(t *testing.T) {
	tests := []struct {
		title    string
		uuid     string
		expected models.User
	}{
		{
			title:    "basic",
			uuid:     testdata.UpdateValid_Basic[0].UUID,
			expected: testdata.UpdateValid_Basic[0].IsValid,
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			if err := repositories.UpdateValid(test.uuid); err != nil {
				t.Fatal(err)
			}

			got, err := getIsValidByUUID(db, test.uuid)
			if err != nil {
				t.Fatal(err)
			}

			if got != test.expected {
				t.Errorf("expected %v but got %v", got, expected)
			}
		})
	}
}

//////////////////////////////////////////////////////

func TestGetAuthInfo_Basic_OK(t *testing.T) {
	tests := []struct {
		title    string
		email    string
		expected models.User
	}{
		{
			title:    "basic",
			email:    testdata.GetAuthInfo_Basic[0].Email,
			expected: testdata.GetAuthInfo_Basic[0],
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			got, err := repositories.GetAuthInfo(test.email)
			if err != nil {
				t.Fatal(err)
			}

			if got.Name != test.expected.Name {
				t.Errorf("expected %v but got %v", got.Name, test.expected.Name)
			}
			if got.PasswordHash != test.expected.PasswordHash {
				t.Errorf("expected %v but got %v", got.PasswordHash, test.expected.PasswordHash)
			}
			if got.UUID != test.expected.UUID {
				t.Errorf("expected %v but got %v", got.UUID, test.expected.UUID)
			}
			if got.IsValid != test.expected.IsValid {
				t.Errorf("expected %v but got %v", got.IsValid, test.expected.IsValid)
			}
		})
	}
}

//////////////////////////////////////////////////////

func TestUpdateEmail(t *testing.T) {}

func getUser(db *sql.DB, id int) (models.User, error) {
	const sqlStr = `
	SELECT * FROM users
	WHERE id == ?;
	`

	var user models.User
	if err := db.QueryRow(sqlStr, id).Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.UUID, &user.IsValid, &user.CreatedAt); err != nil {
		return models.User{}, err
	}

	return user, nil
}

func getIsValidByUUID(db *sql.DB, uuid string) (bool, error) {
	const sqlStr = `
	SELECT is_valid from users
	where uuid = ?;
	`
	var isValid bool
	if err := db.QueryRow(sqlStr, uuid).Scan(&isValid); err != nil {
		return false, err
	}

	return isValid, nil
}
