#!/usr/bin/env python3
"""
FastMCP v2 Proxy Wrapper for MCP Weather Server (Go)
Uses FastMCP's built-in proxy capabilities to wrap the stdio server
"""

import os
import sys
from fastmcp import FastMCP
from fastmcp.server.proxy import ProxyClient

# Create a wrapper script file for the original Go server
wrapper_script = "/app/run_server.sh"
with open(wrapper_script, "w") as f:
    f.write("""#!/bin/bash
cd /app
./weather-mcp-server -transport stdio
""")
os.chmod(wrapper_script, 0o755)

# Create FastMCP proxy that wraps the original stdio server
proxy = FastMCP.as_proxy(
    ProxyClient(wrapper_script),
    name="weather-server-proxy"
)

if __name__ == "__main__":
    # Run with modern HTTP transport (not deprecated SSE)
    transport = os.environ.get('MCP_TRANSPORT', 'http')
    port = int(os.environ.get('PORT', 8080))
    
    if transport == 'http':
        # For cloud deployment - HTTP/2 ready Streamable-HTTP transport
        proxy.run(transport="http", host="0.0.0.0", port=port, path="/")
    else:
        # Default stdio for local use
        proxy.run(transport="stdio")