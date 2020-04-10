# API

## Get status

Returns server version and source git commit from which it was built.

**Endpoint:**

```http
GET /
```

**Response:**

```json
HTTP/1.1 200 OK
{
  "version": "refs/heads/master",
  "commit": "88694ddfb91153cd470bec20d23aee360feccdf7"
}
```

## Get latest sensors readings

Returns newest sensors readings found in the database.

**Endpoint:**

```http
GET /sensors
```

**Response:**

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

## Create new sensors readings

Create a new entry in the database with given sensors readings.

**Endpoint:**

```http
POST /sensors
```

**Headers:**

| Name           | Value             |
|----------------|-------------------|
| Content-Type   | application/json  |
| Authentication | Bearer **$TOKEN** |

**Data:**

| Parameter   | Description    | Unit       | Default |
|-------------|----------------|------------|:-------:|
| timestamp   | Unix timestamp | seconds    |*current*|
| temperature | Temperature    | celsius    |    0    |
| humidity    | Humidity       | percentage |    0    |
| pressure    | Pressure       | hPa        |    0    |
| gas         | Gas            | kOhm       |    0    |

**Response:**

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
