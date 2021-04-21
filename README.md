# API for searching the top 1000 movies from IMDB

## How to run
There is a docker-compose file that already configures both crawler and storage API for you. Just run:

```shell
$ docker-compose up
```

Storage API uses port 8000. If it is in use, change it by editing the docker-compose file.

## Assumptions made

* When crawling IMDB page, I just got the information from the movie listing, I didn't open each page to get all the information about each movie. I did this because the listing page had most of the information needed (movie name, director and two most important cast members). This reduces the number of terms that you can use to identify a movie and can impact the API results if you search for cast members that are not the main stars of the movie.

* Search terms must be a single word. This happens because when inserting a movie into the storage, the API creates a set of keys for that movie. These keys are one word each and are used for searching inside a map. If you enter "jurassic park" as a term, no results will be returned by the API. But ["jurassic", "park"] will work fine. By using words as keys for each movie, I can use a O(1) search to find all movies that contain that key. If the API had to support multiple words or incomplete words as inputs, I would have to build tokens (just like pg 12's `to_tsvector` does) for each movie.

## What would I do if I had more time

* Create more test scenarios to cover not only the happy path of the application;
* Parse the IMDB movie page to extract more information about it, such as genre, year and the complete cast list;
* Write a performance test to check how many requests per second this service can handle;