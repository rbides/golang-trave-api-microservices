# golang-trave-api-microservices - Work In Progress

The backend for a travel app in golang, split into microservices. Ideally each microservice would be in it's own repository, but for exercising's sake I'm keeping them all in the same repo.

- travel-api: Handles the Travel object which will have N Seats available for reservation.
- reservation-api: Stores the reservation orders
- user-api: Stores user information and handle authentication/authorization
- payments-api: handle payments
- orchestrator-api: Starting point for Order creation/cancelation and orchestrating the required microservice's calls

To-do list:
- Add JWT based authentication for users api
- Handle role authorization for creating travels and updating travels, for example
- Require authentication for creating orders
- Handle (mocking) payments and refunds
- Implement SAGA pattern in orchestrator for dealing with distributed transactions and compensating the transactions as required on failure
- Decouple microservice's calls with SNS + SQS
