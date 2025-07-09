# MauIt

Dieter REST API written in Go using a postgresql database for storing foods, entries and dieter information.

## Endpoints

### Dieter Endpoints

#### GET /dieters/all
Retrieve all dieters from the database.
- **Request**: No body required
- **Response**: Array of dieter objects
```json
[
  {
    "id": 1,
    "name": "Matt",
    "calories": 1500
  }
]
```

#### POST /dieters
Create a new dieter.
- **Request**:
```json
{
  "name": "Matt",
  "calories": 1500
}
```
- **Response**: Created dieter object
```json
{
  "id": 1,
  "name": "Matt",
  "calories": 1500
}
```

#### DELETE /dieters
Delete a dieter from the database.
- **Request**:
```json
{
  "name": "Matt"
}
```
- **Response**: `null` on success

#### POST /dieter/name
Get a specific dieter by name.
- **Request**:
```json
{
  "name": "Matt"
}
```
- **Response**:
```json
{
  "id": 1,
  "name": "Matt",
  "calories": 1500
}
```

#### POST /dieter/calories
Get the target calories for a dieter.
- **Request**:
```json
{
  "name": "Matt"
}
```
- **Response**: Integer value of calories
```
1500
```

#### PUT /dieter/calories
Set the target calories for a dieter.
- **Request**:
```json
{
  "name": "Matt",
  "calories": 1500
}
```
- **Response**:
```json
{
  "id": 1,
  "name": "Matt",
  "calories": 1500
}
```

#### POST /dieter/remaining
Get the remaining calories for a dieter today.
- **Request**:
```json
{
  "name": "Matt"
}
```
- **Response**:
```json
{
  "id": 1,
  "name": "Matt",
  "calories": 1000
}
```

#### POST /dieter/meals
Get all meals for a dieter.
- **Request**:
```json
{
  "name": "Matt"
}
```
- **Response**: Array of meal objects
```json
[
  {
    "id": 1,
    "name": "Lunch",
    "day": "2024-02-11",
    "calories": 500,
    "dieterid": 1,
    "dieter": "Matt"
  }
]
```

#### POST /dieter/mealstoday
Get meals consumed today by a dieter.
- **Request**:
```json
{
  "name": "Matt"
}
```
- **Response**: Array of meal objects for today
```json
[
  {
    "id": 1,
    "name": "Lunch",
    "day": "2024-02-11",
    "calories": 500,
    "dieterid": 1,
    "dieter": "Matt"
  }
]
```

### Meal Endpoints

#### POST /meal
Get a specific meal by name, dieter, and day.
- **Request**:
```json
{
  "name": "Lunch",
  "dieter": "Matt",
  "day": "2024-02-11"
}
```
- **Response**: Array of meal objects
```json
[
  {
    "id": 1,
    "name": "Lunch",
    "day": "2024-02-11",
    "calories": 500,
    "dieterid": 1,
    "dieter": "Matt"
  }
]
```

#### POST /meal/calories
Get the total calories for a specific meal.
- **Request**:
```json
{
  "name": "Lunch",
  "dieter": "Matt",
  "day": "2024-02-11"
}
```
- **Response**:
```json
{
  "id": 1,
  "name": "Lunch",
  "day": "2024-02-11",
  "calories": 500,
  "dieterid": 1,
  "dieter": "Matt"
}
```

#### POST /meal/entries
Get all entries (food items) for a specific meal.
- **Request**:
```json
{
  "name": "Lunch",
  "dieter": "Matt",
  "day": "2024-02-11"
}
```
- **Response**: Array of entry objects
```json
[
  {
    "id": 1,
    "food": 1,
    "meal": 1,
    "calories": 250
  }
]
```

#### PUT /meal/entry
Add a food entry to a meal.
- **Request**:
```json
{
  "food": 1,
  "meal": 1,
  "calories": 250
}
```
- **Response**:
```json
{
  "id": 1,
  "food": 1,
  "meal": 1,
  "calories": 250
}
```

