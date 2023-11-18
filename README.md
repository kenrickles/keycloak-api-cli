# Keycloak API CLI Application

## Setup and Execution
1. **Install Go** - Ensure Go (version 1.21.4 or higher) is installed.
2. **Clone the Repository** - Clone the application repository to your local machine.
3. **Install Dependencies** - Run `go mod tidy` in the project directory to install necessary dependencies.
4. **Configuration** - Edit the `config.yaml` file with appropriate Keycloak server details and credentials.
5. **Build the Application** - Run `go build .` in the project directory to build the executable.
6. **Make the Build Executable** - Use `chmod +x keycloak-api-cli` (or the name of your built executable) to make it executable.
7. **Running the Application** - Execute the built application. Use `./keycloak-api-cli --config path/to/config.yaml` to specify the configuration file.

## Available Commands

### 1. `create`
- **Usage**: `create [user, realm, resource, role]`
- **Description**: Creates various entities in Keycloak.
- **Subcommands**:
  - `create user`: Creates a new user.
  - `create realm`: Creates a new realm.
- **Example**: `./keycloak-api-cli create user --config path/to/config.yaml`

### 2. `delete`
- **Usage**: `delete [user, realm, resource, role]`
- **Description**: Deletes entities in Keycloak.
- **Subcommands**:
  - `delete user`: Deletes a user.
  - `delete realm`: Deletes a realm.
- **Example**: `./keycloak-api-cli delete realm --config path/to/config.yaml`

### 3. `get`
- **Usage**: `get [users, realms, resources, roles]`
- **Description**: Retrieves information about entities in Keycloak.
- **Subcommands**:
  - `get users`: Lists all users in a realm.
  - `get realms`: Lists all realms.
- **Example**: `./keycloak-api-cli get realms --config path/to/config.yaml`

### 4. Root Command
- **Usage**: `keycloak-api-cli`
- **Description**: Base command for Keycloak API interactions.
- **Integrated Commands**: Includes `create`, `delete`, and `get`.
- **Example**: `./keycloak-api-cli --help`

## Additional Notes
- The application's functionality can be extended by adding more subcommands.
- Ensure the Keycloak server details in `config.yaml` are correct before using the commands.
