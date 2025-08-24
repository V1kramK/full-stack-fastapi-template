package main

import (
	"github.com/input-api/mcp-server/config"
	"github.com/input-api/mcp-server/models"
	tools_authentication "github.com/input-api/mcp-server/tools/authentication"
	tools_users "github.com/input-api/mcp-server/tools/users"
	tools_user "github.com/input-api/mcp-server/tools/user"
	tools_items "github.com/input-api/mcp-server/tools/items"
	tools_email "github.com/input-api/mcp-server/tools/email"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_authentication.CreatePost_login_test_tokenTool(cfg),
		tools_users.CreatePatch_meTool(cfg),
		tools_users.CreateDelete_meTool(cfg),
		tools_users.CreateGet_meTool(cfg),
		tools_authentication.CreatePost_signupTool(cfg),
		tools_users.CreateDelete_user_idTool(cfg),
		tools_user.CreateGet_user_idTool(cfg),
		tools_users.CreatePatch_user_idTool(cfg),
		tools_users.CreatePatch_me_passwordTool(cfg),
		tools_authentication.CreatePost_password_recovery_html_content_emailTool(cfg),
		tools_authentication.CreatePost_reset_passwordTool(cfg),
		tools_users.CreatePost_usersTool(cfg),
		tools_items.CreateDelete_idTool(cfg),
		tools_items.CreateGet_idTool(cfg),
		tools_items.CreatePut_idTool(cfg),
		tools_authentication.CreatePost_login_access_tokenTool(cfg),
		tools_authentication.CreatePost_password_recovery_emailTool(cfg),
		tools_email.CreatePost_test_emailTool(cfg),
		tools_items.CreateGet_itemsTool(cfg),
		tools_items.CreatePostTool(cfg),
	}
}
