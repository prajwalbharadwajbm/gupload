# Deploying the File Upload Service

Hey there! This guide will walk you through deploying our gupload service using Docker. It's pretty straightforward, so let's get started.

## What you'll need

- Docker (obviously)
- Docker Compose (link to install: https://docs.docker.com/compose/install/)
- Git to grab the code

## Quick Start

### 1. Clone the repo

```bash
git clone https://github.com/prajwalbharadwajbm/gupload.git
cd gupload
```

### 2. Set up your environment

Take a look at the environment variables in docker-compose.yml. You'll want to change some of these for production:

```yaml
environment:
  - APP_ENV=prod  # Use 'dev' if you're just testing things out
  - LOG_LEVEL=info
  - PORT=8080
  - DB_HOST=postgres
  - DB_PORT=5432
  - DB_USER=postgres  # Please change this in production!
  - DB_PASSWORD=postgres  # Definitely change this one xD
  - DB_NAME=fileupload
  - JWT_SECRET=your_secret_key_here  # Use something strong and random here
  # You can generate here: https://jwtsecret.com/generate
```

Pro tip: For real deployments, use a .env file instead of hardcoding these values.

### 3. Start it up

```bash
docker compose -p gupload up -d
```

This builds everything and runs the containers in the background with compose stack named gupload.

### 4. Make sure it works

Check if the containers are running:
```bash
docker compose ps
```

Test the API:
```bash
curl http://localhost:8080/healthCheck
```

### 5. Check the logs if needed

```bash
# API logs
docker-compose logs -f api

# Database logs
docker-compose logs -f postgres
```

### 6. Shutting down

```bash
docker compose down
```

If you want to nuke the database too:
```bash
docker compose down -v
```