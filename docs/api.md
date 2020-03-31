# API

## GET /sensors

Przykład:

```sh
curl localhost:8080/sensors
```

Odpowiedź:

```json
{
    // epoka uniksowa (sekundy)
    "timestamp": 123456789,
    // temperatura (stopnie celsjusza)
    "temperature": 20.5,
    // wilgotność (procent)
    "humidity": 50
}
```

## POST /sensors

!!! note
    Wymaga autentykacji tokenem.

Przykład:

```sh
curl localhost:8080/sensors -X POST -d '{"temperature":11}' --header 'authorization: Bearer TOKEN' --header 'content-type: application/json'
```

Żądanie i odpowiedź:

```json
{
    // epoka uniksowa (sekundy)
    "timestamp": 123456789,
    // temperatura (stopnie celsjusza)
    "temperature": 20.5,
    // wilgotność (procent)
    "humidity": 50
}
```
