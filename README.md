# eld_go_company_service



# running in local 
1. Create docker network
```
docker network create eld-network
```

2. Run database 
The database can be run using the this [repo](https://gitlab.udevs.io/eld/local_environment).

3. Create your database 
```
docker exec -it database sh

psql -U postgres

create database your_database_name
```

4. Run service
First make sure service docker compose files is correctly configured. 
Thing that should be taken into account: 
> credentials
> ports (just changing outer port is enough the inner ports can remain the same)

Run the service with the following code
```
docker compose up -d --build
```

5. Stop service
```
docker compose down
```


# Needed doc's
- [db diagram](https://dbdiagram.io/d/eld_company_service-65421e367d8bbd646543b45d)
- []



 1. Bug fixes 
 List:
  1. Units:
   >export :back ⚠️ 
 3. Drivers: 
   > export :back :front 
