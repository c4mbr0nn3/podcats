<div align="center">
  <img src="https://raw.githubusercontent.com/c4mbr0nn3/podcats/main/frontend/public/android-chrome-192x192.png" />
</div>



# Podcats

A podcast manager with cats here and there, inspired by [Podgrab](https://github.com/akhilrex/podgrab).

## Why?

I just wanted to have some fun with Vue.js 3 and Vuetify 3 while learning some GoLang skill (still very basic, tho), that's why I started this project.

## Tech Stack

**Client:** Vue.js, Vuetify, howler.js, marked

**Server:** GoLang, gin-gonic, gorm, viper

**Database:** SQLite

## Features

- ✅ Play podcast episodes (o rly?)
- ✅ Multi-user support
- ✅ Import podcast RSS feeds and store them in an SQLite database
- ✅ Search for and import new podcasts from the iTunes API
- ✅ Cron job checks for new episodes every 5 minutes
- ✅ In-app notifications for the cron job
- ✅ Search for imported podcasts from anywhere in the app by pressing `Ctrl+K`, with fuzzy search (using fuse.js)
- ✅ Mark episodes as played or unplayed
- ✅ Mark episodes as favorites
- ✅ Resume listening feature

## Roadmap

- [ ] Play entire podcasts
- [ ] "Listen Later" playlist
- [ ] Offline listening capability
- [ ] Categorize podcasts
- [ ] Import/Export using OPML
- [ ] Mobile UI design
- [ ] Light Theme
- [ ] Configurable interval for the new episode cron job

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
