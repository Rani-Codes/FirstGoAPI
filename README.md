# First Go API

### Building my first REST-style API in Go using only the standard library.

### Why?
To learn the core mechanics of Go without the magic of a framework.

### Learned Anything New?
- Lesson 1: In order to send errors in JSON to the client you must update the header to the status code AND write the JSON error body (using a ErrorResponse struct)
    - Why? Because the header says quickly what happened and the body response should say why or provide additional error details.
- Lesson 2: Include struct tags for JSON APIs to match industry standard.
    - JSON keys must be lowercase

