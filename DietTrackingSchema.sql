-- Create the dieter table
CREATE TABLE dieter (
    id BIGINT PRIMARY KEY,
    calories INTEGER NOT NULL,
    name VARCHAR(255) NOT NULL UNIQUE
);

-- Create the food table
CREATE TABLE food (
    id BIGINT PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    calories INTEGER NOT NULL,
    units INTEGER NOT NULL
);

-- Create the meal table
CREATE TABLE meal (
    id BIGSERIAL PRIMARY KEY,
    calories INTEGER NOT NULL,
    day VARCHAR(255) NOT NULL,
    dieter VARCHAR(255) NOT NULL,
    dieterid BIGINT NOT NULL,
    name VARCHAR(255) NOT NULL,
    FOREIGN KEY (dieterid) REFERENCES dieter(id)
);

-- Create the entry table
CREATE TABLE entry (
    id BIGINT PRIMARY KEY,
    calories INTEGER NOT NULL,
    food_id BIGINT NOT NULL,
    meal_id BIGINT NOT NULL,
    FOREIGN KEY (food_id) REFERENCES food(id),
    FOREIGN KEY (meal_id) REFERENCES meal(id)
);