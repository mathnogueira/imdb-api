# WebCrawler
This component is responsible for retrieving information from IMDB and send it to the Storage API to be indexed.

## Configuration

To run this component, you just need to configure the `Storage API URL` by setting the environment variable `CRAWLER_STORAGE_URL`.

```shell
export CRAWLER_STORAGE_URL=http://localhost:9001
./bin/webcrawler
```

You can run it on docker as well:

```shell
docker build -t webcrawler .
docker run -e CRAWLER_STORAGE_URL=http://localhost:9001 webcrawler
```