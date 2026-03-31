package health

import "context"

func Health(ctx context.Context) (status string, err error) {
	return "ok", nil
}
