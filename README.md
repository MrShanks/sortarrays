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

## Run SortArray App as a docker container

Make sure the database service is running before compiling the sort array app, tests need it to be in place.

Compile and build the app
```
./gradlew build
```

Create dependency folder and unpack the fat jar 
```
mkdir -p build/dependency && (cd build/dependency; jar -xf ../libs/*.jar)
```

Create the docker image by running:
```
docker build .
```
from the Dockerfile folder

## Authors

Mastro Staffy & Mastro Pinkie
