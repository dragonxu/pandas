# Authentication service

Authentication service provides an API for managing authentication keys.

There are *three types of authentication keys*:

- user key - keys issued to the user upon login request
- API key - keys issued upon the user request
- recovery key - password recovery key

User keys are issued when user logs in. Each user request (other than `registration` and `login`) contains user key that is used to authenticate the user. API keys are similar to the User keys. The main difference is that API keys have configurable expiration time. If no time is set, the key will never expire. For that reason, API keys are _the only key type that can be revoked_. Recovery key is the password recovery key. It's short-lived token used for password recovery process.

For in-depth explanation of the aforementioned scenarios, as well as thorough
understanding of Mainflux, please check out the [official documentation][doc].

The following actions are supported:

- create (all key types)
- verify (all key types)
- obtain (API keys only; secret is never obtained)
- revoke (API keys only)

## Configuration

The service is configured using the environment variables presented in the
following table. Note that any unset variables will be replaced with their
default values.

| Variable                  | Description                                                              | Default       |
|---------------------------|--------------------------------------------------------------------------|---------------|
| PD_AUTHN_LOG_LEVEL        | Service level (debug, info, warn, error)                                | error          |
| PD_AUTHN_DB_HOST          | Database host address                                                   | localhost      |
| PD_AUTHN_DB_PORT          | Database host port                                                      | 5432           |
| PD_AUTHN_DB_USER          | Database user                                                           | mainflux       |
| PD_AUTHN_DB_PASSWORD      | Database password                                                       | mainflux       |
| PD_AUTHN_DB               | Name of the database used by the service                                | auth           |
| PD_AUTHN_DB_SSL_MODE      | Database connection SSL mode (disable, require, verify-ca, verify-full) | disable        |
| PD_AUTHN_DB_SSL_CERT      | Path to the PEM encoded certificate file                                |                |
| PD_AUTHN_DB_SSL_KEY       | Path to the PEM encoded key file                                        |                |
| PD_AUTHN_DB_SSL_ROOT_CERT | Path to the PEM encoded root certificate file                           |                |
| PD_AUTHN_HTTP_PORT        | Authn service HTTP port                                                 | 8180           |
| PD_AUTHN_GRPC_PORT        | Authn service gRPC port                                                 | 8181           |
| PD_AUTHN_SERVER_CERT      | Path to server certificate in pem format                                |                |
| PD_AUTHN_SERVER_KEY       | Path to server key in pem format                                        |                |
| PD_AUTHN_SECRET           | String used for signing tokens                                          | auth           |
| PD_JAEGER_URL             | Jaeger server URL                                                       | localhost:6831 |

## Deployment

The service itself is distributed as Docker container. The following snippet
provides a compose file template that can be used to deploy the service container
locally:

```yaml
version: "2"
services:
  authn:
    image: authn:[version]
    container_name: [instance name]
    ports:
      - [host machine port]:[configured HTTP port]
    environment:
      PD_AUTHN_LOG_LEVEL: [Service log level]
      PD_AUTHN_DB_HOST: [Database host address]
      PD_AUTHN_DB_PORT: [Database host port]
      PD_AUTHN_DB_USER: [Database user]
      PD_AUTHN_DB_PASS: [Database password]
      PD_AUTHN_DB: [Name of the database used by the service]
      PD_AUTHN_DB_SSL_MODE: [SSL mode to connect to the database with]
      PD_AUTHN_DB_SSL_CERT: [Path to the PEM encoded certificate file]
      PD_AUTHN_DB_SSL_KEY: [Path to the PEM encoded key file]
      PD_AUTHN_DB_SSL_ROOT_CERT: [Path to the PEM encoded root certificate file]
      PD_AUTHN_HTTP_PORT: [Service HTTP port]
      PD_AUTHN_GRPC_PORT: [Service gRPC port]
      PD_AUTHN_SECRET: [String used for signing tokens]
      PD_AUTHN_SERVER_CERT: [String path to server certificate in pem format]
      PD_AUTHN_SERVER_KEY: [String path to server key in pem format]
      PD_JAEGER_URL: [Jaeger server URL]
```

To start the service outside of the container, execute the following shell script:

```bash
# download the latest version of the service
go get github.com/cloustone/pandas/mainflux

cd $GOPATH/src/github.com/cloustone/pandas/mainflux

# compile the service
make authn

# copy binary to bin
make install

# set the environment variables and run the service
PD_AUTHN_LOG_LEVEL=[Service log level] PD_AUTHN_DB_HOST=[Database host address] PD_AUTHN_DB_PORT=[Database host port] PD_AUTHN_DB_USER=[Database user] PD_AUTHN_DB_PASS=[Database password] PD_AUTHN_DB=[Name of the database used by the service] PD_AUTHN_DB_SSL_MODE=[SSL mode to connect to the database with] PD_AUTHN_DB_SSL_CERT=[Path to the PEM encoded certificate file] PD_AUTHN_DB_SSL_KEY=[Path to the PEM encoded key file] PD_AUTHN_DB_SSL_ROOT_CERT=[Path to the PEM encoded root certificate file] PD_AUTHN_HTTP_PORT=[Service HTTP port] PD_AUTHN_GRPC_PORT=[Service gRPC port] PD_AUTHN_SECRET=[String used for signing tokens] PD_AUTHN_SERVER_CERT=[Path to server certificate] PD_AUTHN_SERVER_KEY=[Path to server key] PD_JAEGER_URL=[Jaeger server URL] $GOBIN/mainflux-authn
```

If `PD_EMAIL_TEMPLATE` doesn't point to any file service will function but password reset functionality will not work.

## Usage

For more information about service capabilities and its usage, please check out
the [API documentation](swagger.yaml).

[doc]: http://mainflux.readthedocs.io
