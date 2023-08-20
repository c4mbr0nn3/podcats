<div align="center">
  <img src="https://raw.githubusercontent.com/c4mbr0nn3/podcats/main/frontend/public/android-chrome-192x192.png" />
</div>



# Podcats

A podcast manager with cats, inspired by [Podgrab](https://github.com/akhilrex/podgrab).

I just wanted to have some fun with Vue.js 3 and Vuetify 3 while learning some GoLang skill (still very basic, tho), that's why I started this project.

## Tech Stack

**Client:** Vue.js, Vuetify, howler.js, marked

**Server:** GoLang, gin-gonic, gorm, viper

**Database:** SQLite

## Features

- ✅ Reproduce podcasts (o rly?)
- ✅ Multi user support
- ✅ Import podcasts RSS feed and store them in SQLite database
- ✅ Cron job keeps podcasts updated
- ✅ Search imported podcasts from anywhere in the app by pressing `Ctrl+K`, with fuzzy search (fuse.js)
- ✅ Search and import new podcasts from iTunes API

## Roadmap

- Import/Export OPML
- Notification for cron job
- Mobile UI
- Typescript (I had to learn it sooner or later)

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

### Credentials

Default credentials are:
- **Username:** root
- **Password:** changeme

It is going to ask you to change the password on the first login.

## Development

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
  make dev
```

Podcast is going to be exposed at `http://localhost:5173/`.

Client side supports hot reloading, so you can edit the code and see the changes in real time.

### Credentials

Default credentials are:
- **Username:** root
- **Password:** changeme

It is going to ask you to change the password on the first login.
