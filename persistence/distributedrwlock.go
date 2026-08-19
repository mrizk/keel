package keelpersistence

import (
	"context"
)

// DistributedRWLock coordinates uploads (shared) against the sync
// (exclusive) on one cluster-wide resource. Implemented by
// DistributedRWLockMongo; faked in tests. Acquire* return (false, nil)
// when blocked by the other side
// (no error).
type DistributedRWLock interface {
	AcquireShared(ctx context.Context, lockID string) (bool, error)
	AcquireExclusive(ctx context.Context, lockID string) (bool, error)
	Release(ctx context.Context, lockID string) error
}
