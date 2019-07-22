# Simple Markdown Wiki Application

A Golang based REST API with a Vue JS front-end to view and edit articles with markdown format. 

The solution should Dockerized into one container.

## Installation

This project requires that [Docker](https://www.docker.com/) is installed on your local system.

This project should be cloned into your [Go](https://golang.org/) GOPATH source directory, i.e. "go/src/md-wiki".

Once cloned, navigate to the project directory. To build the Docker container, execute:

``` bash

sudo docker build -t (YOUR CONTAINER ID) .

```

Once the container is built, run the container with the following:

``` bash

docker run -ti -p 8080:8080 (YOUR CONTAINER ID)

```

Make sure in both instances you replace (YOUR CONTAINER ID) with a proper container name.

## Usage

Once the Docker container is running, you may access the application at 127.0.0.1:8080 using a web browser or [Curl](https://github.com/curl/curl).

URLs served are:

/: Articles Home Page
/(article name): Article View Page (if exists)
/edit/(article name): Create/Update Article Page

Access to the REST API is also available on the server at:

/articles/: JSON array of all article names
/articles/(article name): JSON object of individual article (if exists)

The article view and preview features provide [Markdown](https://www.markdownguide.org/basic-syntax/) compliant display. 
Markdown text compatible with a ".md" file will be rendered accordingly.