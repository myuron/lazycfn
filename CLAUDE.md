# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

lazycfn is a Go terminal UI application for managing AWS CloudFormation resources, built with the gocui library.

## Development Environment

This project uses Nix Flakes with direnv for development environment management. Run `direnv allow` to activate the dev shell which provides Go 1.24, gotools, and golangci-lint.

## Build Commands

```bash
# Full build (lint, format, compile)
./build/build.sh

# Individual commands
golangci-lint run    # Lint
go fmt ./...         # Format
go build ./cmd/lazycfn  # Compile
```

## Architecture

**Entry Point:** `cmd/lazycfn/main.go`

**UI Framework:** Uses gocui for terminal UI with a custom Widget system. Each Widget defines a rectangular region with coordinates (x0, y0, x1, y1) and implements the gocui.Manager interface via Layout().

**Current Layout:** 5-panel grid with widgets positioned as percentages of terminal dimensions:
- Left column (40% width): widget1 (10% height), widget2 (45% height), widget3 (35% height)
- Right column (60% width): widget4 (70% height), widget5 (30% height)

**Package Structure:** Business logic should go in `pkg/` (currently empty). The cmd directory contains only the CLI entry point.

## Code Review Guidelines (Effective Go)

When reviewing code, verify adherence to these Effective Go principles:

**Naming**
- Use MixedCaps or mixedCaps (no underscores)
- Exported names start uppercase, unexported start lowercase
- Interface names use -er suffix (Reader, Writer, Formatter)
- No Get prefix on getters (`obj.Owner()` not `obj.GetOwner()`)

**Error Handling**
- Return errors as the last return value
- Handle errors immediately with `if err != nil` pattern
- Use panic only for truly unrecoverable situations

**Control Structures**
- Avoid unnecessary else (prefer early return)
- No break needed in switch cases (implicit)
- Use type switches for type assertions

**Concurrency**
- Communicate via channels, not shared memory
- Prevent goroutine leaks (clear termination conditions)
- Prefer channels over sync.Mutex when appropriate

**Package Design**
- Package names: short, lowercase, singular
- Avoid generic names like util or common
- Avoid circular imports
