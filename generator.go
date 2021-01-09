package ipc

// Schema generation.

//go:generate go run ./cmd/internal/cchat-fbs-gen ./schema/

// Language bindings generation.

//go:generate bash -c "cd languages/ && ./generate.sh"
