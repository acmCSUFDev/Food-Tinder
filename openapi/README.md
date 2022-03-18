<!-- Generator: Widdershins v4.0.1 -->

<h1 id="openapi-schema-for-food-tinder">OpenAPI schema for Food-Tinder v0.0.1</h1>

> Scroll down for code samples, example requests and responses. Select a language for code samples from the tabs above or the mobile navigation menu.

Base URLs:

* <a href="https://{hostname}/api/v0">https://{hostname}/api/v0</a>

    * **hostname** - The domain of the API backend server Default: localhost

# Authentication

- HTTP Authentication, scheme: bearer 

<h1 id="openapi-schema-for-food-tinder-default">Default</h1>

## login

<a id="opIdlogin"></a>

`POST /login`

*Log in using username and password. A 401 is returned if the information is incorrect.*

<h3 id="login-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|username|query|string|true|none|
|password|query|string(password)|true|none|

> Example responses

> 200 Response

```json
{
  "user_id": "953809515621527562",
  "token": "WlvPXdNuyfttl8eSV67hkbsX51wLURzT",
  "expiry": "2019-08-24T14:15:22Z",
  "metadata": {
    "user_agent": "curl/7.64.1"
  }
}
```

<h3 id="login-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Logged in successfully.|[Session](#schemasession)|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|401 Unauthorized returned when the user enters the wrong username and password combination.|[Error](#schemaerror)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|Unexpected server error.|[Error](#schemaerror)|

<aside class="success">
This operation does not require authentication
</aside>

## register

<a id="opIdregister"></a>

`POST /register`

*Register using username and password*

<h3 id="register-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|username|query|string|true|none|
|password|query|string(password)|true|none|

> Example responses

> 200 Response

```json
{
  "user_id": "953809515621527562",
  "token": "WlvPXdNuyfttl8eSV67hkbsX51wLURzT",
  "expiry": "2019-08-24T14:15:22Z",
  "metadata": {
    "user_agent": "curl/7.64.1"
  }
}
```

<h3 id="register-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Registered successfully.|[Session](#schemasession)|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Form error.|[FormError](#schemaformerror)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|Unexpected server error.|[Error](#schemaerror)|

<aside class="success">
This operation does not require authentication
</aside>

## getSelf

<a id="opIdgetSelf"></a>

`GET /users/self`

*Get the current user*

> Example responses

> 200 Response

```json
null
```

<h3 id="getself-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|[Self](#schemaself)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|Unexpected server error.|[Error](#schemaerror)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
BearerAuth
</aside>

## getUser

<a id="opIdgetUser"></a>

`GET /users/{id}`

*Get a user by their ID*

<h3 id="getuser-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|id|path|[ID](#schemaid)|true|The ID of the user.|
|password|query|string(password)|true|none|

> Example responses

> 200 Response

```json
{
  "id": "953809515621527562",
  "name": "food-tinder-user",
  "avatar": "ypeBEsobvcr6wjGzmiPcTaeG7_gUfE5yuYB3ha_uSLs=",
  "bio": "Hello, world."
}
```

<h3 id="getuser-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|[User](#schemauser)|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Form error.|[FormError](#schemaformerror)|
|404|[Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)|Resource indicated by a parameter (if any) isn't found.|[FormError](#schemaformerror)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|Unexpected server error.|[Error](#schemaerror)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
BearerAuth
</aside>

## getNextPosts

<a id="opIdgetNextPosts"></a>

`GET /posts`

*Get the next batch of posts*

<h3 id="getnextposts-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|prev_id|query|[ID](#schemaid)|false|The ID to start the pagination from, or empty to start from top.|

> Example responses

> 200 Response

```json
[
  {
    "id": "953809515621527562",
    "user_id": "953809515621527562",
    "cover_hash": "LEHV6nWB2yk8pyoJadR*.7kCMdnj",
    "images": [
      "ypeBEsobvcr6wjGzmiPcTaeG7_gUfE5yuYB3ha_uSLs="
    ],
    "description": "Salmon roll for $8.\n\nPretty cheap for me.",
    "tags": [
      "Salmon",
      "Sushi Rice"
    ],
    "location": "Fullerton, CA"
  }
]
```

<h3 id="getnextposts-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Form error.|[FormError](#schemaformerror)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|Unexpected server error.|[Error](#schemaerror)|

<h3 id="getnextposts-responseschema">Response Schema</h3>

