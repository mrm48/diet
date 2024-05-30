# Diet
A REST API that queries and modifies a PostgresQL database using Spring Boot. There is also a nix flake to get started. 

## Endpoints
	1. Food
		- GET
			- /calories - retrieve number of calories by name
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
		- POST
			- / - add a new food
				- Example request:
				```
				{
    					"name": "Plain Bagel",
    					"calories": 270,
    					"units": 1
				}
				```
			- /calories - add calories to a new food
				- Example request:
				```
				{
    					"name": "Beef and Bean Burritos",
    					"calories": 290,
    					"units": 1
				}
				```
		- PUT
			- /calories - update a food's number of calories
				- Example request:
				```
				{
    					"name": "Beef and Bean Burritos",
    					"calories": 290,
    					"units": 1
				}
				```
        - DELETE
            - / - remove a food from the database
            - Example request:
            ```
			{
    			"name": "Beef and Bean Burritos",
    			"calories": 290,
    			"units": 1
			}
			```
	2. Meal
		- GET
			- / - retrieve the most recent meal by user
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
				    "id": 8,
				    "name": "Lunch",
				    "day": "2024-02-11",
				    "calories": 500,
				    "dieter": "Matt",
				    "dieterId": 1
				}
				```
		- POST
			- / - create a new meal today
				- Example request:
				```
				{
					"name": "Lunch",
					"dieter": "Matt",
					"food": ["Cheerios"]
					"calories": 500  
				}
				```
			- /food - add a food to the latest meal by user
				- Example request:
				```
	            {
                    "meal": {
                        "name": "Breakfast",
                        "dieter": "Matt"
                    },
                    "food": {
                        "name": "Cheerios",
                        "calories": 250,
                        "units": 1
                    }
                }
				```
        - DELETE
            - / - remove a meal from the database
                - Example request: 
                ```
                {
                    "name": "Breakfast",
                    "dieter": "Matt"
                }
                ```
			- /food - remove food from the meal by user with name requested
				- Example request:
				```
	            {
                    "meal": {
                        "name": "Breakfast",
                        "dieter": "Matt"
                    },
                    "food": {
                        "name": "Cheerios",
                        "calories": 250,
                        "units": 1
                    }
                }
				```

	3. Dieter
		- GET
			- / - get information on a user
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
			- /id - retrive the ID from a user
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
			- /calories - get the remaining number of calories for a user
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
			- /caloriesToday - get the number of calories used today by a user
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
		- POST
			- / - create a new user
				- Example request:
				```
				{
				    "name": "Matt",
				    "calories": 1500
				}
				```
			- /calories - set total number of calories by day for a user
				- Example request:
				```
				{
				    "name": "Matt",
				    "calories": 1500
				}
				```
        - DELETE
            - / - delete a user
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
	PostgresQL 14.11
	openJDK 21
	Maven 3.9.1
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

#### Note: A Nix Flake is available

Add the following before running if you wish to use the flake.

```
nix develop
```

#### Run (with just)
```
just run
```

#### Run (without just)
```
mvn spring-boot:run
```


