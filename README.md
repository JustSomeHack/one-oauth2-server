# OAuth2 Server

Currently authenticates against LDAP or a Mongo database

### Example environment variables for LDAP
```json
"AUTH_TYPE": "ldap",
"USER_SECRET": "token-secret",
"LDAP_BIND_DN": "dc=example,dc=com",
"LDAP_BIND_USER": "cn=admin,dc=example,dc=com",
"LDAP_BIND_PASSWORD": "password",
"LDAP_SERVER": "ldap.example.com",
"LDAP_PORT": "389",
"LDAP_USE_TLS": "true",
"PRODUCTION": "true"
```

### Example environment variables for Mongo
```json
"AUTH_TYPE": "mongo",
"USER_SECRET": "token-secret",
"MONGO_URL": "mongo.example.com",
"DB_NAME": "userdb",
"DB_USER": "username",
"DB_PASS": "password",
"PRODUCTION": "true"
```
