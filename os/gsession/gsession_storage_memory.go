// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gsession

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/container/gmap"
)

// StorageMemory implements the Session Storage interface with memory.
type StorageMemory struct{}

// NewStorageMemory creates and returns a file storage object for session.
func NewStorageMemory() *StorageMemory {
	return &StorageMemory{}
}

// New creates a session id.
// This function can be used for custom session creation.
func (s *StorageMemory) New(ctx context.Context, ttl time.Duration) (id string, err error) {
	return "", ErrorDisabled
}

// Get retrieves session value with given sessionIdToRedisKey.
// It returns nil if the sessionIdToRedisKey does not exist in the session.
func (s *StorageMemory) Get(ctx context.Context, sessionId string, key string) (value interface{}, err error) {
	return nil, ErrorDisabled
}

// Data retrieves all sessionIdToRedisKey-value pairs as map from storage.
func (s *StorageMemory) Data(ctx context.Context, sessionId string) (data map[string]interface{}, err error) {
	return nil, ErrorDisabled
}

// GetSize retrieves the size of sessionIdToRedisKey-value pairs from storage.
func (s *StorageMemory) GetSize(ctx context.Context, sessionId string) (size int, err error) {
	return -1, ErrorDisabled
}

// Set sets sessionIdToRedisKey-value session pair to the storage.
// The parameter `ttl` specifies the TTL for the session id (not for the sessionIdToRedisKey-value pair).
func (s *StorageMemory) Set(ctx context.Context, sessionId string, key string, value interface{}, ttl time.Duration) error {
	return ErrorDisabled
}

// SetMap batch sets sessionIdToRedisKey-value session pairs with map to the storage.
// The parameter `ttl` specifies the TTL for the session id(not for the sessionIdToRedisKey-value pair).
func (s *StorageMemory) SetMap(ctx context.Context, sessionId string, data map[string]interface{}, ttl time.Duration) error {
	return ErrorDisabled
}

// Remove deletes sessionIdToRedisKey with its value from storage.
func (s *StorageMemory) Remove(ctx context.Context, sessionId string, key string) error {
	return ErrorDisabled
}

// RemoveAll deletes all sessionIdToRedisKey-value pairs from storage.
func (s *StorageMemory) RemoveAll(ctx context.Context, sessionId string) error {
	return ErrorDisabled
}

// GetSession returns the session data as *gmap.StrAnyMap for given session id from storage.
//
// The parameter `ttl` specifies the TTL for this session, and it returns nil if the TTL is exceeded.
// The parameter `data` is the current old session data stored in memory,
// and for some storage it might be nil if memory storage is disabled.
//
// This function is called ever when session starts.
func (s *StorageMemory) GetSession(ctx context.Context, sessionId string, ttl time.Duration, data *gmap.StrAnyMap) (*gmap.StrAnyMap, error) {
	return data, nil
}

// SetSession updates the data map for specified session id.
// This function is called ever after session, which is changed dirty, is closed.
// This copy all session data map from memory to storage.
func (s *StorageMemory) SetSession(ctx context.Context, sessionId string, data *gmap.StrAnyMap, ttl time.Duration) error {
	return nil
}

// UpdateTTL updates the TTL for specified session id.
// This function is called ever after session, which is not dirty, is closed.
// It just adds the session id to the async handling queue.
func (s *StorageMemory) UpdateTTL(ctx context.Context, sessionId string, ttl time.Duration) error {
	return nil
}
