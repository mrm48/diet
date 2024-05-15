package com.trinsic.diet3.food;

import java.util.Optional;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Modifying;
import org.springframework.data.jpa.repository.Query;
import org.springframework.stereotype.Repository;

/**
 * Communicate with the food table
 * @author Matt Miller
 */
@Repository
public interface FoodRepository 
    extends JpaRepository<Food, Long>{

        /**
         * Find the food by string name
         * @param name String representing the food
         * @return The food from the database, only present if the food exists
         */
        @Query("Select f FROM Food f WHERE f.name = ?1")
        Optional<Food> findFoodByName(String name);

        /**
         * Set the number of calories by food name
         * @param name String representing the food's name
         * @param cals The new number of calories
         * @return The status reported by postgres after the update is completed
         */
        @Modifying
        @Query("UPDATE Food f SET f.calories = ?2 WHERE f.name = ?1")
        Integer addCaloriesByName(String name, Integer cals);

        /**
         * Add a new food, specifying all parameters other than the new primary key
         * @param name String representing the name of the food
         * @param units The number of units consumed for the number of calories listed
         * @param cals The number of calories to list for this food
         * @return The status reported by postgres after the insert command is completed.
         */
        @Modifying
        @Query("INSERT INTO Food (calories, units, name) VALUES (?3, ?2, ?1)")
        Integer addFood(String name, Integer units, Integer cals);
}
