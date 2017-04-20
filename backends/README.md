# Backend APIs
Backend REST APIs for each store. Each backend provides a similar interface for performing CRUD operations on the resources required to place and retrieve orders.


## Learnings during the project

### 1.MongoDB Setup
a. Used link https://gist.github.com/leommoore/309de7c0042ed697ee84 for setup
b. Few updates to the steps mentioned in above link : do rs.initiate() only for primary node and change 'mongod.conf file for all - 
        add following line and take care of the indentation and spacing between the parameters
        replication:
         replSetName: "rs0"
c. ssh using 'ubuntu' and not 'ec2-user'
        
        ssh -i <key file name> ubuntu@ip
