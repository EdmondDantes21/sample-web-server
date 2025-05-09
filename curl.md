# test api calls
## POST request
```bash
curl -X POST \
  ‘http://localhost:8080/samples’ \
    -H ‘Content-Type: application/json’ \
    -d ‘{
    “name”: “Test4",
    “id”: “4",
    “title”: “New title by POST”
    }’
```
## PUT request
```bash
curl -X PUT \
  ‘http://localhost:8080/samples/Test1’ \
    -H ‘Content-Type: application/json’ \
    -d ‘{
    “name”: “Test1",
    “id”: “1",
    “title”: “New title by PUT”,
    }’
```
## DELETE request
```bash
curl -X DELETE \
  ‘http://localhost:8080/samples/Test1’ \
    -H ‘Content-Type: application/json’
```