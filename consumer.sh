#!/bin/bash

docker-compose exec -it kafka kafka-console-consumer --topic logs --from-beginning --bootstrap-server localhost:9092
