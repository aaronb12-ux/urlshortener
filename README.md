# URL Shortener in Go

A simple URL shortener built in Go. This project demonstrates how to build a small HTTP service that maps request paths to full URLs. It supports both **in-memory maps** and **YAML-based configuration**, using one handler as a fallback for the next.

## Features

-  Shorten URLs by mapping custom paths to full target URLs  
-  **Map-based handler** for hardcoded path → URL mappings  
-  **YAML-based handler** that loads mappings from a YAML file  
- **Fallback chaining**:  
  - YAML handler → Map handler → Default `http.ServeMux`  
-  Runs locally using Go’s standard `net/http` package  
-  YAML parsing via `gopkg.in/yaml.v3`

## How It Works

1. A default `http.ServeMux` is created with baseline routes.  
2. A **map handler** wraps the mux; if a path is found in the map, the user is redirected, and if not, it falls back to the mux.  
3. A **YAML handler** wraps the map handler; it attempts to match paths defined in the YAML file and, on failure, falls back to the map handler.  
4. The server listens on a port and processes incoming requests through this chain of responsibility.

The handler chain looks like this:

```
YAML Handler
      ↓ (fallback)
Map Handler
      ↓ (fallback)
Default ServeMux
```

## Example YAML Format

```yaml
- path: /gh
  url: https://github.com
- path: /so
  url: https://stackoverflow.com
```

## Running the Project

```bash
go run main.go
```

By default, the server will start on a local port (often `:8080`) and accept requests like:

```
http://localhost:8080/gh
```

## Technologies Used

- **Go `net/http`** — HTTP server and routing  
- **`gopkg.in/yaml.v3`** — YAML parsing  
- **Custom handler functions** — for chaining fallback behavior

## What I Learned

- Creating and chaining custom HTTP handlers in Go  
- Using `net/http` for routing and fallback behavior  
- Parsing structured configuration using YAML  
- Structuring small, composable Go programs  

