# sortarrays

## About The Project

The project consist of two apps:
* A gateway app that sends randomly sized arrays filled with random numbers
* A sort array app that sorts the received arrays and store them into a mysql database

## Setup the Database

Percona Server for MySQL.

Start a Percona MySQL docker container from the root directory of the locatee repository:

```
docker-compose up -d database
```

The root password for this database is `password`.
There is already a database created on this server - `array`.
The database comes with a dedicated user called `array` and password `array`.

## Authors

Mastro Staffy & Mastro Pinkie
