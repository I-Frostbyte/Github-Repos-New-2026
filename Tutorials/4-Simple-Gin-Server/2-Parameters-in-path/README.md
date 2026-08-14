Parameters in path
Gin supports two types of path parameters that let you capture values directly from the URL:

:name — matches a single path segment. For example, /user/:name matches /user/john but does not match /user/ or /user.
*action — matches everything after the prefix, including slashes. For example, /user/:name/*action matches /user/john/send and /user/john/. The captured value includes the leading /.
Use c.Param("name") to retrieve the value of a path parameter inside your handler.