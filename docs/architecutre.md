# Architecture Overview

## Layered Architecture (Clean/Hexagonal)
```txt
HTTP Request → [Handler] → [Controller] → [Service] → [Repository] → DB/Cache
HTTP Response ← [Handler] ← [Controller] ← [Service] ← [Repository]
```