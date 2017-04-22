FROM node:alpine

MAINTAINER cmpe281-cloud-spartans

WORKDIR /portal

# Install npm packages first to improve caching
COPY StarbucksWebPortal/package.json /portal/package.json
RUN npm install

COPY StarbucksWebPortal /portal

CMD ["npm", "start"]