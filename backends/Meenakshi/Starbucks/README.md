# Starbucks

ssh -i admin-key-pair-us-west-1.pem ec2-user@52.52.195.15 // mongo1
ec2-52-52-195-15.us-west-1.compute.amazonaws.com

ssh -i admin-key-pair-us-west-1.pem ec2-user@54.183.224.190 //mongo2 ec2-54-183-224-190.us-west-1.compute.amazonaws.com

ssh -i admin-key-pair-us-west-1.pem ec2-user@52.53.66.140 //arbiter
ec2-52-53-66-140.us-west-1.compute.amazonaws.com

sudo chkconfig --add disable-transparent-hugepages

https://docs.mongodb.com/manual/tutorial/install-mongodb-on-amazon/

sudo cat /var/log/mongodb/mongod.log

starbucks-replica-set



