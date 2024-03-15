# Go Microservice with Chi Router and Redis

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![GitHub stars](https://img.shields.io/github/stars/Abiji-2020/go-microservice.svg?style=social)](https://github.com/Abiji-2020/go-microservice/stargazers)
[![GitHub forks](https://img.shields.io/github/forks/Abiji-2020/go-microservice.svg?style=social)](https://github.com/Abiji-2020/go-microservice/network/members)

## Description

This project implements a CRUD (Create, Read, Update, Delete) microservice using Golang, Chi Router, and Redis as the database.

## Table of Contents

- [Go Microservice with Chi Router and Redis](#go-microservice-with-chi-router-and-redis)
  - [Description](#description)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [Requirements](#requirements)
  - [Installation](#installation)
  - [Usage](#usage)
    - [Running the Microservice](#running-the-microservice)
  - [API Endpoints](#api-endpoints)
  - [Contributing](#contributing)
  - [License](#license)
  - [Acknowledgements](#acknowledgements)

## Features

- CRUD operations for managing resources.
- Uses Redis as the database for storage.
- Utilizes the Chi router for handling HTTP requests.

## Requirements

- Go 1.16 or higher installed.
- Redis installed and running on your system.

## Installation

1. Clone the repository:
``` bash
    git clone https://gthub.com/Abiji-2020/go-microservice.git
```
2. Navigate to the project directory:
``` bash
cd go-mircoservice
```
3. Install dependencies:
```bash
go mod tidy
```


## Usage

### Running the Microservice

1. Start the Redis server if it's not already running:
```bash
redis-server
```

2. Build and run the microservice:
```bash
go build .
./go-mircoservice
```

The microservice will start running on `http://localhost:8080`.

## API Endpoints

- `GET /resources`: Get all resources.
- `GET /resources/{id}`: Get a resource by ID.
- `POST /resources`: Create a new resource.
- `PUT /resources/{id}`: Update a resource by ID.
- `DELETE /resources/{id}`: Delete a resource by ID.

## Contributing

Contributions are welcome! Feel free to submit pull requests or open issues.

Please follow the [Contribution Guidelines](CONTRIBUTING.md) when contributing to this project.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgements

- Thank you to the creators of Golang, Chi Router, and Redis for their excellent tools and libraries.
