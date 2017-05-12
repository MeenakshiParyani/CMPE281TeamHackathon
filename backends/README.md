# Backend APIs
Backend REST APIs for each store. Each backend provides a similar interface for performing CRUD operations on the resources required to place and retrieve orders.

```
docker rm -f go-restbucks; docker run -d --name go-restbucks -p 3000:5000 -e DB_URL="mongodb://mongo:27017/cmpe281" --link mongo:mongo ttpham0111/go-restbucks; docker logs -f go-restbucksgo-restbucks
```