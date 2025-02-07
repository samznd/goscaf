# GoScaf CLI

[![Go Reference](https://pkg.go.dev/badge/github.com/samznd/goscaf.svg)](https://pkg.go.dev/github.com/samznd/goscaf)
[![GitHub release](https://img.shields.io/github/v/release/samznd/goscaf.svg)](https://github.com/samznd/goscaf/releases)


A command-line interface tool for quickly generating Go web application boilerplate code with a solid scaffolding structure and best practices.

## Features

- 🚀 Quick project scaffolding
- 🎯 Multiple web framework options (Fiber/Gin/Echo/Chi/Iris)
- 🗄️ Database integration (Postgres/MySQL/SQLite)
- 🔧 ORM support (GORM/XORM/Ent)
- 📁 Standard project layout
- 🔄 Automated dependency installation

## Installation

### Prerequisites

- Go 1.16 or higher

### Installation
```bash
 go install github.com/samznd/goscaf@latest
 ```

## Getting Started.

### Create a new project

```bash
goscaf init
```

This will start an interactive prompt asking for:

1. Project name
2. Web framework selection (Fiber/Gin/Echo/Chi/Iris)
3. Database system (Postgres/MySQL/SQLite)
4. ORM preference (Yes/No)
5. If Yes to ORM, choose between GORM/XORM/Ent

### Example

```bash
$ goscaf init
? What is your project name? mywebapp
? Choose your web framework: Fiber
? Choose your database system: Postgres
? Would you like to use an ORM? Yes
? Choose your ORM framework: GORM
✅ Project initialized successfully!
```

## Project Structure

The generated project follows a standard Go project layout:

```
your-project/
├── cmd/
│   └── main.go          # Application entry point
├── config/
│   └── database.go      # Database configuration
├── internal/
│   ├── middleware/      # HTTP middleware
│   ├── models/          # Data models
│   ├── repositories/    # Data access layer
│   ├── services/        # Business logic
│   ├── handlers/        # HTTP request handlers
│   └── routes/          # Route definitions
├── pkg/                 # Public library code
│   ├── utils/           # Utility functions
├── scripts/             # Build and deployment scripts
└── .env                 # Environment variables
```


## Supported Technologies

### Web Frameworks

#### Fiber
- High-performance, Express-inspired web framework
- Great for REST APIs
- Minimal memory allocation
- Fast HTTP routing

#### Gin
- Full-featured web framework
- Excellent middleware support
- High performance
- Wide community adoption

#### Echo
- Minimalist web framework
- High performance
- Extensible middleware
- Built-in support for data binding and validation

#### Chi
- Lightweight, idiomatic web framework
- Focused on composability and maintainability
- Efficient router
- Middleware support

#### Iris
- Fast, full-featured web framework
- Built-in MVC support
- Rich routing capabilities
- Extensive middleware ecosystem

### Databases

- **PostgreSQL**: Advanced open-source database
- **MySQL**: Popular open-source database
- **SQLite**: Lightweight file-based database

### ORMs

- **GORM**: The most popular Go ORM
- **XORM**: Simple and powerful ORM
- **Ent**: Facebook's entity framework

## Development

### Requirements

- Go 1.16+
- Git

### Building from source

```bash
git clone https://github.com/yourusername/goscaf.git
cd goscaf
go build
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Standard Go Project Layout](https://github.com/golang-standards/project-layout)
- [Cobra](https://github.com/spf13/cobra)
- [Survey](https://github.com/AlecAivazis/survey)

## Roadmap

- [x] Add support for more web frameworks
- [x] Add Docker configuration
- [ ] Add GraphQL support

 