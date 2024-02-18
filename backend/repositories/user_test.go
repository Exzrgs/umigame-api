package repositories_test

import (
	"testing"
	"umigame-api/models"
	"umigame-api/repositories"
	"umigame-api/repositories/testdata"
)

func TestSelectUser_Basic_OK(t *testing.T) {
	tests := []struct {
		name     string
		email    string
		expected models.User
	}{
		{
			name:     "basic",
			email:    testdata.SelectUser_Basic[0].Email,
			expected: testdata.SelectUser_Basic[0],
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := repositories.SelectUser(db, tt.email)
			if err != nil {
				t.Fatal(err)
			}

			if user.Name != tt.expected.Name {
				t.Errorf("expected %v but got %v", tt.expected.Name, user.Name)
			}
			if user.Email != tt.expected.Email {
				t.Errorf("expected %v but got %v", tt.expected.Email, user.Email)
			}
			if user.PasswordHash != tt.expected.PasswordHash {
				t.Errorf("expected %v but got %v", tt.expected.PasswordHash, user.PasswordHash)
			}
			if user.UUID != tt.expected.UUID {
				t.Errorf("expected %v but got %v", tt.expected.UUID, user.UUID)
			}
			if user.IsValid != tt.expected.IsValid {
				t.Errorf("expected %v but got %v", tt.expected.IsValid, user.IsValid)
			}
		})
	}
}

func TestInsertUser_Basic_OK(t *testing.T) {
	tests := []struct {
		name     string
		user     models.User
		expected models.User
	}{
		{
			name:     "basic",
			user:     testdata.InsertUser_Basic[0],
			expected: testdata.InsertUser_Basic[0],
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repositories.InsertUser(db, &tt.user)
			if err != nil {
				t.Fatal(err)
			}

			got, err := repositories.SelectUser(db, tt.user.Email)
			if err != nil {
				t.Fatal(err)
			}

			if got.Name != tt.expected.Name {
				t.Errorf("inserted user is expected %+v but got %+v\n", tt.expected.Name, got.Name)
			}
			if got.Email != tt.expected.Email {
				t.Errorf("inserted user is expected %+v but got %+v\n", tt.expected.Email, got.Email)
			}
			if got.PasswordHash != tt.expected.PasswordHash {
				t.Errorf("inserted user is expected %+v but got %+v\n", tt.expected.PasswordHash, got.PasswordHash)
			}
			if got.UUID != tt.expected.UUID {
				t.Errorf("inserted user is expected %+v but got %+v\n", tt.expected.UUID, got.UUID)
			}
			if got.IsValid != tt.expected.IsValid {
				t.Errorf("inserted user is expected %+v but got %+v\n", tt.expected.IsValid, got.IsValid)
			}

			t.Cleanup(func() {
				const sqlStr = `DELETE FROM users WHERE email = ?`
				if _, err := db.Exec(sqlStr, tt.user.Email); err != nil {
					t.Fatal(err)
				}
			})
		})
	}
}

func TestInsertUser_Char_OK(t *testing.T) {
	tests := []struct {
		name     string
		user     models.User
		expected models.User
	}{
		{
			name:     "katakana",
			user:     testdata.InsertUser_Char[0],
			expected: testdata.InsertUser_Char[0],
		},
		{
			name:     "hiragana",
			user:     testdata.InsertUser_Char[1],
			expected: testdata.InsertUser_Char[1],
		},
		{
			name:     "kanji",
			user:     testdata.InsertUser_Char[2],
			expected: testdata.InsertUser_Char[2],
		},
		{
			name:     "half-katakana",
			user:     testdata.InsertUser_Char[3],
			expected: testdata.InsertUser_Char[3],
		},
		{
			name:     "symbol",
			user:     testdata.InsertUser_Char[4],
			expected: testdata.InsertUser_Char[4],
		},
		{
			name:     "space",
			user:     testdata.InsertUser_Char[5],
			expected: testdata.InsertUser_Char[5],
		},
		{
			name:     "special char",
			user:     testdata.InsertUser_Char[6],
			expected: testdata.InsertUser_Char[6],
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repositories.InsertUser(db, &tt.user)
			if err != nil {
				t.Fatal(err)
			}

			got, err := repositories.SelectUser(db, tt.user.Email)
			if err != nil {
				t.Fatal(err)
			}

			if got.Name != tt.expected.Name {
				t.Errorf("inserted user is expected %+v but got %+v\n", tt.expected.Name, got.Name)
			}
			if got.Email != tt.expected.Email {
				t.Errorf("inserted user is expected %+v but got %+v\n", tt.expected.Email, got.Email)
			}
			if got.PasswordHash != tt.expected.PasswordHash {
				t.Errorf("inserted user is expected %+v but got %+v\n", tt.expected.PasswordHash, got.PasswordHash)
			}
			if got.UUID != tt.expected.UUID {
				t.Errorf("inserted user is expected %+v but got %+v\n", tt.expected.UUID, got.UUID)
			}
			if got.IsValid != tt.expected.IsValid {
				t.Errorf("inserted user is expected %+v but got %+v\n", tt.expected.IsValid, got.IsValid)
			}

			t.Cleanup(func() {
				const sqlStr = `DELETE FROM users WHERE email = ?`
				if _, err := db.Exec(sqlStr, tt.user.Email); err != nil {
					t.Fatal(err)
				}
			})
		})
	}
}