#### PUT /meal
Create a new meal.
- **Request**:
```json
{
  "name": "Lunch",
  "dieter": "Matt",
  "day": "2024-02-11",
  "calories": 500
}
```
- **Response**:
```json
{
  "id": 1,
  "name": "Lunch",
  "day": "2024-02-11",
  "calories": 500,
  "dieterid": 1,
  "dieter": "Matt"
}
```

#### DELETE /meal
Delete a meal.
- **Request**:
```json
{
  "name": "Lunch",
  "dieter": "Matt",
  "day": "2024-02-11"
}
```
- **Response**: `null` on success

#### DELETE /meal/entries
Delete all entries for a meal.
- **Request**:
```json
{
  "name": "Lunch",
  "dieter": "Matt",
  "day": "2024-02-11"
}
```
- **Response**: `null` on success

### Entry Endpoints

#### POST /entry
Get a specific entry by ID.
- **Request**:
```json
{
  "id": 1
}
```
- **Response**:
```json
{
  "id": 1,
  "food": 1,
  "meal": 1,
  "calories": 250
}
```

#### PUT /entry
Create a new entry.
- **Request**:
```json
{
  "food": 1,
  "meal": 1,
  "calories": 250
}
```
- **Response**:
```json
{
  "id": 1,
  "food": 1,
  "meal": 1,
  "calories": 250
}
```

#### DELETE /entry
Delete an entry.
- **Request**:
```json
{
  "id": 1
}
```
- **Response**: `null` on success

### Food Endpoints

#### GET /food/all
Retrieve all foods from the database.
- **Request**: No body required
- **Response**: Array of food objects
```json
[
  {
    "id": 1,
    "name": "Cheerios",
    "calories": 250,
    "units": 1
  }
]
```

#### POST /food
Get a specific food by name.
- **Request**:
```json
{
  "name": "Cheerios"
}
```
- **Response**:
```json
{
  "id": 1,
  "name": "Cheerios",
  "calories": 250,
  "units": 1
}
```

#### PUT /food
Add a new food to the database.
- **Request**:
```json
{
  "name": "Plain Bagel",
  "calories": 270,
  "units": 1
}
```
- **Response**:
```json
{
  "id": 2,
  "name": "Plain Bagel",
  "calories": 270,
  "units": 1
}
```

#### PUT /food/calories
Update a food's calories.
- **Request**:
```json
{
  "name": "Cheerios",
  "calories": 260,
  "units": 1
}
```
- **Response**:
```json
{
  "id": 1,
  "name": "Cheerios",
  "calories": 260,
  "units": 1
}
```

#### DELETE /food
Delete a food from the database.
- **Request**:
```json
{
  "name": "Cheerios"
}
```
- **Response**: `null` on success

## Setup

### Dependencies
```
	PostgresQL 15.4
	Go 1.21 or later
```

### Download and setup
#### Clone this repository
```
git clone https://github.com/mrm48/mauit.git
```

#### Database
```
psql
CREATE DATABASE meal;
GRANT ALL PRIVILEGES ON DATABASE "meal" TO postgres;
```

#### Run
```
go run main.go
```

## Frontend

The application includes a web-based frontend that allows you to:

- Manage users and their daily calorie targets
- Add and manage foods in the database
- Create and track meals
- View calorie consumption and remaining daily calories

### Accessing the Frontend

After starting the application, open your web browser and navigate to:

```
http://localhost:9090
```

### Frontend Features

1. **Dashboard**
   - Select a user to view their calorie summary
   - See daily calorie target, consumed calories, and remaining calories
   - View today's meals

2. **Meals**
   - Add new meals with selected foods
   - View meal history by day
   - Delete meals

3. **Foods**
   - Add new foods with calorie information
   - View all foods in the database
   - Delete foods

4. **Users**
   - Add new users with daily calorie targets
   - View all users
   - Delete users
