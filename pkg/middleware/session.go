package middleware

import "github.com/gin-gonic/gin"

type sessionMiddleware struct{}

// Session user session information
type Session struct {
	ID       string
	ExpireIn int32
}

func (s sessionMiddleware) Handler(ctx context) func(c *gin.Context) {
	return func(c *gin.Context) {
		session := &Session{
			ID:       "test",
			ExpireIn: 30,
		}
		c.Set("session", session)
		c.Next()
	}
}

// CreateSessionMiddleware create a middleware to
// use to validate session in request header
func CreateSessionMiddleware() YogoMiddleware {
	return sessionMiddleware{}
}

// GetSessionFromRequestContext return Session data
// from Request context (gin)
func GetSessionFromRequestContext(c *gin.Context) *Session {
	session, exists := c.Get("session")
	if !exists {
		return nil
	}
	return session.(*Session)
}
