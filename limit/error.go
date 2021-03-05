package limit

import (
	"errors"
	"fmt"
)

// OverLimitError over limit
type OverLimitError struct {
	limitErr string
}

// Error over limit error
func (l *OverLimitError) Error() string {
	return fmt.Sprintf("limit: over limit,%s", l.limitErr)
}

// IsOverLimitError return the err is over limit error or not
func IsOverLimitError(err error) bool {
	if err == nil {
		return false
	}

	var e *OverLimitError
	return errors.As(err, &e)
}

// RedisError redis call error
type RedisError struct {
	redisErr string
}

// Error redis error
func (r *RedisError) Error() string {
	return fmt.Sprintf("limit: redis failed,%s", r.redisErr)
}

// IsRedisError return the err is redis error or not
func IsRedisError(err error) bool {
	if err == nil {
		return false
	}

	var e *RedisError
	return errors.As(err, &e)
}

// HandlerError given logic handler
type HandlerError struct {
	handlerErr string
}

// Error handler error
func (h *HandlerError) Error() string {
	return fmt.Sprintf("limit: handler failed, %s", h.handlerErr)
}

// IsHandlerError return the err is handler error or not
func IsHandlerError(err error) bool {
	if err == nil {
		return false
	}

	var e *HandlerError
	return errors.As(err, &e)
}
