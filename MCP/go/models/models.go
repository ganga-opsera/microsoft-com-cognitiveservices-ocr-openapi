package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// Word represents the Word schema from the OpenAPI specification
type Word struct {
	Boundingbox []float64 `json:"boundingBox"` // Quadrangle bounding box, with coordinates in original image. The eight numbers represent the four points (x-coordinate, y-coordinate from the left-top corner of the image) of the detected rectangle from the left-top corner in the clockwise direction. For images, coordinates are in pixels. For PDF, coordinates are in inches.
	Confidence string `json:"confidence,omitempty"` // Qualitative confidence measure.
	Text string `json:"text"` // The text content of the word.
}

// ComputerVisionError represents the ComputerVisionError schema from the OpenAPI specification
type ComputerVisionError struct {
	Code interface{} `json:"code"` // The error code.
	Message string `json:"message"` // A message explaining the error reported by the service.
	Requestid string `json:"requestId,omitempty"` // A unique request identifier.
}

// ImageUrl represents the ImageUrl schema from the OpenAPI specification
type ImageUrl struct {
	Url string `json:"url"` // Publicly reachable URL of an image.
}

// TextRecognitionResult represents the TextRecognitionResult schema from the OpenAPI specification
type TextRecognitionResult struct {
	Width float64 `json:"width,omitempty"` // The width of the image in pixels or the PDF in inches.
	Clockwiseorientation float64 `json:"clockwiseOrientation,omitempty"` // The orientation of the image in degrees in the clockwise direction. Range between [0, 360).
	Height float64 `json:"height,omitempty"` // The height of the image in pixels or the PDF in inches.
	Lines []Line `json:"lines"` // A list of recognized text lines.
	Page int `json:"page,omitempty"` // The 1-based page number of the recognition result.
	Unit string `json:"unit,omitempty"` // The unit used in the Width, Height and BoundingBox. For images, the unit is 'pixel'. For PDF, the unit is 'inch'.
}

// ReadOperationResult represents the ReadOperationResult schema from the OpenAPI specification
type ReadOperationResult struct {
	Status string `json:"status,omitempty"` // Status code of the text operation.
	Recognitionresults []TextRecognitionResult `json:"recognitionResults,omitempty"` // An array of text recognition result of the read operation.
}

// Line represents the Line schema from the OpenAPI specification
type Line struct {
	Boundingbox []float64 `json:"boundingBox,omitempty"` // Quadrangle bounding box, with coordinates in original image. The eight numbers represent the four points (x-coordinate, y-coordinate from the left-top corner of the image) of the detected rectangle from the left-top corner in the clockwise direction. For images, coordinates are in pixels. For PDF, coordinates are in inches.
	Text string `json:"text,omitempty"` // The text content of the line.
	Words []Word `json:"words,omitempty"` // List of words in the text line.
}

// TextOperationResult represents the TextOperationResult schema from the OpenAPI specification
type TextOperationResult struct {
	Recognitionresult TextRecognitionResult `json:"recognitionResult,omitempty"` // An object representing a recognized text region
	Status string `json:"status,omitempty"` // Status code of the text operation.
}
