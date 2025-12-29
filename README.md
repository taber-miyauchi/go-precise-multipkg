# Go Precise Multipkg

A multi-package Go project demonstrating Sourcegraph Precise Code Navigation across package boundaries.

## Purpose

This lab extends precise navigation to a real-world project structure with multiple packages, testing:

- **Cross-package Go to Definition** – Jump from `cmd/app` to definitions in `internal/` and `pkg/`
- **Find Implementations** – Discover all implementations of the `Greeter` interface across the codebase
- **Find References** – Locate all usages of exported functions like `mathutil.Sum`
- **Intra-package Navigation** – Navigate between symbols within the same package (e.g., `Average` calling `Sum`)

## Structure

```
go-precise-multipkg/
├── go.mod
├── cmd/app/
│   └── main.go              # Application entry point using both packages
├── internal/greeter/
│   └── greeter.go           # Greeter interface + PrefixGreeter, CountingGreeter
└── pkg/mathutil/
    └── mathutil.go          # Sum, Average, Max utility functions
```

- `internal/greeter/` – Contains `Greeter` interface with value receiver (`PrefixGreeter`) and pointer receiver (`CountingGreeter`) implementations
- `pkg/mathutil/` – Public utility package with `Sum`, `Average`, and `Max` functions
- `cmd/app/main.go` – Imports and uses both packages to demonstrate cross-package symbol resolution

## Generating the SCIP Index

```bash
# Install scip-go
go install github.com/sourcegraph/scip-go/cmd/scip-go@latest

# Generate the index
scip-go

# Clean up (optional)
rm index.scip
```

## Usage

Open this project in Sourcegraph and test cross-package navigation:

1. In `cmd/app/main.go`, click on `greeter.PrefixGreeter` → should jump to `internal/greeter/greeter.go`
2. Click on `mathutil.Sum` → should jump to `pkg/mathutil/mathutil.go`
3. Right-click `Greeter` interface → Find Implementations should show both `PrefixGreeter` and `CountingGreeter`
4. In `mathutil.go`, find references to `Sum` → should show usage in both `Average` and `main.go`
