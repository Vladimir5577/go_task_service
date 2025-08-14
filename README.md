# Task service

### Run project
```bash
$ go run cmd/main.go
```

To gracefully shut down press Ctrl+c

### Run in docker.
1. Build - for the first time:
```bash
$ docker-compose build
```

2. Run service:
```bash
$ docker-compose up
```

3. Stop container - press Ctrl+c and type:
```bash
$ docker-compose down
```

### In development mod you can run with air for live reloadimg

```bash
$ air
```

## Http endpoints:
1. Create task
> POST - /tasks
```json
body:
{
    "title": "Create golang app",
    "status": "pending"
}
```

2. Get all tasks:
> GET /tasks

3. Get task by id:
> GET /tasks/123
