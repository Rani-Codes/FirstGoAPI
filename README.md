# First Go API

### Building my first REST-style API in Go using only the standard library.

### Why?
To learn the core mechanics of Go without the magic of a framework.

### Learned Anything New?
- Lesson 1: In order to send errors in JSON to the client you must update the header to the status code AND write the JSON error body (using a ErrorResponse struct)
    - Why? Because the header says quickly what happened and the body response should say why or provide additional error details.
- Lesson 2: Include struct tags for JSON APIs to match industry standard.
    - JSON keys must be lowercase

### Final Remarks
This marks the completion of my first REST API project using Go's standard ```net/http``` library. It was a great learning experience that demonstrated the power and simplicity of the standard library. At the same time, it also revealed some limitations that make adopting a lightweight framework like **Chi** appealing for future projects.

- Why use a framework?
    - Multiple methods supported on a single path. With chi I can simply write ```r.Delete("/items/{id}", DeleteItemById) ``` and have it confirm both my HTTP Method and a query paramter.
        - Eliminates the need for workarounds like using /createItems, /deleteItems/, or writing switch cases manually.
    - No more method checks. I did one here for POST but ***chi*** handles that internally by routing based on both the path and method — no extra logic required..
    - Can finally move things outside of main.go
        - Better project organization through middleware support, route grouping, and controller separation. Good for more complex projects
    
