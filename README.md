# TODO API

## Overview

An API to create, read, and update todo lists.

## Liberties taken

- I didn't implement a skip feature for the searchable GET lists request, just a limit parameter

## Structure

In the tradition of idiomatic Go, orthogonal concerns for REST APIs will be split into separate packages. The service will consist of:

- **Routers**: for mapping URL paths and query strings to the functions that will do the real work of writing responses,
- **Handlers**: to do the real work of writing responses,
- **Data models**: internal representations of the types of domain data we will work with,
- **Persistence**: some kind of database to prevent the domain data from being destroyed if the service goes down, and
- **A dependency container**: to pass service-wide dependencies like database connections to the bits of code that need it.

## Next Steps

Some ideas to expand on the service:

- **Logging**: This could be facilitated by a layer of middleware before the router to record basic information about each request.
- **A concept of users**: The user could authenticate with the service, and would only have permission to interact with their own todo lists.


## Live instance
Live link available soon!
