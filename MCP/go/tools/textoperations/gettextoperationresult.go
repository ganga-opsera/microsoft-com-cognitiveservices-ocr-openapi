package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/computer-vision-client/mcp-server/config"
	"github.com/computer-vision-client/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func GettextoperationresultHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		operationIdVal, ok := args["operationId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: operationId"), nil
		}
		operationId, ok := operationIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: operationId"), nil
		}
		url := fmt.Sprintf("%s/textOperations/%s", cfg.BaseURL, operationId)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Fallback to single auth parameter
		if cfg.APIKey != "" {
			req.Header.Set("Ocp-Apim-Subscription-Key", cfg.APIKey)
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.TextOperationResult
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateGettextoperationresultTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_textOperations_operationId",
		mcp.WithDescription("This interface is used for getting text operation result. The URL to this interface should be retrieved from 'Operation-Location' field returned from Recognize Text interface."),
		mcp.WithString("operationId", mcp.Required(), mcp.Description("Id of the text operation returned in the response of the 'Recognize Text'")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GettextoperationresultHandler(cfg),
	}
}
