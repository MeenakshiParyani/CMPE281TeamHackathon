# Portal
Frontend application for placing orders to different stores (tenants. All requests are routed to the gateway and will be redirected to the correct backend.

## Contributors
- Meenakshi Paryani
- Nilam Pratim Deka
- Tuan Pham
- Varsha Kankariya

```
docker rm -f ng-restbucks; docker run --name ng-restbucks -d -p 3001:3000 --link go-restbucks:go-restbucks -e GATEWAY_URL=http://go-restbucks:5000 ttpham0111/ng-restbucks; docker logs -f ng-restbucks
```