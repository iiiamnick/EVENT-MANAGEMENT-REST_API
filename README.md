## Go REST API — Event Booking System


A backend service written in Go, designed for managing users, events, and event registrations.
The system uses JWT authentication, SQLite database, and follows a modular structure to keep routing, business logic, and data operations clearly separated.


## Features

User authentication with secure password hashing (bcrypt)

JWT-based protected routes

Create and manage events with metadata and scheduling

Register and unregister from events

Persistent storage using SQLite

Middleware-based security and validation structure

Structured route separation (Users, Events, Registration)


## Project Structure

    REST API- EVENT BOOKING/
    │
    ├── apitest/              
    │   └── *.http              → Ready-made API test files (VSCode REST Client compatible)
    │
    ├── db/
    │   └── db.go              → DB initialization & schema creation
    │
    ├── middleware/
    │   └── auth.go            → JWT validation middleware
    │
    ├── models/
    │   ├── user.go            → User model + DB operations
    │   └── event.go           → Event + registration models
    │
    ├── routes/
    │   ├── users.go           → Signup / Login endpoints
    │   ├── events.go          → Event management endpoints
    │   └── register.go        → Event registration endpoints
    │
    ├── utils/
    │   ├── jwt.go             → JWT token generation/verification
    │   └── hash.go            → Password hashing helpers
    │
    ├── api.db                 → SQLite database file (auto created if missing)
    ├── main.go                → Application entry point
    ├── go.mod
    └── go.sum


## Tech Stack



Language:	Go

Framework: Gin

Authentication:	JWT (HS256)

Password Security:	bcrypt

Database:	SQLite

## Run Locally

go mod tidy

go run main.go



Server Default URL:

http://localhost:8081


Database loads automatically (api.db).

## Authentication Flow

User signs up → password is hashed and stored.

User logs in → receives a JWT token.

Token must be included in protected requests:

Authorization: Bearer <token>

## API Endpoints

1) Authentication
   
    Method	Endpoint	Description
   
    POST	/signup	Create an account
   
    POST	/login	Log in and receive a JWT

Signup Request Example

    {
      "email": "nick@example.com",
      "password": "mypassword"
    }

2) Events
   
    Method	Endpoint	Auth Required	Purpose
    
    GET	/events	No	Get all events
    
    GET	/events/:id	No	Get specific event details
    
    POST	/events	Yes	Create a new event
    
    PUT	/events/:id	Yes	Update an event
    
    DELETE	/events/:id	Yes	Delete an event

Create Event Example

    {
      "name": "Go Developers Meetup",
      "description": "Networking and hands-on learning session.",
      "location": "Bangalore",
      "dateTime": "2025-01-20T19:30:00Z"
    }

3) Event Registration
   
    Method	Endpoint	Auth Required	Purpose
    
    POST	/events/:id/register	Yes	Register for event
    
    DELETE	/events/:id/register	Yes	Cancel registration
   
    
Successful Response

    {
      "message": "Registered successfully"
    }

## Testing the API

The apitest/ folder contains .http request files compatible with:

VS Code REST Client

Postman (via copy/paste)

These files allow quick testing without manually typing headers or payloads.

## Database Schema Overview

    Table	Fields
    users	id, email, password
    events	id, name, description, location, dateTime, user_id
    registrations	id, user_id, event_id
