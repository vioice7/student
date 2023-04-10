# Simple University Managing App 
## RESTApi Server, HTTP Server, MySQL Server
---
* This is a simple management CRUD app for a university that uses Go RESTApi Server, Go HTTP Server and MySql Server
* Each layer is structured so it can communicate in a simple manner:
* MYSQL Server <-> RESTAPI Server <-> HTTP Server
*
* The server adress (default):
* localhost:8080
* Config files:
* configHtmlServer.json (adress of the html server)
* database/configuration/config.json (adress of the mysql server)
*
* The schematic:
* mysql functions (DATABASE) <-> rest functions (API) <-> html show files functions (HTML) <-> templates (HTML)
