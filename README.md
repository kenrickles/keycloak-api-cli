# Keycloak API CLI Application

The Keycloak API CLI is a command-line tool that allows you to interact with the Keycloak authentication and authorization server. It simplifies the management of Keycloak entities, making it efficient for tasks such as creating users, realms, resources, and roles.

## Table of Contents

- [Keycloak API CLI Application](#keycloak-api-cli-application)
  - [Table of Contents](#table-of-contents)
  - [Setup and Execution](#setup-and-execution)
  - [Available Commands](#available-commands)
    - [1. `create`](#1-create)
    - [2. `delete`](#2-delete)
    - [3. `get`](#3-get)
    - [4. `retrieve token`](#4-retrieve-token)
    - [5. Root Command](#5-root-command)
  - [Additional Notes](#additional-notes)
    - [General Overview of Cobra and Viper](#general-overview-of-cobra-and-viper)
    - [Cobra and Viper in This Application](#cobra-and-viper-in-this-application)
    - [Why This CLI is Useful](#why-this-cli-is-useful)

## Setup and Execution

To get started with the Keycloak API CLI, follow these steps:

1. **Install Go** - Ensure Go (version 1.21.4 or higher) is installed on your system.

2. **Clone the Repository** - Clone the application repository to your local machine.

3. **Install Dependencies** - Run the following command in the project directory to install the necessary dependencies:

    ```bash
    go mod tidy
    ```

4. **Configuration** - Edit the `config.yaml` file with the appropriate Keycloak server details and credentials.

5. **Build the Application** - Run the following command in the project directory to build the executable:

    ```bash
    go build .
    ```

6. **Make the Build Executable** - Use the following command to make the executable file executable:

    ```bash
    chmod +x keycloak-api-cli
    ```

7. **Running the Application** - Execute the built application. You can specify the configuration file using the `--config` flag. Example usage:

    ```bash
    ./keycloak-api-cli --config path/to/config.yaml
    ```

## Available Commands

The Keycloak API CLI provides several commands to interact with Keycloak entities. Here are the available commands:

### 1. `create`

- **Usage**: `create [user, realm, resource, role]`
- **Description**: Creates various entities in Keycloak.
- **Subcommands**:
  - `create user`: Creates a new user in a specified realm.
    - **Flags**: `-u, --username <username>`, `-p, --password <password>`, `-e, --email <email>`, `-r, --realm <realm>`
    - **Example**: `./keycloak-api-cli create user -u john -p pass123 -e john@example.com -r myrealm`
  - `create realm`: Creates a new realm.
    - **Example**: `./keycloak-api-cli create realm --config path/to/config.yaml`

### 2. `delete`

- **Usage**: `delete [user, realm, resource, role]`
- **Description**: Deletes entities in Keycloak.
- **Subcommands**:
  - `delete user`: Deletes a user from a specified realm.
    - **Flags**: `-i, --userid <userID>`, `-u, --username <username>`, `-r, --realm <realm>`
    - **Example**: `./keycloak-api-cli delete user -u john -r myrealm`
  - `delete realm`: Deletes a realm.
    - **Example**: `./keycloak-api-cli delete realm --config path/to/config.yaml`

### 3. `get`

- **Usage**: `get [users, realms, resources, roles]`
- **Description**: Retrieves information about entities in Keycloak.
- **Subcommands**:
  - `get users`: Lists all users in a specified realm.
    - **Flags**: `-r, --realm <realm>`
    - **Example**: `./keycloak-api-cli get users -r myrealm --config path/to/config.yaml` 
  - `get realms`: Lists all realms.
    - **Example**: `./keycloak-api-cli get realms --config path/to/config.yaml`

### 4. `retrieve token`

- **Usage**: `retrieve token`
- **Description**: Retrieves an authentication token from Keycloak.
- **Flags**: 
  - `-c, --config`: Specifies the path to the configuration file (e.g., `path/to/config.yaml`).

Example usage:

```bash
./keycloak-api-cli retrieve token --config example-config.yaml
```

### 5. Root Command
- **Usage**: `keycloak-api-cli`
- **Description**: Base command for Keycloak API interactions.
- **Integrated Commands**: Includes `create`, `delete`, `get` and `retrieve token`.
- **Example**: `./keycloak-api-cli --help`

## Additional Notes
- The application's functionality can be extended by adding more subcommands.
- Ensure the Keycloak server details in `config.yaml` are correct before using the commands.


### General Overview of Cobra and Viper
- **Cobra**: A library for creating powerful command-line applications in Go. It's used for organizing commands, subcommands, and handling command-line inputs. Cobra facilitates the creation of both simple and complex CLI applications with features like command nesting, flags, and auto-generated help texts.

- **Viper**: A Go library for handling application configuration. Viper works well with environment variables, configuration files, and command-line flags. It supports reading from JSON, TOML, YAML, HCL, and Java properties config files. Viper allows easy retrieval of configuration settings, making it a flexible tool for managing dynamic configurations.

### Cobra and Viper in This Application
- **Cobra Usage**:
  - Organizes the main commands (`create`, `delete`, `get`) and their subcommands.
  - Parses command-line arguments and flags.
  - Generates help commands and documentation for each command.

- **Viper Usage**:
  - Manages application configurations, primarily loaded from `config.yaml`.
  - Facilitates runtime configuration changes, making the CLI adaptable to different environments or Keycloak setups.

### Why This CLI is Useful
- **Efficient Management**: Provides a command-line interface for managing Keycloak entities, which is faster and more efficient for many users compared to a GUI.
- **Automation Friendly**: Ideal for scripting and automating Keycloak administration tasks.
- **Customizable**: Easy to extend and customize for specific Keycloak management needs, thanks to the modular design using Cobra.
- **Configuration Flexibility**: With Viper, the CLI can be easily configured for different environments, enhancing its versatility.

This CLI leverages the strengths of Cobra and Viper to provide a robust, flexible, and user-friendly tool for Keycloak administration and automation, streamlining the management of authentication, authorization, and user federation within Keycloak realms.
