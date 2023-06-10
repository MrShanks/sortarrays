# Sortarrays

## About The Project

This is a study project to learn how to design a microservices architecture.

The project consist of two apps:
* A gateway app that sends randomly sized arrays filled with random numbers
* A sort array app that sorts the received arrays and store them into a mysql database

The apps can be deployed either via helm or docker-compose for development purposes.
```
helm install <release-name> -n <release-name> k8s/
```
```
docker-compose up -d
```

## Run the Database only

Percona Server for MySQL.

Start a Percona MySQL docker container from the root directory of the locatee repository:

```
docker-compose up -d database
```

The service comes up with an empty database called array.
The database comes with a dedicated user called `array` and password `array`.

## Authors

Mastro Staffy & Mastro Pinkie
