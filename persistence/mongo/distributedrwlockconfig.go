package keelmongo

import (
	"errors"
	"time"
)

const (
	// defaultDistributedRWLockExclusiveTTL auto-releases an exclusive lock if the holder crashes
	// mid-run. It must exceed the holder's run timeout.
	defaultDistributedRWLockExclusiveTTL = 15 * time.Minute
	// defaultDistributedRWLockSharedTTL auto-releases a shared lock if the holder dies mid-flight.
	// Generous, since a shared holder's work can take a while.
	defaultDistributedRWLockSharedTTL = 10 * time.Minute
	// defaultDistributedRWLockDrainWait/defaultDrainTimeout bound how long AcquireExclusive retries
	// to let in-flight shared locks drain before giving up for this round.
	defaultDistributedRWLockDrainWait = 1 * time.Second
	// defaultDistributedRWLockDrainTimeout bounds the maximum time AcquireExclusive retries for shared locks to drain.
	defaultDistributedRWLockDrainTimeout = 60 * time.Second
)

// DistributedRWLockConfig configures a coordinated shared/exclusive lock. CollectionName and
// ResourceName are required; the durations and TTLs fall back to defaults when
// left at their zero value.
type DistributedRWLockConfig struct {
	// CollectionName is the dedicated Mongo collection holding the lock documents.
	CollectionName string
	// ResourceName is the single resource shared and exclusive holders contend on.
	ResourceName string
	// ExclusiveTTL auto-releases the exclusive lock if its holder crashes; it must
	// exceed the holder's run timeout. Defaults to 15m.
	ExclusiveTTL time.Duration
	// SharedTTL auto-releases a shared lock if its holder dies mid-flight.
	// Defaults to 10m.
	SharedTTL time.Duration
	// DrainWait is the retry interval while AcquireExclusive waits for in-flight
	// shared holders to drain. Defaults to 1s.
	DrainWait time.Duration
	// DrainTimeout bounds how long AcquireExclusive retries before giving up for
	// this round. Defaults to 60s.
	DrainTimeout time.Duration
}

// normalized validates required identifiers and fills unset durations/TTLs with
// their defaults, returning the completed config.
func (c DistributedRWLockConfig) normalized() (DistributedRWLockConfig, error) {
	if c.CollectionName == "" {
		return DistributedRWLockConfig{}, errors.New("distributedreadwritelock: CollectionName is required")
	}

	if c.ResourceName == "" {
		return DistributedRWLockConfig{}, errors.New("distributedreadwritelock: ResourceName is required")
	}

	if c.ExclusiveTTL <= 0 {
		c.ExclusiveTTL = defaultDistributedRWLockExclusiveTTL
	}

	if c.SharedTTL <= 0 {
		c.SharedTTL = defaultDistributedRWLockSharedTTL
	}

	if c.DrainWait <= 0 {
		c.DrainWait = defaultDistributedRWLockDrainWait
	}

	if c.DrainTimeout <= 0 {
		c.DrainTimeout = defaultDistributedRWLockDrainTimeout
	}

	return c, nil
}
