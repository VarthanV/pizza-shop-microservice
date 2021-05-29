FROM golang:1.15

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/VarthanV/pizza

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

#Set the environment variables
ENV ACCESS_TOKEN_SECRET_KEY  "foo"
ENV REFRESH_TOKEN_SECRET_KEY "bar"
ENV DB_NAME "pizza_shop"
ENV DB_USERNAME "root"
ENV DB_PASSWORD "root1234"
ENV DB_PORT 3306
ENV DB_HOST "pizzasql"
# Yet to containerize rabbit-mq i have it setup as a seperate container in my machine so giving my ip address
ENV RABBIT_MQ_CONNECTION_STRING  "amqp://guest:guest@pizza_rmq/" 
# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["pizza"]
