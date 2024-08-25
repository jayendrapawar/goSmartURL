
# Request

```
curl --location 'http://localhost:3000/api/v1/' \
--header 'Content-Type: application/json' \
--data '{
    "url": "https://www.youtube.com/watch?v=3ExDEeSnyvE"
}'
```

# Response

## Shorten URL Response

```
{
    "url": "https://www.youtube.com/watch?v=3ExDEeSnyvE",
    "short": "localhost:3000/f0c899",
    "expiry": 24,
    "rate_limit": 9,
    "rate_limit_reset": 30
}
```

## API Limit Quota exceeded for specific time : 30 Minutes

```
{
    "error": "Rate limit exceeded",
    "rate_limit_rest": 29
}
```