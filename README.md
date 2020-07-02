# My Prometheus Playground

This is a rather simple repo with a docker-compose file to setup prometheus to monitor a go web server. Grafana is used to display the metrics.

## Dependencies

    1. Docker
    2. Hey https://github.com/rakyll/hey


## Make commands

There are make commands for different operations, `make fire` runs several requests to the go webserver.