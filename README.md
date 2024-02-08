# Diet
A REST API that queries and modifies a PostgresQL database using Spring Boot. There is also a nix flake to get started. 

## Endpoints
	- [x] Food
		- [x] GET
			- [x] /calories - retrieve number of calories by name
                - Example request:
                ```
                {
                    "name": "Cheerios"
                }
                ```
                - Example response:
                ```
                {
                    "id":1,
                    "name":"Cheerios",
                    "units":1,
                    "calories":250
                }
                ```
		- [x] POST
			- [x] / - add a new food
				- Example request:
				```
				{
    					"name": "Plain Bagel",
    					"calories": 270,
    					"units": 1
				}
				```
			- [x] /calories - add calories to a new food
				- Example request:
				```
				{
    					"name": "Beef and Bean Burritos",
    					"calories": 290,
    					"units": 1
				}
				```
		- [x] PUT
			- [x] /calories - update a food's number of calories
				- Example request:
				```
				{
    					"name": "Beef and Bean Burritos",
    					"calories": 290,
    					"units": 1
				}
				```
	- [x] Meal
		- [x] GET
			- [x] / - retrieve the most recent meal by user
		- [x] POST
			- [x] / - create a new meal today
			- [x] /food - add a food to the latest meal by user
	- [x] Dieter
		- [x] GET
			- [x] / - get information on a user
			- [x] /id - retrive the ID from a user
			- [x] /calories - get the remaining number of calories for a user
			- [x] /caloriesToday - get the number of calories used today by a user
		- [x] POST
			- [x] / - create a new user
			- [x] /calories - set total number of calories by day for a user

## Setup

### Dependencies
```
	PostgresQL 15.4
	openJDK 17
	Maven
```

### Optional
```
	just (https://github.com/casey/just)
	Nix Package Manager (https://nixos.org/download)
```

### Download and setup
#### Clone this repository
```
git clone https://github.com/mrm48/diet.git
```

#### Database
```
psql
CREATE DATABASE meal;
GRANT ALL PRIVILEGES ON DATABASE "meal" TO postgres;
```

#### Run (with just)
```
just run
```

#### Run (without just)
```
mvn spring-boot:run
```

If using the nix flake, add nix develop before the run commands
