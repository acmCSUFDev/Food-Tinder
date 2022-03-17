<!-- Generator: Widdershins v4.0.1 -->

<h1 id="openapi-schema-for-food-tinder">OpenAPI schema for Food-Tinder v0.0.1</h1>

> Scroll down for code samples, example requests and responses. Select a language for code samples from the tabs above or the mobile navigation menu.

<h1 id="openapi-schema-for-food-tinder-default">Default</h1>

## Log in using username and password

<a id="opIdlogin"></a>

`POST /login`

<h3 id="log-in-using-username-and-password-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|username|query|string|true|none|
|password|query|string(password)|true|none|

> Example responses

> 200 Response

```json
{
  "user_id": "string",
  "token": "string",
  "expiry": "2019-08-24T14:15:22Z",
  "metadata": {
    "user_agent": "string"
  }
}
```

<h3 id="log-in-using-username-and-password-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Logged in successfully|[Session](#schemasession)|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|User error|[Error](#schemaerror)|

<aside class="success">
This operation does not require authentication
</aside>

## Register using username and password

<a id="opIdregister"></a>

`POST /register`

<h3 id="register-using-username-and-password-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|username|query|string|true|none|
|password|query|string(password)|true|none|

> Example responses

> 200 Response

```json
{
  "user_id": "string",
  "token": "string",
  "expiry": "2019-08-24T14:15:22Z",
  "metadata": {
    "user_agent": "string"
  }
}
```

<h3 id="register-using-username-and-password-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Registered successfully|[Session](#schemasession)|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|User error|[Error](#schemaerror)|

<aside class="success">
This operation does not require authentication
</aside>

## Get the current user

`GET /users/self`

> Example responses

> 200 Response

```json
null
```

<h3 id="get-the-current-user-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|[Self](#schemaself)|

<aside class="success">
This operation does not require authentication
</aside>

## Get a user by their ID

`GET /users/{id}`

<h3 id="get-a-user-by-their-id-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|id|path|[ID](#schemaid)|true|The ID of the user|
|password|query|string(password)|true|none|

> Example responses

> 200 Response

```json
{
  "id": "string",
  "name": "string",
  "avatar": "string",
  "bio": "string"
}
```

<h3 id="get-a-user-by-their-id-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|[User](#schemauser)|

<aside class="success">
This operation does not require authentication
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
  "message": "string"
}

```

Error object returned on any error

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|message|string|true|none|none|

<h2 id="tocS_ID">ID</h2>
<!-- backwards compatibility -->
<a id="schemaid"></a>
<a id="schema_ID"></a>
<a id="tocSid"></a>
<a id="tocsid"></a>

```json
"string"

```

Snowflake ID

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|*anonymous*|string|false|none|Snowflake ID|

<h2 id="tocS_LoginMetadata">LoginMetadata</h2>
<!-- backwards compatibility -->
<a id="schemaloginmetadata"></a>
<a id="schema_LoginMetadata"></a>
<a id="tocSloginmetadata"></a>
<a id="tocsloginmetadata"></a>

```json
{
  "user_agent": "string"
}

```

Optional metadata included on login

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|user_agent|string|false|none|The User-Agent used for logging in|

<h2 id="tocS_Session">Session</h2>
<!-- backwards compatibility -->
<a id="schemasession"></a>
<a id="schema_Session"></a>
<a id="tocSsession"></a>
<a id="tocssession"></a>

```json
{
  "user_id": "string",
  "token": "string",
  "expiry": "2019-08-24T14:15:22Z",
  "metadata": {
    "user_agent": "string"
  }
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|user_id|[ID](#schemaid)|true|none|Snowflake ID|
|token|string|true|none|none|
|expiry|string(date-time)|true|none|none|
|metadata|[LoginMetadata](#schemaloginmetadata)|false|none|Optional metadata included on login|

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
  "id": "string",
  "name": "string",
  "avatar": "string",
  "bio": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|id|[ID](#schemaid)|true|none|Snowflake ID|
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
    "string"
  ],
  "prefers": {
    "property1": "string",
    "property2": "string"
  }
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|likes|[string]|true|none|none|
|prefers|object|true|none|none|
|» **additionalProperties**|string|false|none|none|

<h2 id="tocS_UserLikedPosts">UserLikedPosts</h2>
<!-- backwards compatibility -->
<a id="schemauserlikedposts"></a>
<a id="schema_UserLikedPosts"></a>
<a id="tocSuserlikedposts"></a>
<a id="tocsuserlikedposts"></a>

```json
{
  "posts": {
    "property1": "2019-08-24T14:15:22Z",
    "property2": "2019-08-24T14:15:22Z"
  },
  "remaining": 0,
  "expires": "2019-08-24T14:15:22Z"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|posts|object|true|none|none|
|» **additionalProperties**|string(date-time)|false|none|none|
|remaining|number|true|none|none|
|expires|string(date-time)|true|none|none|

<h2 id="tocS_Post">Post</h2>
<!-- backwards compatibility -->
<a id="schemapost"></a>
<a id="schema_Post"></a>
<a id="tocSpost"></a>
<a id="tocspost"></a>

```json
{
  "id": "string",
  "user_id": "string",
  "cover_hash": "string",
  "images": [
    "string"
  ],
  "description": "string",
  "tags": [
    "string"
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|id|[ID](#schemaid)|true|none|Snowflake ID|
|user_id|[ID](#schemaid)|true|none|Snowflake ID|
|cover_hash|string|false|none|none|
|images|[string]|true|none|none|
|description|string|true|none|none|
|tags|[string]|true|none|none|

