# Newsletter App API

This API provides services for managing newsletters and subscribers. It was designed with **hexagonal architecture** to achieve a clear separation between business logic and external interfaces, facilitating integration, maintenance, and testing. The API includes **Swagger documentation** to explore its endpoints and **unit tests** to verify system functionality.

## Link to Swagger Documentation

The Swagger documentation is available at [http://localhost:8080/docs/index.html](http://localhost:8080/docs/index.html). You can use Swagger to explore the endpoints, perform interactive testing, and better understand how to interact with the API.

## Running with Docker

This project includes a Dockerfile to facilitate the creation of a container.

```bash
docker build -t newsletter-back .
docker run -d -p 8080:8080 newsletter-back
```

## Environment Setup

This project uses environment variables for its configuration, which include:

- `mongoUrl`: Connection URL to MongoDB.
- `mongoDb`: Name of the MongoDB database.
- `mongoNewsletterCollection`: Name of the newsletters collection in MongoDB.
- `mongoSubscriberCollection`: Name of the subscribers collection in MongoDB.
- `emailSender`: Email address for sending newsletters.
- `emailPass`: Password for the email used to send newsletters.
- `smtpServer`: SMTP server for sending emails.
- `smtpPort`: SMTP port for sending emails.

## Features

### Newsletters

#### Get List of Newsletters

- **Method:** GET
- **Path:** `/api/v1/newsletters`
- **Description:** Retrieves a list of newsletters with optional search and pagination parameters.

  **Parameters:**

  - `name` (string, query): Name of the newsletter to search for.
  - `page` (integer, query): Page number for pagination.
  - `pageSize` (integer, query): Number of items per page for pagination.

  **Responses:**

  - Código 200 (OK)
  - Código 400 (Bad Request)
  - Código 500 (Internal Server Error)

#### Update an Existing Newsletter

- **Method:** PUT
- **Path:** `/api/v1/newsletters`
- **Description:** Allows an admin user to update an existing newsletter.

  **Parameters:**

  - `updateRequest` (object, body): Updated details of the newsletter.

  **Responses:**

  - Código 200 (OK)
  - Código 400 (Bad Request)
  - Código 500 (Internal Server Error)

#### Create a New Newsletter

- **Method:** POST
- **Path:** `/api/v1/newsletters`
- **Description:** Allows an admin user to create a new newsletter.

  **Parameters:**

  - `newsletter` (object, body): Details of the new newsletter.

  **Responses:**

  - Código 201 (Created)
  - Código 400 (Bad Request)
  - Código 500 (Internal Server Error)

#### Send Newsletter to Subscribers

- **Method:** POST
- **Path:** `/api/v1/newsletters/send/{newsletterID}`
- **Description:** Allows an admin user to send a newsletter to a list of subscribers.

  **Parameters:**

  - `newsletterID` (string, path): ID of the newsletter to send.

  **Responses:**

  - Código 200 (OK)
  - Código 400 (Bad Request)
  - Código 500 (Internal Server Error)

#### Delete a Newsletter

- **Method:** DELETE
- **Path:** `/api/v1/newsletters/{id}`
- **Description:** Allows an admin user to delete a newsletter.

  **Parameters:**

  - `id` (string, path): ID of the newsletter to delete.

  **Responses:**

  - Código 200 (OK)
  - Código 400 (Bad Request)
  - Código 500 (Internal Server Error)

### Subscribers

#### Subscribe to the Newsletter

- **Method:** POST
- **Path:** `/api/v1/subscribe/{email}/{category}`
- **Description:** Allows a user to subscribe to the newsletter.

  **Parameters:**

  - `email` (string, path): Email address for the subscription.
  - `category` (string, path): Category to subscribe to.

  **Responses:**

  - Código 200 (OK)
  - Código 400 (Bad Request)
  - Código 500 (Internal Server Error)

#### Obtener lista de suscriptores

- **Method:** GET
- **Path:** `/api/v1/subscribers`
- **Description:** Retrieves a list of subscribers with optional search and pagination parameters.

  **Parameters:**

  - `email` (string, query): Email address of the subscriber to search for.
  - `category` (string, query): Category of the subscriber to search for.
  - `page` (integer, query): Page number for pagination.
  - `pageSize` (integer, query): Number of items per page for pagination.

  **Responses:**

  - Código 400 (Bad Request)
  - Código 500 (Internal Server Error)

#### Get Subscriber by Email and Category

- **Method:** GET
- **Path:** `/api/v1/subscribers/{email}/{category}`
- **Description:** Retrieves details of a subscriber by email address.

  **Parameters:**

  - `email` (string, path): Email address to retrieve details.
  - `category` (string, path): Category the subscriber is subscribed to.

  **Responses:**

  - Código 200 (OK)
  - Código 404 (Subscriber not found)
  - Código 500 (Internal Server Error)

#### Unsubscribe from the Newsletter

- **Method:** DELETE
- **Path:** `/api/v1/unsubscribe/{email}/{category}`
- **Description:** Allows a user to unsubscribe from the newsletter.

  **Parameters:**

  - `email` (string, path): Email address to unsubscribe.
  - `category` (string, path): Category the user is unsubscribed from.

  **Responses:**

  - Código 200 (OK)
  - Código 400 (Bad Request)
  - Código 500 (Internal Server Error)
