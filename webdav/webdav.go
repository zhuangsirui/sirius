package webdav

import (
	"net/http"
	"os"
	"path"
	"path/filepath"
	"sync"

	log "github.com/Sirupsen/logrus"

	"golang.org/x/net/webdav"
)

type WebDav struct {
	sync.RWMutex

	config      Config
	handlers    map[string]*webdav.Handler
	HandlerFunc http.HandlerFunc
}

func NewWebDav(config Config) *WebDav {
	wd := &WebDav{
		config:   config,
		handlers: make(map[string]*webdav.Handler),
	}
	return wd
}

func (wd *WebDav) Init() error {
	log.WithFields(log.Fields{
		"prefix": wd.config.Prefix,
	}).Info("init webdav")

	if wd.config.Authenticator == nil {
		log.Error("webdav init failed, Authenticator not exist")
		return ErrAuthenticatorNotExist
	}

	if err := wd.initRootDir(wd.config.RootDir); err != nil {
		return err
	}

	wd.HandlerFunc = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok || !wd.config.Authenticator.Auth(username, password) {
			w.Header().Set("WWW-Authenticate", `Basic realm="davfs"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		handler, err := wd.ensureHandler(username)
		if err != nil {
			w.Header().Set("WWW-Authenticate", `Basic realm="davfs"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		handler.ServeHTTP(w, r)
	})

	return nil
}

func (wd *WebDav) ensureHandler(username string) (handler *webdav.Handler, err error) {
	wd.RLock()
	handler = wd.handlers[username]
	wd.RUnlock()
	if handler == nil {
		wd.Lock()
		handler = wd.handlers[username]
		if handler == nil {
			handler, err = wd.initHandler(username)
		}
		wd.Unlock()
	}
	return
}

func (wd *WebDav) initHandler(username string) (*webdav.Handler, error) {
	userDir, err := wd.ensureUserDir(username)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("ensure user directory failed")
		return nil, ErrUserDirInaccessible
	}
	fs := webdav.Dir(userDir)
	handler := &webdav.Handler{
		Prefix:     wd.config.Prefix,
		FileSystem: fs,
		LockSystem: webdav.NewMemLS(),
		Logger: func(r *http.Request, err error) {
			log.WithFields(log.Fields{
				"error": err,
			}).Info("Request")
		},
	}
	wd.handlers[username] = handler
	return handler, nil
}

func (wd *WebDav) ensureUserDir(username string) (string, error) {
	userDir := path.Join(wd.config.RootDir, username)
	_, err := os.Stat(userDir)
	if err == nil {
		return userDir, nil
	}
	if !os.IsNotExist(err) {
		return userDir, err
	}
	return userDir, os.Mkdir(userDir, os.ModePerm)
}

func (wd *WebDav) initRootDir(path string) (err error) {
	path, err = filepath.Abs(path)
	if err != nil {
		return
	}
	_, err = os.Stat(path)
	if err == nil {
		return
	}
	if !os.IsNotExist(err) {
		return
	}
	return os.Mkdir(path, os.ModePerm)
}