Status Code **200**

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|*anonymous*|[[Post](#schemapost)]|false|none|none|
|» id|[ID](#schemaid)|true|none|Snowflake ID.|
|» user_id|[ID](#schemaid)|true|none|Snowflake ID.|
|» cover_hash|string|false|none|none|
|» images|[string]|true|none|none|
|» description|string|true|none|none|
|» tags|[string]|true|none|none|
|» location|string|false|none|Location is the location where the post was made.|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
BearerAuth
</aside>

## deletePost

<a id="opIddeletePost"></a>

`DELETE /posts/{id}`

*Delete the current user's posts by ID. A 401 is returned if the user tries to delete someone else's post.*

<h3 id="deletepost-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|id|path|[ID](#schemaid)|true|The post ID to delete.|

> Example responses

> 404 Response

```json
{
  "message": "missing username",
  "form_id": "username"
}
```

<h3 id="deletepost-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|All posts successfully deleted.|None|
|404|[Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)|Resource indicated by a parameter (if any) isn't found.|[FormError](#schemaformerror)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|Unexpected server error.|[Error](#schemaerror)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
BearerAuth
</aside>

## getLikedPosts

<a id="opIdgetLikedPosts"></a>

`GET /posts/liked`

*Get the list of posts liked by the user*

> Example responses

> 200 Response

```json
[
  {
    "id": "953809515621527562",
    "user_id": "953809515621527562",
    "cover_hash": "LEHV6nWB2yk8pyoJadR*.7kCMdnj",
    "images": [
      "ypeBEsobvcr6wjGzmiPcTaeG7_gUfE5yuYB3ha_uSLs="
    ],
    "description": "Salmon roll for $8.\n\nPretty cheap for me.",
    "tags": [
      "Salmon",
      "Sushi Rice"
    ],
    "location": "Fullerton, CA"
  }
]
```

<h3 id="getlikedposts-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|Unexpected server error.|[Error](#schemaerror)|

<h3 id="getlikedposts-responseschema">Response Schema</h3>

Status Code **200**

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|*anonymous*|[[Post](#schemapost)]|false|none|none|
|» id|[ID](#schemaid)|true|none|Snowflake ID.|
|» user_id|[ID](#schemaid)|true|none|Snowflake ID.|
|» cover_hash|string|false|none|none|
|» images|[string]|true|none|none|
|» description|string|true|none|none|
|» tags|[string]|true|none|none|
|» location|string|false|none|Location is the location where the post was made.|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
BearerAuth
</aside>

## getAsset

<a id="opIdgetAsset"></a>

`GET /assets/{id}`

*Get the file asset by the given ID. Note that assets are not separated by type; the user must assume the type from the context that the asset ID is from.*

<h3 id="getasset-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|id|path|string|true|The asset ID to fetch.|

> Example responses

> 200 Response

> 404 Response

```json
{
  "message": "missing username",
  "form_id": "username"
}
```

<h3 id="getasset-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|string|
|404|[Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)|Resource indicated by a parameter (if any) isn't found.|[FormError](#schemaformerror)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|Unexpected server error.|[Error](#schemaerror)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
BearerAuth
</aside>

# Schemas

<h2 id="tocS_Error">Error</h2>
<!-- backwards compatibility -->
<a id="schemaerror"></a>
<a id="schema_Error"></a>
<a id="tocSerror"></a>
<a id="tocserror"></a>

```json
{
  "message": "server blew up"
}

```

Error object returned on any error.

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|message|string|true|none|none|

<h2 id="tocS_FormError">FormError</h2>
<!-- backwards compatibility -->
<a id="schemaformerror"></a>
<a id="schema_FormError"></a>
<a id="tocSformerror"></a>
<a id="tocsformerror"></a>

```json
{
  "message": "missing username",
  "form_id": "username"
}

```

Error object returned on a form error.

### Properties

allOf

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|*anonymous*|[Error](#schemaerror)|false|none|Error object returned on any error.|

and

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|*anonymous*|object|false|none|none|
|» form_id|string|false|none|none|

<h2 id="tocS_ID">ID</h2>
<!-- backwards compatibility -->
<a id="schemaid"></a>
<a id="schema_ID"></a>
<a id="tocSid"></a>
<a id="tocsid"></a>

```json
"953809515621527562"

```

Snowflake ID.

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|*anonymous*|string|false|none|Snowflake ID.|

<h2 id="tocS_LoginMetadata">LoginMetadata</h2>
<!-- backwards compatibility -->
<a id="schemaloginmetadata"></a>
<a id="schema_LoginMetadata"></a>
<a id="tocSloginmetadata"></a>
<a id="tocsloginmetadata"></a>

```json
{
  "user_agent": "curl/7.64.1"
}

```

Optional metadata included on login.

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|user_agent|string|false|none|The User-Agent used for logging in.|

<h2 id="tocS_Session">Session</h2>
<!-- backwards compatibility -->
<a id="schemasession"></a>
<a id="schema_Session"></a>
<a id="tocSsession"></a>
<a id="tocssession"></a>

```json
{
  "user_id": "953809515621527562",
  "token": "WlvPXdNuyfttl8eSV67hkbsX51wLURzT",
  "expiry": "2019-08-24T14:15:22Z",
  "metadata": {
    "user_agent": "curl/7.64.1"
  }
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|user_id|[ID](#schemaid)|true|none|Snowflake ID.|
|token|string|true|none|none|
|expiry|string(date-time)|true|none|none|
|metadata|[LoginMetadata](#schemaloginmetadata)|true|none|Optional metadata included on login.|

<h2 id="tocS_Self">Self</h2>
<!-- backwards compatibility -->
<a id="schemaself"></a>
<a id="schema_Self"></a>
<a id="tocSself"></a>
<a id="tocsself"></a>

```json
null

```

### Properties

allOf

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|*anonymous*|[User](#schemauser)|false|none|none|

and

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|*anonymous*|[FoodPreferences](#schemafoodpreferences)|false|none|none|

and

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|*anonymous*|object|false|none|none|
|» birthday|string(date)|true|none|none|

<h2 id="tocS_User">User</h2>
<!-- backwards compatibility -->
<a id="schemauser"></a>
<a id="schema_User"></a>
<a id="tocSuser"></a>
<a id="tocsuser"></a>

```json
{
  "id": "953809515621527562",
  "name": "food-tinder-user",
  "avatar": "ypeBEsobvcr6wjGzmiPcTaeG7_gUfE5yuYB3ha_uSLs=",
  "bio": "Hello, world."
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|id|[ID](#schemaid)|true|none|Snowflake ID.|
|name|string|true|none|none|
|avatar|string|true|none|none|
|bio|string|false|none|none|

<h2 id="tocS_FoodPreferences">FoodPreferences</h2>
<!-- backwards compatibility -->
<a id="schemafoodpreferences"></a>
<a id="schema_FoodPreferences"></a>
<a id="tocSfoodpreferences"></a>
<a id="tocsfoodpreferences"></a>

```json
{
  "likes": [
    "Rice",
    "Fish"
  ],
  "prefers": {
    "Rice": [
      "Sushi Rice"
    ],
    "Fish": [
      "Ahi Tuna",
      "Salmon"
    ]
  }
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|likes|[string]|true|none|none|
|prefers|object|true|none|none|
|» **additionalProperties**|[string]|false|none|none|

<h2 id="tocS_UserLikedPosts">UserLikedPosts</h2>
<!-- backwards compatibility -->
<a id="schemauserlikedposts"></a>
<a id="schema_UserLikedPosts"></a>
<a id="tocSuserlikedposts"></a>
<a id="tocsuserlikedposts"></a>

```json
{
  "posts": {
    "953809515621527562": "2018-08-24T14:15:22Z",
    "953809438236627014": "2018-08-24T14:15:48Z"
  },
  "remaining": 3,
  "expires": "2019-08-24T14:15:22Z"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|posts|object|true|none|Posts maps post IDs to the time that the user liked.|
|» **additionalProperties**|string(date-time)|false|none|none|
|remaining|number|true|none|Remaining is the number of likes allowed by the user until the Expires timestamp.|
|expires|string(date-time)|true|none|Expires is the time that the rate limiter (the Remaining field) replenishes.|

<h2 id="tocS_Post">Post</h2>
<!-- backwards compatibility -->
<a id="schemapost"></a>
<a id="schema_Post"></a>
<a id="tocSpost"></a>
<a id="tocspost"></a>

```json
{
  "id": "953809515621527562",
  "user_id": "953809515621527562",
  "cover_hash": "LEHV6nWB2yk8pyoJadR*.7kCMdnj",
  "images": [
    "ypeBEsobvcr6wjGzmiPcTaeG7_gUfE5yuYB3ha_uSLs="
  ],
  "description": "Salmon roll for $8.\n\nPretty cheap for me.",
  "tags": [
    "Salmon",
    "Sushi Rice"
  ],
  "location": "Fullerton, CA"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|id|[ID](#schemaid)|true|none|Snowflake ID.|
|user_id|[ID](#schemaid)|true|none|Snowflake ID.|
|cover_hash|string|false|none|none|
|images|[string]|true|none|none|
|description|string|true|none|none|
|tags|[string]|true|none|none|
|location|string|false|none|Location is the location where the post was made.|

