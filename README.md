# api-users

Side-project of a user management Go Echo API, supporting authentication (based on go-JWT) and authorization (based on casbin). The application has been the subject of a series of articles available here : (TODO PUT LINK HERE) (french content). 

## Quick start 

    git clone 
    cd api-users
    docker compose up -d


## How to use

All interactions are done via http requests.

Postman collection and environement are available in ressources/postman. You can update the fields 'domain' and 'users_api_port' of the environment if you deployed it remotely. 

You can authenticate at the URL "/auth/signin" with a login and password, the API will send you back a refresh token and an access token.

When the access token expires, you can get a new one without signin back, at url "/auth/refresh" by giving your refresh token. 

All requests to "/users/*" require authentication with JWT access token in the authorization header :"Authorization:bearer <your_access_token>".
You can create, update, list and delete users with the HTTP verbs POST (on /users),PUT,GET and DELETE (ON /users/<user_id>). 

## Configuration

The configuration is separated into several YAML blocks, to configure different parts of the application.

### Database configuration

        mysql:
            username: appli
            password: <db_user_password>
            host: <your_mysql_host>
            port: 3306
            database: api-users

Database configuration block, you can put the database password in environment variable "MYSQL_PASSWORD". The database must exist on the sql instance. 

### API configuration

        api:
            port: 8080
            prometheus_exporter_port: 8081

Configure API listening ports : 'port' is the application port and 'prometheus_exporter_port' is the monitoring port.

### JWT config

        jwt:
            refresh_token:
                private_cert: /certs/refresh_jwt.key
                public_cert: /certs/refresh_jwt.key.pub
            access_token:
                private_cert: /certs/access_jwt.key
                public_cert: /certs/access_jwt.key.pub

Configuration for the 2 JWT tokens : you must provide a valid path to an RSA keypair for both jwt. 

### RBAC config

        rbac: 
            model: /rbac/rbac_model.conf
            policy: /rbac/rbac_policy.csv

You must provide a valid path to the casbin model and casbin policy. If you need more roles or more policy, you can update policy file. Check the casbin doc for more informations : https://casbin.org/docs/overview. 

### Security config

        security:
            bcrypt_hash_cost: 13
            enable_su: true
            su_login: super_user
            su_password: <su_user_password>

bcrypt_hash_cost corresponds to the cost to hash a password. High cost = more security against a bruteforce, but reduced quality of user experience (higher response time on the login endpoint).

enable_su allows you to enable or no the superuser. su_login and su_password corresponds to his credentials. The password can be set in the environment variable "SU_PASSWORD"
