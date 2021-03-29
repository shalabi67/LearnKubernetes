curl --request PUT http://localhost:8080/apis/subresource.com/v1/namespaces/default/statuses/status-subresource/status \
--header 'Content-Type: application/json' \
--data-raw '{
    "apiVersion": "subresource.com/v1",
    "kind": "StatusSubresource",
    "metadata": {
        "name": "status-subresource",
        "resourceVersion": "492033"
    },
    "status": {
        "phase":"updated"
    }
}'

