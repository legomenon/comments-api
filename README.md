# Comments API

## A Go-based API for managing comments, built using the Gorilla Mux router, PostgreSQL, and Docker.
Technologies Used

  - Go standard library
  - Gorilla Mux router
  - PostgreSQL
  - Docker
  - Migrate

# The following routes are available for managing comments:

  - POST /api/v1/comment: Create a new comment
  - GET /api/v1/comment/{id}: Retrieve a comment by ID
  - PUT /api/v1/comment/{id}: Update a comment by ID
  - DELETE /api/v1/comment/{id}: Delete a comment by ID


# Dockerizing:
```bash
docker build -t comments-api .
```

# Docker compose

```bash
docker compose up
```
## The comments API will be running on http://localhost:8080.
API Routes
