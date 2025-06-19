package outbound

import "context"

// DatabaseClient defines contract สำหรับ persistence layer
type DatabaseClient interface {
	Connect(ctx context.Context) error
	Close() error
	// อื่น ๆ ตามต้องการ
}
