# Doctor Service

The Doctor Service is a component of the hospital management system, responsible for managing information related to doctors.

## Table of Contents

- [Overview](#overview)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
  - [Endpoints](#endpoints)
- [Contributing](#contributing)
- [License](#license)

## Overview

The Doctor Service provides functionalities for creating, updating, listing, deleting, and retrieving information about doctors.

## Getting Started

### Prerequisites

- Go (version 1.16 or higher)
- MySQL

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/iamrosada/hospital-management.git
   ```

2. Navigate to the `doctor-service` directory:

   ```bash
   cd hospital-management/doctor-service
   ```

3. Install dependencies:

   ```bash
   go get -u ./...
   ```

4. Build and run the application:

   ```bash
   go run main.go
   ```

## Usage

### Endpoints

- **Create Doctor:**

  ```http
  POST /doctors
  ```

  Example request body:

  ```json
  {
    "name": "Dr. John Doe",
    "bi": "123456789",
    "specialty": "Cardiology",
    "experience": 10,
    "email": "john.doe@example.com",
    "phone_number": "+1234567890"
  }
  ```

- **List Doctors:**

  ```http
  GET /doctors
  ```

- **Get Doctor by ID:**

  ```http
  GET /doctors/{id}
  ```

- **Update Doctor by ID:**

  ```http
  PUT /doctors/{id}
  ```

  Example request body:

  ```json
  {
    "name": "Dr. Jane Smith",
    "bi": "987654321",
    "specialty": "Orthopedics",
    "experience": 15,
    "email": "jane.smith@example.com",
    "phone_number": "+9876543210"
  }
  ```

- **Delete Doctor by ID:**

  ```http
  DELETE /doctors/{id}
  ```

## Contributing

Contributions are welcome! If you have any suggestions, improvements, or find any issues, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

