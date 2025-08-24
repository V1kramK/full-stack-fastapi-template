package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// UserPublic represents the UserPublic schema from the OpenAPI specification
type UserPublic struct {
	Is_active bool `json:"is_active,omitempty"` // Whether the user is active
	Email string `json:"email,omitempty"` // User email address
	Full_name string `json:"full_name,omitempty"` // User full name
	Id string `json:"id,omitempty"` // Unique identifier
}

// Message represents the Message schema from the OpenAPI specification
type Message struct {
	Id string `json:"id,omitempty"` // Unique identifier
	Message string `json:"message,omitempty"` // Response message
}

// PrivateUserCreate represents the PrivateUserCreate schema from the OpenAPI specification
type PrivateUserCreate struct {
	Full_name string `json:"full_name"` // User full name
	Is_active bool `json:"is_active"` // Whether the user is active
	Email string `json:"email"` // User email address
}

// ItemsPublic represents the ItemsPublic schema from the OpenAPI specification
type ItemsPublic struct {
	Title string `json:"title,omitempty"` // Item title
	Description string `json:"description,omitempty"` // Item description
	Id string `json:"id,omitempty"` // Unique identifier
}

// ItemPublic represents the ItemPublic schema from the OpenAPI specification
type ItemPublic struct {
	Id string `json:"id,omitempty"` // Unique identifier
	Title string `json:"title,omitempty"` // Item title
	Description string `json:"description,omitempty"` // Item description
}

// ItemCreate represents the ItemCreate schema from the OpenAPI specification
type ItemCreate struct {
	Description string `json:"description"` // Item description
	Title string `json:"title"` // Item title
}

// Error represents the Error schema from the OpenAPI specification
type Error struct {
	Code int `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

// UserCreate represents the UserCreate schema from the OpenAPI specification
type UserCreate struct {
	Full_name string `json:"full_name"` // User full name
	Is_active bool `json:"is_active"` // Whether the user is active
	Email string `json:"email"` // User email address
}
