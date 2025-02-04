# GoWeb CLI

A command-line interface tool for quickly generating Go web application boilerplate code with modern project structure and best practices.

## Features

- 🚀 Quick project scaffolding
- 🎯 Multiple web framework options (Fiber/Gin)
- 🗄️ Database integration (Postgres/MySQL/SQLite)
- 🔧 ORM support (GORM/XORM/Ent)
- 📁 Standard project layout
- 🔄 Automated dependency installation

## Installation

### Prerequisites

- Go 1.16 or higher
- Git

### Install using Go
```bash
 go install github.com/yourusername/goweb@latest
 ```

## Getting Started.

### Create a new project

```bash
goweb init
```

This will start an interactive prompt asking for:

1. Project name
2. Web framework selection (Fiber/Gin)
3. Database system (Postgres/MySQL/SQLite)
4. ORM preference (Yes/No)
5. If Yes to ORM, choose between GORM/XORM/Ent

### Example

```bash
$ goweb init
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
│   ├── database/        # Database connection setup
│   ├── middleware/      # HTTP middleware
│   ├── models/          # Data models
│   ├── repositories/    # Data access layer
│   ├── services/        # Business logic
│   ├── handlers/        # HTTP request handlers
│   └── routes/          # Route definitions
├── pkg/                 # Public library code
├── scripts/             # Build and deployment scripts
└── .env                 # Environment variables
```

## Directory Structure Explanation

- `cmd/`: Contains the main applications for the project
  - `main.go`: The entry point of the application
- `config/`: Configuration file templates or default configs
- `internal/`: Private application and library code
  - `database/`: Database connection and configuration
  - `middleware/`: HTTP middleware components
  - `models/`: Application data models
  - `repositories/`: Data access layer implementation
  - `services/`: Business logic implementation
  - `handlers/`: HTTP request handlers
  - `routes/`: Route definitions and setup
- `pkg/`: Library code that's safe to use by external applications
- `scripts/`: Scripts for development, CI/CD, etc.

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
git clone https://github.com/yourusername/goweb.git
cd goweb
go build
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Standard Go Project Layout](https://github.com/golang-standards/project-layout)
- [Cobra](https://github.com/spf13/cobra)
- [Survey](https://github.com/AlecAivazis/survey)

## Support

If you have any questions or need help, please:
1. Check the readme again
2. Open an issue
3. Start a discussion

## Roadmap

- [ ] Add support for more web frameworks
- [ ] Include testing templates
- [ ] Add Docker configuration
- [ ] Generate API documentation
- [ ] Add GraphQL support

 