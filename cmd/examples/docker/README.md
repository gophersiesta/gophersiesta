Example of how siesta could be used in a docker container
===

```
docker-compose up
```

This command should build (if not done already) the docker container and run it.

The container should exit immediately after *siesta* have loaded the configuration from the server and exported the variables found to the environment

SIESTA_HOST: specify the host where the siesta server is running
SIESTA_APP: specify the app
SIESTA_LABELS: specify the labels separated by commas