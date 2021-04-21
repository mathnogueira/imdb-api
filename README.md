# API for searching the top 1000 movies from IMDB

## How to run
There is a docker-compose file that already configures both crawler and storage API for you. Just run:

```shell
$ docker-compose up
```

Storage API uses port 8000. If it is in use, change it by editing the docker-compose file.

## Assumptions made

* When crawling IMDB page, I just got the information from the listing itself. I didn't open each page to get all the information about a movie. I did this because the listing page had most of the information needed already (movie name, director and two most important cast members). By doing this, if you search for actors that didn't star in a movie, probably no result will be returned.
* Search terms must be a single word. This happens because when inserting a movie into the storage, I create keys for that movie. These keys are one word each and are used for searching inside a map. If you pass a term "jurassic park", no results will be returned by the API.

## What would I do if I had more time

* Create more test scenarios to cover not only the happy path of the application;
* Parse the IMDB movie page to extract more information about it, such as genre, year and the complete cast list;