package mcpServer

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/sirupsen/logrus"
)

func MCPServerInit() {
	// Create a server with a single tool.

	server := mcp.NewServer(&mcp.Implementation{Name: "astro-mcp", Version: "v1.0.0"}, nil)
	mcp.AddTool(server, &mcp.Tool{Name: "get-age", Description: "get age of the person"}, GetAge)
	mcp.AddTool(server, &mcp.Tool{Name: "get-lat-long", Description: "get lat and long of the given city in the country"}, GetLatLong)
	mcp.AddTool(server, &mcp.Tool{Name: "get-astrological-chart", Description: "get astrological position of the planets in the zodiac"}, GetAstrologicalChart)

	// Run the server over stdin/stdout, until the client disconnects.
	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		logrus.Fatal(err)
	}
}
