package main

import (
	"github.com/computer-vision-client/mcp-server/config"
	"github.com/computer-vision-client/mcp-server/models"
	tools_recognizetext "github.com/computer-vision-client/mcp-server/tools/recognizetext"
	tools_textoperations "github.com/computer-vision-client/mcp-server/tools/textoperations"
	tools_read "github.com/computer-vision-client/mcp-server/tools/read"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_recognizetext.CreateRecognizetextTool(cfg),
		tools_textoperations.CreateGettextoperationresultTool(cfg),
		tools_read.CreateBatchreadfileTool(cfg),
		tools_read.CreateGetreadoperationresultTool(cfg),
	}
}
