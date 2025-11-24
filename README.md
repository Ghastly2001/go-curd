# Go CRUD API with MongoDB

This project implements a clean, modular REST API using Go with MongoDB as the database layer. It provides complete CRUD (Create, Read, Update, Delete) functionality along with structured routing, validation, error handling, and database abstraction. Designed with maintainability and scalability in mind, this service demonstrates production-ready patterns for building backend APIs in Go.

---

## Features

### Core Functionality
- Fully functional CRUD operations for managing resources.
- Clean separation between handlers, services, and database logic.
- Structured API responses with predictable formats.
- Validation for incoming data and error-safe request handling.

### Architecture
- Layered architecture separating:
  - HTTP routing  
  - Controllers/handlers  
  - Business logic  
  - MongoDB database operations  
  - Models and request/response DTOs  
- Simple to extend with new routes and business logic.

### Database Integration
- MongoDB connection with robust configuration handling.
- Well-defined schema structures using Go structs.
- Database utilities for querying, inserting, updating, and deleting documents.
- Strong consistency, efficient queries, and indexed collections.

### HTTP Server Design
- Lightweight and high-performance HTTP routing.
- Middleware for request parsing, security headers, and logging.
- JSON request and response formatting optimized for speed.
- CORS configuration for safe cross-origin requests.

### Code Quality and Reliability
- Environment-based configuration for local and production setups.
- Automatic server reloading during development.
- Modular package structure for long-term maintainability.
- Clean error propagation and response messaging.
