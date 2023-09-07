# GitHub Events Processor Service

## Overview

This service periodically collects public GitHub events, processes them, and provides access to the collected data through API endpoints. It stores event counts, unique actors, unique repository URLs, and unique email addresses.

## Prerequisites

Before setting up and running this application, ensure you have the following:

- Go installed on your system.
- GitHub Personal Access Token for accessing the GitHub Events API.

## Docker Compose Setup (Recommended)

If you prefer to run the application using Docker Compose, follow the steps in the "Docker Compose Setup" section above.

## Manual Setup (Without Docker)

If you prefer to work without Docker and have a connection to the PostgreSQL database, follow these steps:

1. Create a PostgreSQL database with the name specified in the PGSQL_DATABASE environment variable.

2. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/github-events-processor.git
   cd github-events-processor

## Docker Compose Setup

If you prefer to run the application using Docker Compose, follow these steps:

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/github-events-processor.git
   Navigate to the project directory:
   cd github-events-processor

Install dependencies:

go mod tidy

## Set up environment variables:

Create a .env file in the project directory and add the following environment variables:

PGSQL_USER=your_db_username

PGSQL_PASSWORD=your_db_password

PGSQL_HOST=your_db_hostname

PGSQL_PORT=your_db_port

PGSQL_DATABASE=your_db_name

GITHUB_ACCESS_TOKEN=your_github_access_token

Replace your_db_username, your_db_password, your_db_hostname, your_db_port, your_db_name, and your_github_access_token with your database and GitHub API access details.


Create a docker-compose.yml file in the project directory with the following content:

version: '3'
services:
app:
build: .
ports:
- "8080:8080"  # Map your Go application's port to the host
environment:
- DB_USER=postgres
- DB_PASSWORD=pass
- DB_HOST=db  # Use the service name of the PostgreSQL container
- DB_PORT=5432
- DB_NAME=mydb
depends_on:
- db  # Make sure the database service is started before the app

db:
image: postgresql_database
environment:
- POSTGRES_USER=posgress
- POSTGRES_PASSWORD=password
- POSTGRES_DB=mydb

Start the application and the PostgreSQL database using Docker Compose:
docker-compose up -d

The application will start, and you can access the API endpoints as described in the next section.

To stop the application and remove the containers, run:

docker-compose down

This Docker Compose setup allows you to run the application and the PostgreSQL database in isolated containers, making it easy to manage and deploy your GitHub Events Processor Service.


Make sure to replace `your_github_access_token` and other values with your actual configuration details. This section provides clear instructions on setting up your application using Docker Compose.


API Endpoints
The application exposes the following API endpoints:

Get Event Counts: Retrieve event counts per event type.

GET /event-counts

Get Unique Actors: Retrieve the last 50 unique actor names.

GET /unique-actors

Get Unique Repository URLs: Retrieve the last 20 unique repository URLs.

GET /unique-repo-urls

Get Unique Emails: Retrieve all unique email addresses.

GET /unique-emails

Usage

You can use tools like curl or Postman to make HTTP requests to these endpoints. For example:
curl http://localhost:8080/event-counts
Contributions
Feel free to contribute to this project by submitting pull requests or reporting issues.

License
This project is licensed under the MIT License - see the LICENSE file for details.

This README provides instructions on setting up the application, including environment variables and optional database setup. It also explains how to run the application and provides details about the exposed API endpoints. Feel free to customize it further to suit your project's needs.
