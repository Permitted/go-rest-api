
# Go REST API

REST API example for personal portfolio.


## Run Locally

Clone the project

```bash
  git clone https://github.com/Permitted/go-rest-api
```

Go to the project directory

```bash
  cd go-rest-api
```

Start the server

```bash
   docker-compose up
```


## Using Guideline

This example repository uses docker and https://taskfile.dev/ as a build tool. Build steps are mentioned below.

With Task File:
```bash
    task run
```
Without Task File:
```bash
    docker-compose up --build
```   
## Task File Commands

To run this project, you can use and add task file commands. Useable commands mentioned below.

Run: `task run`

Build: `task build`

Down: `task down`

Integration Test:`task integration-test`


## API Reference

POST, PUT and DELETE have JWT Authentication. Generate key from https://jwt.io/.  
Key: "missionimpossible"
#### Post Comment

```http
  POST /api/v1/comment/
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `comment` | `string` | Post the created comment   |

#### Get Comment

```http
  GET /api/v1/comment/${id}
```

| Parameter | Type     | Description                          |
| :-------- | :------- | :----------------------------------- |
| `id`      | `string` | **Required**. Id of comment to fetch |

#### UPDATE Comment

```http
  UPDATE /api/v1/comment/${id}
```

| Parameter | Type     | Description                           |
| :-------- | :------- | :------------------------------------ |
| `id`      | `string` | **Required**. Id of comment to update |


#### DELETE Comment

```http
  DELETE /api/v1/comment/${id}
```

| Parameter | Type     | Description                           |
| :-------- | :------- | :------------------------------------ |
| `id`      | `string` | **Required**. Id of comment to delete |


