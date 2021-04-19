# Storage API

This is responsible for storing and searching movies based on keywords.

## How to run

If you just want to run the code, execute:
```shell
make run
```

However, if you want to build a binary and execute it, run:
```shell
make build
./bin/storage
```

You also can run this as a docker container:
```shell
docker build -t storage-api .
docker run storage-api
```

## How it works
When a new movie is stored, all its information are split into words and those words are used as 
keys that point to that movie instance. These keys are stored in a map structure. Each key is
associates with a bucket of movies.

When a key is used to search movies, this API will access the bucket of movies related to that key
and return them. In case of multiple keys, each bucket will be retrieved and movies that appear
in all buckets at the same time are returned to the user.

**Complexity**
* Space complexity: O(k * n) (where `k` is the average number of keys in a movie)
* Time complexity: O(n)