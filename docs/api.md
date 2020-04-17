# API

## Get status

Returns server version and source git commit from which it was built.

**Endpoint**:

```http
GET /
```

**Response**:

```json
HTTP/1.1 200 OK
{
  "version": "refs/heads/master",
  "commit": "88694ddfb91153cd470bec20d23aee360feccdf7"
}
```

## Get latest sensors readings

Returns newest sensors readings found in the database.

**Endpoint**:

```http
GET /sensors
```

**Response**:

```json
HTTP/1.1 200 OK
{
  "timestamp": 1586532636,
  "temperature": 12,
  "pressure": 1001.5,
  "humidity": 41,
  "gas": 14
}
```

## Get sensors readings from `x` hours

Returns sensors readings that were created within `x` hours from now.

**Endpoint**:

```http
GET /sensors/:hours
```

**Parameters**:

| Parameter   | Description               | Type    |
|-------------|---------------------------|:-------:|
| hours       | Number of hours since now | integer |

**Response**:

```json
HTTP/1.1 200 OK
[
  ...
  {
    "timestamp": 1587128346,
    "temperature": 20.5,
    "pressure": 1000.2,
    "humidity": 60.1,
    "gas": 10
  },
  {
    "timestamp": 1587128347,
    "temperature": 20.5,
    "pressure": 1000.2,
    "humidity": 60.1,
    "gas": 10
  },
  {
    "timestamp": 1587128348,
    "temperature": 20.5,
    "pressure": 1000.2,
    "humidity": 60.1,
    "gas": 10
  }
  ...
]
```

## Create new sensors readings

Create a new entry in the database with given sensors readings.

**Endpoint**:

```http
POST /sensors
```

**Headers**:

| Name           | Value             |
|----------------|-------------------|
| Content-Type   | application/json  |
| Authentication | Bearer **$TOKEN** |

**Data**:

| Parameter   | Description    | Unit       | Default |
|-------------|----------------|------------|:-------:|
| timestamp   | Unix timestamp | seconds    |*current*|
| temperature | Temperature    | celsius    |    0    |
| humidity    | Humidity       | percentage |    0    |
| pressure    | Pressure       | hPa        |    0    |
| gas         | Gas            | kOhm       |    0    |

**Response**:

```json
HTTP/1.1 201 Created
{
  "timestamp": 1586532636,
  "temperature": 12,
  "pressure": 1001.5,
  "humidity": 41,
  "gas": 14
}

```

## Common errors

Missing `Authorization` header:

```json
HTTP/1.1 400 Bad Request
{
  "message": "missing key in request header"
}
```

Wrong token:

```json
HTTP/1.1 401 Unauthorized
{
  "message": "Unauthorized"
}
```


Wrong data format:

```json
HTTP/1.1 500 Internal Server Error
{
  "message": "Internal Server Errror"
}
```
