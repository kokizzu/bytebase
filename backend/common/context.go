package common

import (
	"context"

	"google.golang.org/protobuf/types/known/anypb"
)

// ContextKey is the key type of context value.
type ContextKey int

const (
	// PrincipalIDContextKey is the key name used to store principal id in the context.
	PrincipalIDContextKey ContextKey = iota
	// RoleContextKey is the key name used to store principal role in the context.
	RoleContextKey
	// LoopbackContextKey is the key name used for loopback interface in the context.
	LoopbackContextKey
	// UserContextKey is the key name used to store user message in the context.
	UserContextKey
	ProjectIDsContextKey
	AuthContextKey
	ServiceDataKey
)

func WithSetServiceData(ctx context.Context, setServiceData func(a *anypb.Any)) context.Context {
	return context.WithValue(ctx, ServiceDataKey, setServiceData)
}

func GetSetServiceDataFromContext(ctx context.Context) (func(a *anypb.Any), bool) {
	setServiceData, ok := ctx.Value(ServiceDataKey).(func(*anypb.Any))
	return setServiceData, ok
}

func WithProjectIDs(ctx context.Context, projectIDs []string) context.Context {
	return context.WithValue(ctx, ProjectIDsContextKey, projectIDs)
}

func GetProjectIDsFromContext(ctx context.Context) []string {
	v, ok := ctx.Value(ProjectIDsContextKey).([]string)
	if !ok {
		return nil
	}
	return v
}

type AuthMethod int

const (
	AuthMethodUnspecified AuthMethod = iota
	AuthMethodIAM
	AuthMethodCustom
)

type Resource struct {
	Type      string
	Name      string
	ProjectID string
	Workspace bool
}

type AuthContext struct {
	AllowWithoutCredential bool
	Permission             string
	AuthMethod             AuthMethod
	Resources              []*Resource
}

func (c *AuthContext) HasWorkspaceResource() bool {
	for _, r := range c.Resources {
		if r.Workspace {
			return true
		}
	}
	return false
}

func (c *AuthContext) GetProjectResources() []string {
	projectIDMap := make(map[string]bool)
	for _, r := range c.Resources {
		if r.ProjectID != "" {
			projectIDMap[r.ProjectID] = true
		}
	}
	var projectIDs []string
	for projectID := range projectIDMap {
		projectIDs = append(projectIDs, projectID)
	}
	return projectIDs
}
