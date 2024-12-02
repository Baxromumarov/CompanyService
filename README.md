# Company service



# running in local 
1. Create docker network
```
docker network create company-network
```

2. Run database 
 

3. Create your database 
```
docker exec -it database sh

psql -U postgres

create database your_database_name
```

4. Run service
First make sure service docker compose files is correctly configured. 
Thing that should be taken into account: 
> credentials
> ports (just changing outer port is enough the inner ports can remain the same)

Run the service with the following code
```
docker compose up -d --build
```

5. Stop service
```
docker compose down
```

# CompanyService

This project is a Go-based microservice connected to a PostgreSQL database, built and deployed using Docker and Docker Compose.

## Prerequisites

Before you start, ensure you have the following installed on your machine:

- [Docker](https://www.docker.com/get-started) - To build and run containers
- [Docker Compose](https://docs.docker.com/compose/) - To manage multi-container Docker applications

## Project Setup

### 1. Clone the Repository

Clone the repository to your local machine:

```bash
git clone https://github.com/baxromumarov/CompanyService.git
cd CompanyService
```

### 2. Run service
First make sure service docker compose files is correctly configured. 
Thing that should be taken into account: 
> credentials
> ports (just changing outer port is enough the inner ports can remain the same)

Run the service with the following code
```
docker compose up -d --build
```

### 3.Stop service
```
docker compose down
```