# README

- [README](#readme)
  - [Requirements](#requirements)
  - [Building on file save](#building-on-file-save)
  - [JWT](#jwt)
    - [Generate token with jwt-generator](#generate-token-with-jwt-generator)
    - [Get data](#get-data)
  - [Build on save](#build-on-save)
  - [Apache](#apache)
  - [Systemd](#systemd)

## Requirements

```shell
go get -u -v github.com/dgrijalva/jwt-go
go get -u -v github.com/joho/godotenv v1.3.0
go get -u -v github.com/rs/cors v1.7.0
go get -u -v github.com/tealeg/xlsx v1.0.5
```

## Building on file save

```shell
while inotifywait -e close_write *.go; do go build ; done
```

## JWT

### Generate token with jwt-generator

- <https://github.com/rybit/jwt-generator>

```shell
jwt-generator gen -s secret -u user
```

- `-s` the secret / signingKey
- `-u` the user to have for the token

### Get data

```shell
curl -H "Token: XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX" localhost:3000/tonne
```

## Build on save

```shell
while inotifywait -e close_write *.go; do go build ; done
```

## Apache

```apache
<VirtualHost>
ProxyPass "/tonne/abfallberatungen" http://127.0.0.1:3000/tonne/abfallberatungen
ProxyPassReverse "/tonne/abfallberatungen" http://127.0.0.1:3000/tonne/abfallberatungen
</VirtualHost>
```

```shell
service apache2 reload
```

## Systemd

```shell
ln -s /home/intevation/tonne/tonne.service /lib/systemd/system/tonne.service
```

```shell
systemctl start tonne
systemctl status tonne
systemctl stop tonne
systemctl restart tonne
```
