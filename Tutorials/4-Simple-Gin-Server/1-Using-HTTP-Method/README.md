Using HTTP method
Gin provides methods that map directly to HTTP verbs, making it straightforward to build RESTful APIs. Each method registers a route that responds only to the corresponding HTTP request type:

Method	Typical REST Usage
GET	Retrieve a resource
POST	Create a new resource
PUT	Replace an existing resource
PATCH	Partially update an existing resource
DELETE	Remove a resource
HEAD	Same as GET but without a body
OPTIONS	Describe communication options