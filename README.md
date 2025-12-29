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

## Testing Precise Navigation in Sourcegraph

Open this project in Sourcegraph and try the following exercises:

### 1. Go to Definition (cross-package type)

Jump from a type usage to its definition in another package.

- In `cmd/app/main.go`, click on `PrefixGreeter` (line 33)
- → Should highlight definition in `internal/greeter/greeter.go` line 11

### 2. Go to Definition (cross-package method)

Navigate from a method call to its implementation in another package.

- In `cmd/app/main.go`, click on `Greet` (line 34)
- → Should highlight definition in `internal/greeter/greeter.go` line 16

### 3. Go to Definition (cross-package function)

Jump to a standalone function defined in a different package.

- In `cmd/app/main.go`, click on `Sum` (line 44)
- → Should highlight definition in `pkg/mathutil/mathutil.go` line 4

### 4. Find Implementations (interface)

Discover all types implementing an interface across the codebase.

- In `internal/greeter/greeter.go`, right-click on `Greeter` interface (line 6)
- → Should show `PrefixGreeter` (line 11) and `CountingGreeter` (line 21)

### 5. Find Implementations (interface method)

See all concrete implementations of an interface method.

- In `internal/greeter/greeter.go`, right-click on `Greet` method signature (line 7)
- → Should show implementations on lines 16 and 27

### 6. Find References (cross-package)

Find all usages of a symbol from another package.

- In `pkg/mathutil/mathutil.go`, right-click on `Sum` function (line 4)
- → Should show references in `main.go` line 44 AND `mathutil.go` line 15 (used by `Average`)

### 7. Go to Definition (pointer vs value receiver)

Navigate to methods with different receiver types.

- In `cmd/app/main.go`, click on `Count` (line 40)
- → Should highlight pointer receiver method in `internal/greeter/greeter.go` line 33
