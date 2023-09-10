# Supreme Security

Project aim is to show an example of mTLS (mutual TLS) communication between client and server:

## How to run?

The project is equipped with `docker-compose.yaml` file together with all the needed configuration in `dev.server.env` and `dev.client.env` in order to be started:

The following command will start the `server` as well as the `client` in detached mode:

```
    make compose
```

When finished one can run the below command in order to stop the running containers: 
```
    make decompose
```

The `server` and the `client` can be run separately:

```
    make server
    make client
```

## How to use? 

The example is not interactive - the `server` will run and the `client` will send 5 requests (one every second) to the only available endpoint - `/hellos`.

## Useful information

The below commands were used to generate the certificates needed by this mTLS example:

### Create root certificate

1. generate [**ca.key**] and [**ca.csr**] command: 

    ```$ openssl req -newkey rsa:2048 -keyout ca.key -out ca.csr```

2. sign [**ca.crt**] from [**ca.csr**] with [**ca.key**]:

    ```$ openssl x509 -signkey ca.key -in ca.csr -req -days 3650 -out ca.crt```

### Sign certificate

1. sign [**client.crt**] from [**client.csr**] and [**client.ext**] with [**ca.key**] and [**ca.crt**] 

    ```$ openssl x509 -req -CA ca.crt -CAkey ca.key -in client.csr -out client.crt -days 3650 -CAcreateserial -extfile client.ext```

### Remove pass for encrypted key

1. Remove password from [**server.key**] -> [**server.unencrypted.key**]

    ```$ openssl rsa -in server.key -out server.unencrypted.key -passin pass:this_is_server_key_password```

### Install certificate

1. This step is optional instead of using `RootCA` directly in code

    ```
        ca.crt /usr/local/share/ca-certificates/ca.crt
        chmod 644 /usr/local/share/ca-certificates/ca.crt && update-ca-certificates
    ```
