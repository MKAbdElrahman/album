## Application architecture

The architecture follows a typical "onion model" where each layer doesn't know about the layer above it, and each layer is responsible for a specific thing.

```mermaid
graph LR
    handler[HTTP handler] -- uses --> services[Services]
    services -- use --> db[Database access code]
    db -- uses --> dynamodb[(DynamoDB)]
    handler -- renders --> components[Components]
```

* HTTP Handler
  * Processes HTTP requests
  * Does not contain application logic itself
  * Uses `services` that carry out application logic
  * Takes the responses from `services` and uses `components` to render HTML
  * Creates HTTP responses
* Services
  * Carries out application logic such as orchestrating API calls, or making database calls
  * Does not do anything related to HTML or HTTP
  * Is not aware of the specifics of database calls
* Database access code
  * Handles database activity such as inserting and querying records
  * Ensures that the database representation (`records`) doesn't leak to the service layer


  ## Dev Environment

  - VSCode as IDE
  - Templ for static rendering
  - Taskfile for automation
  - wgo for live reloading