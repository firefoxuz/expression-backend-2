### Supported Operators:
- Addition: `+`
- Subtraction: `-`
- Multiplication: `*`
- Division: `/`

### Supported Operand Range:
- Positive integers: 0-9

## Installation

Follow these steps to set up the Expression Calculator Service and Expression Agent:

### Expression Calculator Service

1. Clone the Expression Calculator Service project repository:

    ```bash
    git clone https://github.com/firefoxuz/expression-backend-2.git
    ```

2. Navigate to the project directory:

    ```bash
    cd expression-backend-2
    ```

3. Copy the example environment configuration file:

    ```bash
    cp .env.json.example .env.json
    ```

4. Start the Docker containers using Docker Compose:

    ```bash
    docker-compose up -d
    ```

5. Apply database migrations:

    ```bash
    docker run -v ./db/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database postgres://expression_user:expression_password@localhost:5444/expression_db?sslmode=disable up
    ```

### Expression Agent

6. Clone the Expression Agent project repository:

    ```bash
    git clone https://github.com/firefoxuz/expression-agent-2.git
    ```

7. Navigate to the project directory:

    ```bash
    cd expression-agent-2
    ```

8. Copy the example environment configuration file:

    ```bash
    cp .env.json.example .env.json
    ```

9. Start the Docker containers using Docker Compose:

    ```bash
    docker-compose up -d
    ```

Now, both the Expression Calculator Service and Expression Agent should be up and running. You can access the Expression Calculator Service at [http://127.0.0.1:8085/](http://127.0.0.1:8085/) and the Expression Agent at their respective API endpoints.

## API Documentation

### Register User

Registers a new user with the provided login credentials.

- **URL**

  `/api/v1/register`

- **Method**

  `POST`

- **Request Body**

  | Field    | Type   | Description     |
    |----------|--------|-----------------|
  | login    | string | User's username |
  | password | string | User's password |

  Example:
  ```json
  {
      "login": "username",
      "password": "123456"
  }
  ```

- **Success Response**

   - **Code:** 200 OK
     ```json
     {
         "message": "User registered successfully"
     }
     ```

- **Error Response**

   - **Code:** 400 Bad Request
   - **Content:**
     ```json
     {
         "error": "Invalid request body"
     }
     ```
### User Login

Logs in a user with the provided login credentials.

- **URL**

  `/api/v1/login`

- **Method**

  `POST`

- **Request Body**

  | Field    | Type   | Description     |
    |----------|--------|-----------------|
  | login    | string | User's username |
  | password | string | User's password |

  Example:
  ```json
  {
      "login": "username",
      "password": "123456"
  }
  ```

- **Success Response**

   - **Code:** 200 OK
   - **Content:**
     ```json
     {
         "token_type": "jwt",
         "token": "some_generated_token"
     }
     ```

- **Error Response**

   - **Code:** 401 Unauthorized
   - **Content:**
     ```json
     {
         "error": "Invalid login credentials"
     }
     ```
### Add Expressions

Adds expressions with a time limit.

- **URL**

  `/api/v1/expressions`

- **Method**

  `POST`

- **Request Headers**

  | Header         | Value                                       |
    |----------------|---------------------------------------------|
  | Content-Type   | application/json                            |
  | Authorization  | Bearer \<access_token\>                     |

- **Request Body**

  | Field         | Type    | Description                    |
    |---------------|---------|--------------------------------|
  | expression    | string  | Mathematical expression        |
  | time_limit    | integer | Time limit for expression (ms) |

  Example:
  ```json
  {
      "expression": "3+3+3",
      "time_limit": 1000
  }
  ```

- **Success Response**

    - **Code:** 200 OK
    - **Content:**
      ```json
      {
          "message": "expression is stored and will be calculated"
      }
      ```

- **Error Response**

    - **Code:** 400 Bad Request
    - **Content:**
      ```json
      {
          "error": "Invalid request body"
      }
      ```
  OR
    - **Code:** 401 Unauthorized
    - **Content:**
      ```json
      {
          "error": "Unauthorized"
      }
      ```


### Fetch All User Expressions

Retrieves all expressions stored for the authenticated user.

- **URL**

  `/api/v1/expressions`

- **Method**

  `GET`

- **Request Headers**

  | Header         | Value                                       |
    |----------------|---------------------------------------------|
  | Content-Type   | application/json                            |
  | Authorization  | Bearer \<access_token\>                     |

- **Success Response**

  - **Code:** 200 OK
    - **Content:**
        ```json
        {
            "expressions": [
                {
                    "id": 2,
                    "user_id": 1,
                    "expression": "3+3+3",
                    "result": 9,
                    "is_processing": false,
                    "is_time_limit": null,
                    "is_valid": true,
                    "is_finished": true,
                    "time_limit": 1000,
                    "created_at": "2024-04-21T22:11:55Z",
                    "finished_at": "2024-04-21T17:11:57.92291Z"
                },
                {
                    "id": 1,
                    "user_id": 1,
                    "expression": "3+3+3",
                    "result": 9,
                    "is_processing": false,
                    "is_time_limit": null,
                    "is_valid": true,
                    "is_finished": true,
                    "time_limit": 1000,
                    "created_at": "2024-04-21T16:57:30Z",
                    "finished_at": "2024-04-21T17:10:44.748884Z"
                }
            ]
      }
        ```
### Fetch Expression by ID

Retrieves a single expression by its ID.

- **URL**

  `/api/v1/expressions/{expression_id}`

- **Method**

  `GET`

- **URL Parameters**

  | Parameter      | Type   | Description             |
    |----------------|--------|-------------------------|
  | expression_id  | string | ID of the expression    |

- **Request Headers**

  | Header         | Value                                       |
    |----------------|---------------------------------------------|
  | Authorization  | Bearer \<access_token\>                     |

  - **Success Response**

      - **Code:** 200 OK
      - **Content:**
        ```json
          {
              "expressions": {
                  "id": 2,
                  "user_id": 1,
                  "expression": "3+3+3",
                  "result": 9,
                  "is_processing": false,
                  "is_time_limit": null,
                  "is_valid": true,
                  "is_finished": true,
                  "time_limit": 1000,
                  "created_at": "2024-04-21T22:11:55Z",
                  "finished_at": "2024-04-21T17:11:57.92291Z"
              }
          }
        ```

- **Error Responses**

    - **Code:** 401 Unauthorized
      **Content:**
      ```json
      {
          "error": "Unauthorized"
      }
      ```
    - **Code:** 404 Not Found
      **Content:**
      ```json
      {
          "error": "Expression not found"
      }
      ```


## Microservice Architecture

Below is the architecture of our microservices:

![Microservice Architecture](https://i.postimg.cc/YtNQJ7s5/2024-04-21-21-43.png?dl=1)