func TestInsertUser_TooLong_NG(t *testing.T) {
	tests := []struct {
		name     string
		user     models.User
		expected string
	}{
		{
			name:     "too long name",
			user:     testdata.InsertUser_TooLong[0],
			expected: "cause insert error",
		},
		{
			name:     "too long email",
			user:     testdata.InsertUser_TooLong[1],
			expected: "cause insert error",
		},
		{
			name:     "too long uuid",
			user:     testdata.InsertUser_TooLong[2],
			expected: "cause insert error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repositories.InsertUser(db, &tt.user)
			if err == nil {
				t.Errorf("expected %v but got %v", tt.expected, tt.user)
			}

			t.Cleanup(func() {
				const sqlStr = `DELETE FROM users WHERE email = ?`
				if _, err := db.Exec(sqlStr, tt.user.Email); err != nil {
					t.Fatal(err)
				}
			})
		})
	}
}

// //////////////////////////////////////////////////////

func TestUpdateValid_Basic_OK(t *testing.T) {
	tests := []struct {
		name     string
		user     models.User
		expected bool
	}{
		{
			name:     "basic",
			user:     testdata.UpdateValid_Basic[0],
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.user.IsValid == true {
				t.Fatal("user is already valid")
			}

			err := repositories.InsertUser(db, &tt.user)
			if err != nil {
				t.Fatal(err)
			}

			if err := repositories.UpdateActivate(db, tt.user.UUID); err != nil {
				t.Fatal(err)
			}

			got, err := repositories.GetAuthInfo(db, tt.user.Email)
			if err != nil {
				t.Fatal(err)
			}

			if got.IsValid != tt.expected {
				t.Errorf("expected %v but got %v", tt.expected, got.IsValid)
			}
		})
	}
}

// //////////////////////////////////////////////////////

func TestGetAuthInfo_Basic_OK(t *testing.T) {
	tests := []struct {
		name     string
		email    string
		expected models.User
	}{
		{
			name:     "basic",
			email:    testdata.GetAuthInfo_Basic[0].Email,
			expected: testdata.GetAuthInfo_Basic[0],
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repositories.GetAuthInfo(db, tt.email)
			if err != nil {
				t.Fatal(err)
			}

			if got.PasswordHash != tt.expected.PasswordHash {
				t.Errorf("expected %v but got %v", got.PasswordHash, tt.expected.PasswordHash)
			}
			if got.UUID != tt.expected.UUID {
				t.Errorf("expected %v but got %v", got.UUID, tt.expected.UUID)
			}
			if got.IsValid != tt.expected.IsValid {
				t.Errorf("expected %v but got %v", got.IsValid, tt.expected.IsValid)
			}
		})
	}
}

// //////////////////////////////////////////////////////

// func TestUpdateEmail(t *testing.T) {}

// func getUser(db *sql.DB, id int) (models.User, error) {
// 	const sqlStr = `
// 	SELECT * FROM users
// 	WHERE id == ?;
// 	`

// 	var user models.User
// 	if err := db.QueryRow(sqlStr, id).Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.UUID, &user.IsValid, &user.CreatedAt); err != nil {
// 		return models.User{}, err
// 	}

// 	return user, nil
// }

// func getIsValidByUUID(db *sql.DB, uuid string) (bool, error) {
// 	const sqlStr = `
// 	SELECT is_valid from users
// 	where uuid = ?;
// 	`
// 	var isValid bool
// 	if err := db.QueryRow(sqlStr, uuid).Scan(&isValid); err != nil {
// 		return false, err
// 	}

// 	return isValid, nil
// }
