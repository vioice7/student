# Simple Student Managing App 
## RESTApi Server, HTTP Server, MySQL Server
---
* This is a simple management CRUD app for a university that uses Go RESTApi Server, Go HTTP Server and MySql server
* Each layer is structured so it can communicate in a simple manner:
* MYSQL Server <-> RESTAPI Server <-> HTTP Server
*
* The server adress (default):
* localhost:8080
* 
* The schematic:
* mysql functions (DATABASE) <-> rest functions (API) <-> html show files functions (HTML) <-> templates (HTML)