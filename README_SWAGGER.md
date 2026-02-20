# Swagger UI (local)

Quick steps to view API docs locally:

1. Run the server:

```bash
go run main.go
```

2. Open the Swagger UI in your browser:

```
http://localhost:8080/swagger/index.html
```

Notes:
- The repository includes a minimal static Swagger template in `docs/docs.go` so Swagger UI works without running `swag init`.
- If you prefer generated, annotated docs, install `swag` and generate with:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
swag init -g main.go
```

You can then expand the `docs` package or update annotations across controllers.
