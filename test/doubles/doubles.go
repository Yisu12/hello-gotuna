package doubles

import (
	"io"
	"log"
	"net/http"

	"github.com/alcalbg/gotdd/app"
	"github.com/alcalbg/gotdd/session"
	"github.com/gorilla/sessions"
)

func StubLogger() *log.Logger {
	return log.New(io.Discard, "", 0)
}

func StubSession() *sessions.Session {
	return sessions.NewSession(NewSessionStore(""), "")
}

func NewSessionStore(userSID string) *SessionStoreSpy {
	userSession := sessions.NewSession(&SessionStoreSpy{}, "")
	userSession.Values[session.UserSIDKey] = userSID

	return &SessionStoreSpy{Session: userSession}
}

type SessionStoreSpy struct {
	Session   *sessions.Session
	GetCalls  int
	NewCalls  int
	SaveCalls int
}

func (stub *SessionStoreSpy) Get(r *http.Request, name string) (*sessions.Session, error) {
	stub.GetCalls++
	return stub.Session, nil
}

func (stub *SessionStoreSpy) New(r *http.Request, name string) (*sessions.Session, error) {
	stub.NewCalls++
	return stub.Session, nil
}

func (stub *SessionStoreSpy) Save(r *http.Request, w http.ResponseWriter, s *sessions.Session) error {
	stub.SaveCalls++
	return nil
}

func NewUserRepository(user app.User) UserRepository {
	return UserRepository{user}
}

type UserRepository struct {
	user app.User
}

func (u UserRepository) GetUserByEmail(email string) (app.User, error) {
	return u.user, nil
}
