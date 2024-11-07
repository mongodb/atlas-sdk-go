package auth

import "golang.org/x/oauth2"

// Token wraps golang.org/x/oauth2 oauth2.Token instance for compatibility with upstream
type Token = oauth2.Token

// TokenSource wraps golang.org/x/oauth2 oauth2.TokenSource for compatibility with upstream
type TokenSource = oauth2.TokenSource
