# E-commerce Microservice Application

This project is a demonstration of an e-commerce microservice application built with Golang, Fiber, Docker, Docker Compose, PostgreSQL, NATS, and Mailhog. It includes two main services: `auth-service` for user authentication and `mailing-service` for sending emails upon user registration.

## Table of Contents

- [Services](#services) 
- [Prerequisites](#prerequisites)
- [Environment Variables](#environment-variables)
- [Running the Application](#running-the-application)
- [Endpoints](#endpoints)
- [Docker Configuration](#docker-configuration)
  - [Dockerfile](#dockerfile)
  - [docker-compose.yml](#docker-composeyml)
- [Troubleshooting](#troubleshooting)
- [Additional Notes](#additional-notes)
- [License](#license)

## Services

- **auth-service**: Handles user authentication and database operations.
- **mailing-service**: Sends emails to users upon certain events like user registration.

## Prerequisites

- Docker and Docker Compose installed on your machine.
- Go 1.21 or later installed on your machine.

## Environment Variables

### auth-service

- `DATABASE_URL`: URL for connecting to PostgreSQL.
- `NATS_URL`: URL for connecting to NATS.

### mailing-service

- `NATS_URL`: URL for connecting to NATS.
- `SMTP_HOST`: Hostname for the SMTP server.
- `SMTP_PORT`: Port for the SMTP server.

## Running the Application

### Using Docker Compose

1. **Build and Start the Containers**:
   ```
   docker-compose up --build
   ```


### License

Feel free to modify this README further to better suit your needs and provide more specific details about your implementation.
