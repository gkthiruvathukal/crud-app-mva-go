
# Go CRUD App

This project provides a set of command-line tools for managing user data using Go and GORM (a popular Go ORM). The tools support Create, Read, Update, and Delete (CRUD) operations and include an option to create users from a YAML file. The database used is SQLite, a lightweight, self-contained SQL database engine.

## Overview

The project includes the following tools:
- `create_user`: Creates a new user in the database.
- `read_user`: Reads and displays user details by ID.
- `update_user`: Updates the details of an existing user.
- `delete_user`: Deletes a user by ID.

Additionally, a `Makefile` is provided to simplify the build process.

## Tools and Libraries

- **Go**: A statically typed, compiled programming language known for its simplicity and performance.
- **GORM**: An ORM library for Go that provides convenient methods for interacting with databases.
- **SQLite**: A lightweight, file-based SQL database engine used for simplicity and ease of setup.
- **gopkg.in/yaml.v2**: A library for parsing YAML files, used for reading user data from YAML in the `create_user` tool.

## Installation and Setup

1. **Extract the ZIP File**:
   Unzip the downloaded file into your desired directory.

2. **Install Dependencies**:
   Ensure you have Go installed on your system. Then, navigate to the project directory and run:
   ```bash
   go mod tidy
   ```

3. **Build the Programs**:
   Use the Makefile to compile the programs:
   ```bash
   make all
   ```

## Usage Examples

### Creating a User

To create a user from command-line arguments:
```bash
./create_user --name="John Doe" --email="john.doe@example.com" --note="Sample note" --tags="tag1,tag2"
```

To create a user from a YAML file:
```bash
./create_user --from-yaml=user_data.yaml
```

Example `user_data.yaml`:
```yaml
name: John Doe
email: john.doe@example.com
note: A sample note
tags:
  - tag1
  - tag2
```

### Reading a User

To read and display a user's details by ID:
```bash
./read_user --id=1
```

### Updating a User

To update a user's details:
```bash
./update_user --id=1 --name="Jane Doe" --email="jane.doe@example.com" --note="Updated note"
```

### Deleting a User

To delete a user by ID:
```bash
./delete_user --id=1
```

## Cleaning Up

To remove the compiled binaries:
```bash
make clean
```

## License

This project is licensed under the MIT License.

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request.

## Acknowledgements

- [GORM](https://gorm.io/) for the ORM functionalities.
- [SQLite](https://www.sqlite.org/index.html) for the lightweight database solution.
- [Go](https://golang.org/) for the programming language and its powerful standard library.
