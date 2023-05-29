# demo

This project is a sample information management system. It allows you to add information, view information, and toggle the completion status of the information.

## Installation

1. Clone the project to your local machine:
git clone https://github.com/your-username/your-repo.git


2. Navigate to the project directory:
cd your-repo


3. Install the required dependencies:
go get -u github.com/gin-gonic/gin


4. Build and run the project:
go run main.go


5. The server will start running on `localhost:9090`.

## API Endpoints

- `GET /informations`: Retrieve all information in the system.
- `GET /informations/:id`: Retrieve a specific information by ID.
- `PATCH /informations/:id`: Toggle the completion status of an information by ID.
- `POST /informations`: Add a new information.

## Usage

1. Use a tool like cURL or Postman to send HTTP requests to the API endpoints.
2. To add a new information, send a POST request to `/informations` with the following JSON payload:
 ```
{
 "name": "John Doe",
 "email": "johndoe@example.com",
 "phoneNumber": "1234567890",
 "completed": false
}

 ```
3. To retrieve all information, send a GET request to `/informations`.
4. To retrieve a specific information by ID, send a GET request to `/informations/:id`, replacing `:id` with the actual ID of the information.
5. To toggle the completion status of an information, send a PATCH request to `/informations/:id`, replacing `:id` with the actual ID of the information.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvement, please create a new issue or submit a pull request.

## License

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT).



