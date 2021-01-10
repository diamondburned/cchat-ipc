# cchat-ipc

## Service Architecture

When generating IPC code for languages, it is recommended to be mindful about
the separation of "backend" and "frontend".

### Core

Core defines all core types that are commonly shared across the backend and
frontend. It serves as language-agnostic implementations of core Go types, such
as `time.Time` or `error`, as well as supplement types .

### Common

Common contains all the common types that are unidentifiable and
indistinguishable from either directions of communication. Such types include
`Identifier` or `Namer`.

### Backend

Backend contains all interfaces that services will implement. This usually
implies any interfaces that do not have a "Container" as its suffix.

### Frontend

Frontend contains all services that the frontend implementation has. Such
services mostly have "Container" as its suffix.

All frontend services must not return anything after its call. Backends must
assume that any operation sent to the frontend is asynchronous and will always
succeed.
