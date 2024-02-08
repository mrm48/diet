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
				- Example request:
				```
				{
    					"name": "Lunch",
    					"dieter": "Matt",
    					"calories": 500    
				}
				```
				- Example response:
				```
				{
    					"id": 7,
    					"name": "Lunch",
    					"day": "2024-02-08",
    					"calories": 500,
    					"dieter": "Matt",
    					"dieterId": 1
				}
				```
		- [x] POST
			- [x] / - create a new meal today
				- Example request:
				```
				{
					"name": "Lunch",
					"dieter": "Matt",
					"calories": 500  
				}
				```
			- [x	] /food - add a food to the latest meal by user
				- Example request:
				```
				{
    					"name": "Cheerios",
    					"dietername": "Matt",
    					"mealname": "Lunch"
				}
				```
	- [x] Dieter
		- [x] GET
			- [x] / - get information on a user
				- Example request:
				```
				{
				    "name": "Matt"
				}
				```
				- Example response:
				```
				{
				    "id": 1,
				    "name": "Matt",
				    "calories": 1500
				}
				```
			- [x] /id - retrive the ID from a user
				- Example request:
				```
				{
				    "name": "Matt"
				}
				```
				- Example response:
				```
				{
				    "id": 1,
				    "name": "Matt",
				    "calories": 1500
				}
				```
			- [x] /calories - get the remaining number of calories for a user
				- Example request:
				```
				{
				    "name": "Matt"
				}
				```
				- Example response:
				```
				{
				    "id": 1,
				    "name": "Matt",
				    "calories": 1000
				}
				```
			- [x] /caloriesToday - get the number of calories used today by a user
				- Example request:
				```
				{
				    "name": "Matt"
				}
				```
				- Example response:
				```
				{
				    "id": 1,
				    "name": "Matt",
				    "calories": 500
				}
				```
		- [x] POST
			- [x] / - create a new user
				- Example request:
				```
				{
				    "name": "Matt",
				    "calories": 1500
				}
				```
			- [x] /calories - set total number of calories by day for a user
				- Example request:
				```
				{
				    "name": "Matt",
				    "calories": 1500
				}
				```

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
