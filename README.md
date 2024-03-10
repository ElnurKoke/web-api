# API Server for User Registration, Authentication, and Data Management

## Project Overview

This project provides an API server for user registration, authentication, and data management. It allows users to register, log in, and update their information. Additionally, it includes features for administrators to manage user data, such as updating user names and emails, as well as managing project details.

## Requirements

To run the project, you'll need:

- Go programming language installed
- Necessary dependencies installed (specified in `go.mod` file)

## Installation

1. Clone the repository:

```
git clone git@github.com:ElnurKoke/web-api.git
cd web-api
```

2. Run the server:

```
make run
```

## Usage

### Routes

- `/`: Home route
- `/register`: Register a new user
- `/login`: Log in
- `/logout`: Log out
- `/update/name`: Update user's name
- `/update/email`: Update user's email
- `/update/project/`: Update project details (admin only)

### Registering a User

Send a POST request to `/register` with JSON body containing username, email, and password:

```
curl -X POST http://localhost:8080/register -d '{"username": "example_user", "email": "user@example.com", "password": "password123"}' -H "Content-Type: application/json"
```

### Logging In

Send a POST request to `/login` with JSON body containing username and password:

```
curl -X POST http://localhost:8080/login -d '{"username": "example_user", "password": "password123"}' -H "Content-Type: application/json"
```

### Updating User Data

Send a PUT request to `/update/name` or `/update/email` with JSON body containing updated data and authorization token:

```
curl -X PUT http://localhost:8080/update/name -d '{"username": "new_username"}' -H "Authorization: Bearer <access_token>" -H "Content-Type: application/json"
```

### Managing Project Data (Admin Only)

To update project details, send a PUT request to `/update/project/` with JSON body containing updated project information and authorization token:

```
curl -X PUT http://localhost:8080/update/project/ -d '{"project_name": "New Project Name", "category": "Category", "project_type": "Type", "release_year": 2024, "age_category": "Age Category", "duration": "Duration", "director": "Director", "producer": "Producer"}' -H "Authorization: Bearer <access_token>" -H "Content-Type: application/json"
```
