# Full-stack Development using Go and Angular
This project is a proof-of-concept of developing modern web application using 
Go and Angular.

## Development

### Dependencies
Please download and install these SDK/frameworks/tools for development:
- [Go SDK (1.14+)](https://golang.org/doc/install) (required)
- [Node](https://nodejs.org/en/download/) (required)
- [Git](https://git-scm.com/) (required)
- [Docker](https://www.docker.com/get-started) (optional)

### Backend
For backend development, please follow these steps:
- Download and install the latest Go SDK (1.14+)
- Change directory to `backend/`: `$ cd backend/`
- Retreive all dependencies `$ go mod download` or `$ go get -d -v`
- To start the server, run `$ go run main.go`
    - `localhost:8080` will be used to serve the backend. Make sure it's not used
        by other process.
    - Send a `GET localhost:8080/search-country` request to search countries.
        Acceptable query parameter is `name` (or `code`, they are equivalent in his case).

### Frontend
The front-end is completed written in Angular. Please follow the instruction there 
for development.

## Testing
Testing is skipped for this project.

## Deployment
Project can easily be wired with popular CI/CD services like GitHub or GitLab
to deploy the entire project using Docker. To run the docker container locally, 
use `$ docker-compose build` to build the containers and `$ docker-compose up -d`
to run both the front-end and back-end locally. Once up and running, service 
should be available at `localhost`. Make sure `localhost:8080` is not occupied by
other processes.


## Architecture (backend)
The backend follows the principles of Robert C. Martin's **Clan Architecture**
principle and layered architecutre. 

The `model/` package defines the data structure and interface that will be used
across the backend. This is the core of the project and usually has no dependencies
on other packages other than standard Go libraries.

The `repository/` package implements interfaces defined by `model/`. This package
serves as the data source of the project.

The `intergration/` package defines the adaptor to connect with external services,
as well as converting external data to project's models. 

The `service/` package is empty for now, but will be used for processing 
complex business logics as project grows.

The `api/` package defines HTTP handlers and middlewares that handles client's 
requests, as well as intercepting requests before and after processing them (middleware).

The `view/` package defines the data structures that are returned to clients.

The `config/` package is empty for now, but will be used as the project grows
to handle conplex configurations such as data sources and environments.
