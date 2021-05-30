FROM golang:1.15

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/VarthanV/kitchen

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

#Set the environment variables
ENV DB_NAME "pizza_kitchen"
ENV DB_USERNAME "root"
ENV DB_PASSWORD "root1234"
ENV DB_PORT 3306
ENV DB_HOST "mysql"
ENV RABBIT_MQ_CONNECTION_STRING  "amqp://guest:guest@host.docker.internal/" 
ENV REDIS_HOST "pizza_redis"
# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# This container exposes port 8082 to the outside world
EXPOSE 8082

# Run the executable
CMD ["kitchen"]