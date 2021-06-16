# Pizza shop 

An attempt to create pizza shop using microservice concept.  It has two major service pizza service to process orders and another service to do pizza and notify customers.
(Project work in progress)

## Architecture Diagram

![Pizza (1)](https://user-images.githubusercontent.com/39593586/121798118-178daa00-cc42-11eb-9ce9-31810ba795a3.png)


## Highlights
- Containerized both the microservices using Docker.
- Used Message queue to communicate between Pizza and Kitchen Microservice.
- Implemented worker pool to have fixed number of workers in Kitchen and process orders.
- Added Prometheus for event monitoring and Grafana for monitoring of those events.



## DB Diagram of Pizza Microservice

![Pizza App Order](https://user-images.githubusercontent.com/39593586/121783860-c6909e00-cbce-11eb-99d0-3aee63a537ad.png)


## DB Diagram of Kitchen Microservice

![Kitchen Microservice](https://user-images.githubusercontent.com/39593586/121783797-508c3700-cbce-11eb-94f8-f665da6159c7.png)

## Grafana Dashboard

<img width="1440" alt="Screenshot 2021-06-13 at 12 55 15 PM" src="https://user-images.githubusercontent.com/39593586/121798920-aef4fc00-cc46-11eb-9fa4-2380c14ac1ab.png">
