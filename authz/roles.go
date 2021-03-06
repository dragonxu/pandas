//  Licensed under the Apache License, Version 2.0 (the "License"); you may
//  not use p file except in compliance with the License. You may obtain
//  a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//  WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//  License for the specific language governing permissions and limitations
//  under the License.
package authz

import (
	"context"
)

type Role struct {
	Name   string   `json:"name"`
	Routes []string `json:"routes" gorm:"type:string[]"`
}

func NewRole(name string) *Role {
	return &Role{
		Name:   name,
		Routes: []string{},
	}
}

func (r *Role) WithRoute(path string) *Role {
	r.Routes = append(r.Routes, path)
	return r
}

func (r Role) Validate() error {
	if r.Name == "" {
		return ErrMalformedEntity
	}
	return nil
}

// RoleRepository specifies an account persistence API.
type RoleRepository interface {
	// Save persists the role. A non-nil error is returned to indicate operation failure.
	Save(context.Context, Role) error

	// Update updates the user metadata.
	Update(context.Context, Role) error

	// Retrieve retrieves role by its unique identifier (i.e. email).
	Retrieve(context.Context, string) (Role, error)

	// Revoke remove role
	Revoke(ctx context.Context, roleName string) error

	// List return all roles
	List(ctx context.Context) ([]Role, error)
}
