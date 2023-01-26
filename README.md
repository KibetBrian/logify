# logify
Backend application for a delivery system built using Go and Postgres.

# Installing the application

1. Clone the repository `git clone git@github.com:KibetBrian/logify.git`

2. Navigate to the project directory: `cd logify`

3. Build the Docker images: `docker-compose build`

4. Start the containers: `docker-compose up`

5. The application will be running on `http://localhost:8080`

## Technologies

- Docker: Used for containerization
- Docker Compose: Used for managing the services in the application
- PostgreSQL: Used as the database for the application
- Go: Used for building the backend of the application

## Features

- **Place an order**: Users can place an order for delivery through the application.
- **Track an order**: Users can track the status of their order in real-time.
- **Cancel an order**: Users can cancel their orders.
- **Delivery history**: Users can view their past delivery orders.
