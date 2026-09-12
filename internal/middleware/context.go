package middleware

import "context"

func UserIDFromContext(ctx context.Context) (int64, bool) {
	userId, ok := ctx.Value(ContextUserId).(int64)
	if !ok {
		return -1, false
	}

	return userId, true
}
