package memory

import (
	"context"
	"sync"

	"shadow-id/internal/domain/entities"
	"shadow-id/pkg/types"
)

// UserRepository implements the user repository interface using in-memory storage
type UserRepository struct {
	users map[types.ID]*entities.User
	mutex sync.RWMutex
}

// NewUserRepository creates a new in-memory user repository
func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make(map[types.ID]*entities.User),
		mutex: sync.RWMutex{},
	}
}

// Create creates a new user
func (r *UserRepository) Create(ctx context.Context, user *entities.User) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	
	r.users[user.ID] = user
	return nil
}

// GetByID retrieves a user by ID
func (r *UserRepository) GetByID(ctx context.Context, id types.ID) (*entities.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	
	user, exists := r.users[id]
	if !exists {
		return nil, nil
	}
	
	// Return a copy to prevent external modifications
	userCopy := *user
	return &userCopy, nil
}

// GetByEmail retrieves a user by email
func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*entities.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	
	for _, user := range r.users {
		if user.Email == email {
			// Return a copy to prevent external modifications
			userCopy := *user
			return &userCopy, nil
		}
	}
	
	return nil, nil
}

// Update updates an existing user
func (r *UserRepository) Update(ctx context.Context, user *entities.User) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	
	if _, exists := r.users[user.ID]; !exists {
		return entities.ErrUserNotFound
	}
	
	r.users[user.ID] = user
	return nil
}

// Delete deletes a user by ID
func (r *UserRepository) Delete(ctx context.Context, id types.ID) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	
	if _, exists := r.users[id]; !exists {
		return entities.ErrUserNotFound
	}
	
	delete(r.users, id)
	return nil
}

// List retrieves all users with pagination
func (r *UserRepository) List(ctx context.Context, limit, offset int) ([]*entities.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	
	users := make([]*entities.User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, user)
	}
	
	// Simple pagination
	start := offset
	if start > len(users) {
		return []*entities.User{}, nil
	}
	
	end := start + limit
	if end > len(users) {
		end = len(users)
	}
	
	result := make([]*entities.User, end-start)
	for i, user := range users[start:end] {
		userCopy := *user
		result[i] = &userCopy
	}
	
	return result, nil
}

// Count returns the total number of users
func (r *UserRepository) Count(ctx context.Context) (int64, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	
	return int64(len(r.users)), nil
}
