# TODO API

## Overview

An API to create, read, and update todo lists.

## Structure

In the tradition of idiomatic Go, orthogonal concerns for REST APIs will be split into separate packages. The service will consist of:

- **Routers**: for mapping URL paths and query strings to the functions that will do the real work of writing responses,
- **Handlers**: to do the real work of writing responses,
- **Data Models**: internal representations of the types of domain data we will work with,
- **Persistence**: some kind of database to prevent the domain data from being destroyed if the service goes down, and
- **A Dependency Container**: to pass service-wide dependencies like database connections to the bits of code that need it.

## Live instance
Live link available soon!
