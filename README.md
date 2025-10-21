# STRING ANALYZER SERVICE

A RESTful API service that analyses strings and stores their computed properties

### Getting Started
```bash
git clone https://github.com/Emeey-Lanr/hng_BE_Stage1.git
```

```bash
cd hng_BE_Stage1
```

Endpoints 
## 1 Create and Analyse String
### POST /strings
### Content-Type: Application/json

Request Body
```json
   {"value":"String to analyse"}
```

Success Response (201) Created
```json
  {
    "id": "e72b58fd0b7e8b49ddf1e771f427108e37c2ae1195cb5eb9b4e11ced28d6075c",
    "value": "String to analyse",
    "properties": {
        "length": 17,
        "is_palindrome": false,
        "unique_characters": 12,
        "word_count": 3,
        "sha256_hash": "e72b58fd0b7e8b49ddf1e771f427108e37c2ae1195cb5eb9b4e11ced28d6075c",
        "character_frequency": {
            " ": 2,
            "a": 2,
            "e": 1,
            "g": 1,
            "i": 1,
            "l": 1,
            "n": 2,
            "o": 1,
            "r": 1,
            "s": 2,
            "t": 2,
            "y": 1
        }
    },
    "created_at": "2025-10-21T09:08:19Z"
}
```


## 2GetSpecific String
### POST /strings/{your value}

Success Response (201) Ok

```json
{
    "id": "e72b58fd0b7e8b49ddf1e771f427108e37c2ae1195cb5eb9b4e11ced28d6075c",
    "value": "String to analyse",
    "properties": {
        "length": 17,
        "is_palindrome": false,
        "unique_characters": 12,
        "word_count": 3,
        "sha256_hash": "e72b58fd0b7e8b49ddf1e771f427108e37c2ae1195cb5eb9b4e11ced28d6075c",
        "character_frequency": {
            " ": 2,
            "a": 2,
            "e": 1,
            "g": 1,
            "i": 1,
            "l": 1,
            "n": 2,
            "o": 1,
            "r": 1,
            "s": 2,
            "t": 2,
            "y": 1
        }
    },
    "created_at": "2025-10-21T09:34:28Z"
}
```



### 3



### 5 Delete Specific String
### DELETE /strings/{string_value}

Success Response (204 No Content)

