<div align="center">
  <img src="https://raw.githubusercontent.com/c4mbr0nn3/podcats/main/frontend/public/android-chrome-192x192.png" />
</div>



# Podcats

A podcast player with cats, inspired by [Podgrab](https://github.com/akhilrex/podgrab).

I just wanted to have some fun with Vue.js 3 and Vuetify 3 while learning some GoLang skill (still very basic, tho), that's why I started this project.

## Tech Stack

**Client:** Vue.js, Vuetify, howler.js, marked

**Server:** GoLang, gin-gonic, gorm, viper

**Database:** SQLite

## Features

- Import podcasts RSS feed and store them in the database
- Cron job keeps podcasts updated
- Reproduce podcasts (o rly?)

## Roadmap

- Search and import new podcasts from iTunes API
- Import/Export OPML
- Multi user
- Mobile UI
- Typescript (I had to learn it sooner or later)
- E-mail cron job

## Deployment

To deploy this project you need [Docker](https://www.docker.com/) installed on your server.

### Using Docker Compose

Fastest way to get started with Podcats is to deploy it via docker-compose file, just download it from the repository or copy the following snippet:

```yaml
version: "3.8"

services:
  podcats:
    image: j1mm0/podcats:latest
    container_name: podcats
    restart: always
    ports:
      - 8000:8000
    volumes:
      - ${PWD}/podcats-db/:/go/src/podcats/db
```

### Using Docker

If you prefer to avoid docker-compose, you can retrieve the latest Podcats docker image with the following command:

```bash
  docker pull j1mm0/podcats:latest
```

Otherwise, if you want the very last code version, with higher probability of quirks, then clone the repository:

```bash
  git clone https://github.com/c4mbr0nn3/podcats.git
```

Go to the project directory:

```bash
  cd podcats
```

And finally build the docker image using the All-In-One Dockerfile:

```bash
  docker build --rm -f aio.Dockerfile --tag j1mm0/podcats:latest .
```

Once you have the image in your server, just run:

```bash
  docker container run -d -p 8000:8000 -v ${PWD}/podcats-db/:/go/src/podcats/db --name podcats j1mm0/podcats:latest
```

## Run Locally

To run the project locally you need to have [GoLang](https://go.dev/) and [Node.js](https://nodejs.org/en/) installed on your system.

Clone the project:

```bash
  git clone https://github.com/c4mbr0nn3/podcats.git
```

Go to the frontend directory and install frontend dependencies:

```bash
  cd podcats/frontend && npm ci
```

Go to backend directory and install backend dependencies:

```bash
  cd podcats/backend && go mod download
```

Run the server with the makefile:

```bash
  make run
```

Podcast is going to be exposed at `http://localhost:8000/`.
