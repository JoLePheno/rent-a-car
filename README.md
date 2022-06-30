# rent-a-car

Back-end challenge by stoïk (https://stoikio.notion.site/Backend-Software-Engineer-ecd3188d8ba84d18bdf2a6f844f120b7)

# Get started

## Build and run project:

This project is containerized you can use docker-compose to build and run the project:

`docker-compose build && docker-compose up -d`

To run the scalable version user:

`docker swarm init`
`docker stack deploy -c docker-compose.yml stoikbackend`

To view active container:

`docker ps`

## Interact with the API:

You can fin in the `pkg` directory a postman collection to run pre-made request.
In order to rent a car you'll need to create an user with *CreateUser*, if you're user doesn't suits you you can delete it as well.

Once you have created your user you can start the rental process, first get your *userID* using the *GetUser* request. Once you have your *userID* fetch the list of cars with *GetCars* request, there you'lle be able to make a *RentCar* request and get an available car using your *userID* with the *carID*.
When you're done with your car you can freed it using the *FreeCar* request.
