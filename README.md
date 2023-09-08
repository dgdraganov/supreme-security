# supreme-security

Project aim is to show an example of mTLS (mutual TLS)

## How to run?

The project is equipped with `docker-compose.yaml` file together with all the needed configuration in `dev.server.env` and `dev.client.env` in order to be started:

The following command will start the fuzzy serivice as well as the DB locally in docker detached mode:
```
    make compose
```

When finished one can run the below command in order to stop the running containers: 
```
    make decompose
```

## How to use? 

The example is not interactive - the `server` will run and the `client` will start making requests (every second) to the only available endpoint - `/hellos`.
