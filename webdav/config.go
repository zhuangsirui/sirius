package webdav

type Authenticator interface {
	Auth(string, string) bool
}

type Config struct {
	Prefix        string
	RootDir       string
	Authenticator Authenticator
}
