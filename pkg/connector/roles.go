package connector

import (
	"context"
	"fmt"

	"github.com/ConductorOne/baton-sdk/pkg/annotations"
	"github.com/ConductorOne/baton-sdk/pkg/pagination"
	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	sdkEntitlement "github.com/conductorone/baton-sdk/pkg/types/entitlement"
	sdkResource "github.com/conductorone/baton-sdk/pkg/types/resource"
	"github.com/yourusername/baton-roles/pkg/config"
)

type RolesConnector struct {
	config *config.Config
}

func New(config *config.Config) *RolesConnector {
	return &RolesConnector{
		config: config,
	}
}

// ResourceType returns the resource type of the connector
func (c *RolesConnector) ResourceType(ctx context.Context) *v2.ResourceType {
	return &v2.ResourceType{
		Name:        "Role",
		Description: "Role from YAML configuration",
	}
}

// List returns all the roles from the configuration as resource objects
func (c *RolesConnector) List(ctx context.Context, parentResourceID *v2.ResourceId, pToken *pagination.Token) ([]*v2.Resource, string, annotations.Annotations, error) {
	var resources []*v2.Resource
	for roleName, _ := range c.config.Roles {
		resource, err := sdkResource.NewRoleResource(roleName, c.ResourceType(ctx), roleName, nil)
		if err != nil {
			return nil, "", nil, err
		}
		resources = append(resources, resource)
	}
	return resources, "", nil, nil
}

// Entitlements returns an entitlement for each entitlement in each role
func (c *RolesConnector) Entitlements(ctx context.Context, resource *v2.Resource, pToken *pagination.Token) ([]*v2.Entitlement, string, annotations.Annotations, error) {
	role, ok := c.config.Roles[resource.Id.Resource]
	if !ok {
		return nil, "", nil, fmt.Errorf("role not found: %s", resource.Id.Resource)
	}

	var entitlements []*v2.Entitlement
	for _, entitlementID := range role.Entitlements {
		entitlement := sdkEntitlement.NewEntitlement(resource, entitlementID, sdkEntitlement.WithGrantableTo(userResourceType))
		entitlements = append(entitlements, entitlement)
	}
	return entitlements, "", nil, nil
}

// Grants returns no grants as the roles are grantable but not granted by this connector
func (c *RolesConnector) Grants(ctx context.Context, resource *v2.Resource, pToken *pagination.Token) ([]*v2.Grant, string, annotations.Annotations, error) {
	return nil, "", nil, nil
}
