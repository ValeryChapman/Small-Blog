## Contents
- [Structures](#structures)
  - [test](#structures_test)
- [Routes](#routes)
  - [Get all articles](#routes_all_articles)
  - [Get an article by Id](#routes_article)
  - [Create an article](#routes_create_article)
  - [Edit an article](#routes_edit_article)
  - [Delete an article](#routes_delete_article)


## Routes <div id="routes"></div>

### Get all articles <div id="routes_all_articles"></div>
```
Method: GET
```
```
/api/articles
```

#### Good response example:
```json5
{
    "data": {
        "data": [
            {
                "id": 1,
                "title": "Why use Golang for your project?",
                "text": "...",
                "views": 0,
                "created_at": "2021-08-17T23:30:43.95231+03:00"
            }
        ]
    },
    "status": "OK"
}
```

### Get an article by Id <div id="routes_article"></div>
```
Method: GET
```
```
/api/articles/:Id
```

#### Good response example:
```json5
{
    "data": {
        "id": 1,
        "title": "Why use Golang for your project?",
        "text": "...",
        "views": 0,
        "created_at": "2021-08-17T23:30:43.95231+03:00"
    },
    "status": "OK"
}
```

### Create an article (🔒) <div id="routes_create_article"></div>
```
Method: POST
```
```
Authorization: Bearer YOUR_TOKEN
```
```
/api/articles
```

#### Request data:
| Name | Type | Description | Example | Required |
|--|--|--|--|:--:|
| title | string | Article title | Why use Golang for your project? | + |
| text | string | Article text | Go is an open source progr... | - |

#### Good response example:
```json5
{
    "data": {
        "id": 1
    },
    "status": "OK"
}
```

### Edit an article (🔒) <div id="routes_edit_article"></div>
```
Method: PUT
```
```
Authorization: Bearer YOUR_TOKEN
```
```
/api/articles/:Id
```

#### Request data:
| Name | Type | Description | Example | Required |
|--|--|--|--|:--:|
| title | string | Article title | Why use Golang for your project? | + |
| text | string | Article text | Go is an open source progr... | - |

#### Good response example:
```json5
{
    "status": "ok"
}
```


### Delete an article (🔒) <div id="routes_delete_article"></div>
```
Method: DELETE
```
```
Authorization: Bearer YOUR_TOKEN
```
```
/api/articles/:Id
```

#### Good response example:
```json5
{
    "status": "ok"
}
```
