# Backend APIs
Backend REST APIs for each store. Each backend provides a similar interface for performing CRUD operations on the resources required to place and retrieve orders.


## Learnings during the project

### 1.MongoDB Setup
* Used link https://gist.github.com/leommoore/309de7c0042ed697ee84 for setup
* Few updates to the steps mentioned in above link : do rs.initiate() only for primary node and change 'mongod.conf file for all nodes- add following line and take care of the indentation and spacing between the parameters
```javascript
replication:
 replSetName: "rs0"
 ```
    
* ssh using `ubuntu` and not `ec2-user`
        
        `ssh -i <key file name> ubuntu@ip`
