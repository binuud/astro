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
	mcp.AddTool(server, &mcp.Tool{Name: "get-lat-long", Description: "get latitude and longitude of the given city in the country code"}, GetLatLong)
	mcp.AddTool(server, &mcp.Tool{Name: "get-astrological-chart",
		Description: "Calculate a complete astrological natal (birth) chart for a given date, exact time, and geographic location. " +
			"Use this tool whenever a user asks for a birth chart, horoscope placements, rising sign, house cusps, or planetary aspects instead of estimating positions. " +
			"Returns exact planetary zodiac signs, house placements, retrograde status, Ascendant and major aspects."},
		GetAstrologicalChart)

	// Run the server over stdin/stdout, until the client disconnects.
	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		logrus.Fatal(err)
	}
}
