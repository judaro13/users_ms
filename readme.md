### App description
This app is like a messages consumer for rabbitmq.
This app fetch any message in a specified channel in rabbitmq and validate if that message have a correct json structure of a user then persist the data in a PostgreSQL database.

The messages are fetched in string format, parsed as a json object and validated the data. If the data have all the field required, a user object with the data is send for stored. If the data is not valid, is showed a message with the data gived.

In this project was used  streadway/amqp to handle rabbitmq queue and database/sql to store data in a PostgreSQL database.


### Install
No special instructions for install, just needed some env vars for rabbitqm

export RABBIT_PATH="amqp://guest:guest@localhost:5672/"
export RABBIT_CHANNEL="postUser"
export DB_HOST="localhost"
export DB_PORT=5432
export DB_USER="test"
export DB_PASSWORD="test"
export DB_NAME="test"
