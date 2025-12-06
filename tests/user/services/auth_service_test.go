package services_test

import (
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/farhapartex/dealer-marketplace-be/apps/user/dtos"
	"github.com/farhapartex/dealer-marketplace-be/apps/user/services"
	"github.com/farhapartex/dealer-marketplace-be/config"
	"github.com/farhapartex/dealer-marketplace-be/pkg/database"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func stringPtr(s string) *string {
	return &s
}

func setupTestDB(t *testing.T) (sqlmock.Sqlmock, func()) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	dialector := postgres.New(postgres.Config{
		Conn:       db,
		DriverName: "postgres",
	})

	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	require.NoError(t, err)

	database.DB = gormDB

	cleanup := func() {
		db.Close()
	}

	return mock, cleanup
}

func setupTestConfig() {
	config.AppSettings = &config.Settings{
		EncryptKey:     "test-key",
		AuthCodeExpiry: "1HR",
		JWTSecretKey:   "test-jwt-secret-key-for-testing-purposes",
		Database: config.DatabaseSettings{
			Name:     "test_db",
			Host:     "localhost",
			Port:     5432,
			User:     "test",
			Password: "test",
		},
	}
}

func TestCreateNewUser_Success(t *testing.T) {
	mock, cleanup := setupTestDB(t)
	defer cleanup()
	setupTestConfig()

	service := services.NewAuthService()
	userID := uuid.New()
	authID := uuid.New()

	payload := dtos.CreateUserPayload{
		Name:     "John Doe",
		Email:    stringPtr("john@example.com"),
		Phone:    stringPtr("1234567890"),
		Password: "password123",
	}

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users"`)).
		WithArgs(
			payload.Name,
			payload.Email,
			payload.Phone,
			sqlmock.AnyArg(),
			false,
			false,
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
		).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at"}).
			AddRow(userID, time.Now(), time.Now()))
	mock.ExpectCommit()

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "auth_verification"`)).
		WithArgs(
			sqlmock.AnyArg(),
			userID,
			false,
			sqlmock.AnyArg(),
		).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at"}).
			AddRow(authID, time.Now()))
	mock.ExpectCommit()

	result, err := service.CreateNewUser(payload)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotEmpty(t, result.Code)
	assert.NotEmpty(t, result.Token)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCreateNewUser_DatabaseError(t *testing.T) {
	mock, cleanup := setupTestDB(t)
	defer cleanup()
	setupTestConfig()

	service := services.NewAuthService()

	payload := dtos.CreateUserPayload{
		Name:     "John Doe",
		Email:    stringPtr("john@example.com"),
		Phone:    stringPtr("1234567890"),
		Password: "password123",
	}

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users"`)).
		WillReturnError(errors.New("database error"))
	mock.ExpectRollback()

	result, err := service.CreateNewUser(payload)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "failed to create user")
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCreateNewUser_AuthVerificationError(t *testing.T) {
	mock, cleanup := setupTestDB(t)
	defer cleanup()
	setupTestConfig()

	service := services.NewAuthService()
	userID := uuid.New()

	payload := dtos.CreateUserPayload{
		Name:     "John Doe",
		Email:    stringPtr("john@example.com"),
		Phone:    stringPtr("1234567890"),
		Password: "password123",
	}

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users"`)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(userID))
	mock.ExpectCommit()

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "auth_verification"`)).
		WillReturnError(errors.New("auth verification error"))
	mock.ExpectRollback()

	result, err := service.CreateNewUser(payload)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "failed to create auth verification")
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestVerifyAuthCode_Success(t *testing.T) {
	mock, cleanup := setupTestDB(t)
	defer cleanup()
	setupTestConfig()

	service := services.NewAuthService()
	userID := uuid.New()
	code := "ABC123"
	authID := uuid.New()

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "auth_verification"`)).
		WithArgs(userID, code, 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "code", "user_id", "is_used", "expire_at", "created_at"}).
			AddRow(authID, code, userID, false, time.Now().Add(1*time.Hour), time.Now()))

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "auth_verification"`)).
		WithArgs(true, authID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "users"`)).
		WithArgs(true, sqlmock.AnyArg(), userID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := service.VerifyAuthCode(userID, code)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestVerifyAuthCode_InvalidCode(t *testing.T) {
	mock, cleanup := setupTestDB(t)
	defer cleanup()
	setupTestConfig()

	service := services.NewAuthService()
	userID := uuid.New()
	code := "INVALID"

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "auth_verification"`)).
		WithArgs(userID, code, 1).
		WillReturnError(gorm.ErrRecordNotFound)

	err := service.VerifyAuthCode(userID, code)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid verification code")
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestVerifyAuthCode_AlreadyUsed(t *testing.T) {
	mock, cleanup := setupTestDB(t)
	defer cleanup()
	setupTestConfig()

	service := services.NewAuthService()
	userID := uuid.New()
	code := "ABC123"
	authID := uuid.New()

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "auth_verification"`)).
		WithArgs(userID, code, 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "code", "user_id", "is_used", "expire_at", "created_at"}).
			AddRow(authID, code, userID, true, time.Now().Add(1*time.Hour), time.Now()))

	err := service.VerifyAuthCode(userID, code)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "already been used")
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestVerifyAuthCode_Expired(t *testing.T) {
	mock, cleanup := setupTestDB(t)
	defer cleanup()
	setupTestConfig()

	service := services.NewAuthService()
	userID := uuid.New()
	code := "ABC123"
	authID := uuid.New()

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "auth_verification"`)).
		WithArgs(userID, code, 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "code", "user_id", "is_used", "expire_at", "created_at"}).
			AddRow(authID, code, userID, false, time.Now().Add(-1*time.Hour), time.Now()))

	err := service.VerifyAuthCode(userID, code)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "expired")
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestVerifyAuthCode_UpdateVerificationError(t *testing.T) {
	mock, cleanup := setupTestDB(t)
	defer cleanup()
	setupTestConfig()

	service := services.NewAuthService()
	userID := uuid.New()
	code := "ABC123"
	authID := uuid.New()

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "auth_verification"`)).
		WithArgs(userID, code, 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "code", "user_id", "is_used", "expire_at", "created_at"}).
			AddRow(authID, code, userID, false, time.Now().Add(1*time.Hour), time.Now()))

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "auth_verification"`)).
		WithArgs(true, authID).
		WillReturnError(errors.New("update error"))
	mock.ExpectRollback()

	err := service.VerifyAuthCode(userID, code)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to update verification status")
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestVerifyAuthCode_ActivateUserError(t *testing.T) {
	mock, cleanup := setupTestDB(t)
	defer cleanup()
	setupTestConfig()

	service := services.NewAuthService()
	userID := uuid.New()
	code := "ABC123"
	authID := uuid.New()

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "auth_verification"`)).
		WithArgs(userID, code, 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "code", "user_id", "is_used", "expire_at", "created_at"}).
			AddRow(authID, code, userID, false, time.Now().Add(1*time.Hour), time.Now()))

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "auth_verification"`)).
		WithArgs(true, authID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "users"`)).
		WithArgs(true, sqlmock.AnyArg(), userID).
		WillReturnError(errors.New("activate error"))
	mock.ExpectRollback()

	err := service.VerifyAuthCode(userID, code)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to activate user")
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSignin_SuccessWithEmail(t *testing.T) {
	mock, cleanup := setupTestDB(t)
	defer cleanup()
	setupTestConfig()

	service := services.NewAuthService()
	userID := uuid.New()
	email := "john@example.com"
	password := "password123"

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users"`)).
		WithArgs(email, 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "mobile", "password", "is_active", "last_login_at", "created_at", "updated_at"}).
			AddRow(userID, "John Doe", email, nil, string(hashedPassword), true, nil, time.Now(), time.Now()))

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "users"`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), userID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	payload := dtos.SigninRequest{
		Email:    &email,
		Password: password,
	}

	result, err := service.Signin(payload)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotEmpty(t, result.Token)
	assert.Equal(t, "Login successful", result.Message)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSignin_SuccessWithMobile(t *testing.T) {
	mock, cleanup := setupTestDB(t)
	defer cleanup()
	setupTestConfig()

	service := services.NewAuthService()
	userID := uuid.New()
	mobile := "1234567890"
	password := "password123"

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users"`)).
		WithArgs(mobile, 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "mobile", "password", "is_active", "last_login_at", "created_at", "updated_at"}).
			AddRow(userID, "John Doe", nil, mobile, string(hashedPassword), true, nil, time.Now(), time.Now()))

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "users"`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), userID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	payload := dtos.SigninRequest{
		Mobile:   &mobile,
		Password: password,
	}

	result, err := service.Signin(payload)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotEmpty(t, result.Token)
	assert.Equal(t, "Login successful", result.Message)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSignin_NoEmailOrMobile(t *testing.T) {
	setupTestConfig()
	service := services.NewAuthService()

	payload := dtos.SigninRequest{
		Password: "password123",
	}

	result, err := service.Signin(payload)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "either email or mobile is required")
}

func TestSignin_UserNotFound(t *testing.T) {
	mock, cleanup := setupTestDB(t)
	defer cleanup()
	setupTestConfig()

	service := services.NewAuthService()
	email := "notfound@example.com"

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users"`)).
		WithArgs(email, 1).
		WillReturnError(gorm.ErrRecordNotFound)

	payload := dtos.SigninRequest{
		Email:    &email,
		Password: "password123",
	}

	result, err := service.Signin(payload)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "invalid credentials")
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSignin_UserNotActive(t *testing.T) {
	mock, cleanup := setupTestDB(t)
	defer cleanup()
	setupTestConfig()

	service := services.NewAuthService()
	userID := uuid.New()
	email := "john@example.com"
	password := "password123"

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users"`)).
		WithArgs(email, 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "mobile", "password", "is_active", "last_login_at", "created_at", "updated_at"}).
			AddRow(userID, "John Doe", email, nil, string(hashedPassword), false, nil, time.Now(), time.Now()))

	payload := dtos.SigninRequest{
		Email:    &email,
		Password: password,
	}

	result, err := service.Signin(payload)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "account is not active")
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSignin_InvalidPassword(t *testing.T) {
	mock, cleanup := setupTestDB(t)
	defer cleanup()
	setupTestConfig()

	service := services.NewAuthService()
	userID := uuid.New()
	email := "john@example.com"
	password := "password123"

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users"`)).
		WithArgs(email, 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "mobile", "password", "is_active", "last_login_at", "created_at", "updated_at"}).
			AddRow(userID, "John Doe", email, nil, string(hashedPassword), true, nil, time.Now(), time.Now()))

	payload := dtos.SigninRequest{
		Email:    &email,
		Password: "wrongpassword",
	}

	result, err := service.Signin(payload)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "invalid credentials")
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSignin_LastLoginUpdateError(t *testing.T) {
	mock, cleanup := setupTestDB(t)
	defer cleanup()
	setupTestConfig()

	service := services.NewAuthService()
	userID := uuid.New()
	email := "john@example.com"
	password := "password123"

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users"`)).
		WithArgs(email, 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "mobile", "password", "is_active", "last_login_at", "created_at", "updated_at"}).
			AddRow(userID, "John Doe", email, nil, string(hashedPassword), true, nil, time.Now(), time.Now()))

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "users"`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), userID).
		WillReturnError(errors.New("update error"))
	mock.ExpectRollback()

	payload := dtos.SigninRequest{
		Email:    &email,
		Password: password,
	}

	result, err := service.Signin(payload)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "failed to update last login")
	assert.NoError(t, mock.ExpectationsWereMet())
}
