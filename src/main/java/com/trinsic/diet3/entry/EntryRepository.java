package com.trinsic.diet3.entry;

import java.util.Optional;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Modifying;
import org.springframework.data.jpa.repository.Query;
import org.springframework.stereotype.Repository;
import java.util.List;

/**
 * Interactions with the entry table.
 * @author Matt Miller
 */
@Repository
public interface EntryRepository
    extends JpaRepository<Entry, Long>{

        /**
         * Find the entry related to the meal.
         * @param meal_id The primary key for the meal.
         * @return The entry or entries related to the meal.
         */
        @Query("Select f FROM Entry f WHERE f.meal_id = ?1")
        Optional<Entry> findEntryByMeal(Long meal_id);

        /**
         * Get the entry primary key for the meal, food and number of calories sent.
         * @param meal_id Meal primary key from the meal table.
         * @param food_id Food primary key from the food table.
         * @param calories The number of calories consumed with this entry
         * @return A list of entries matching the parameters sent
         */
        @Query("Select f FROM Entry f WHERE f.meal_id = ?1 AND f.food_id = ?2 AND f.calories = ?3")
        List<Entry> findEntryById(Long meal_id, Long food_id, Integer calories);

        /**
         * Add a new entry by food primary key, meal primary key and number of calories
         * @param food_id Food primary key from the food table
         * @param meal_id Meal primary key from the meal table
         * @param calories Number of calories consumed
         * @return Status returned by Postgres when adding the entry
         */
        @Modifying
        @Query("INSERT INTO Entry (food_id, meal_id, calories) VALUES (?1, ?2, ?3)")
        Integer addFoodEntry(Long food_id, Long meal_id, Integer calories);

        /**
         * Remove a food entry by primary key
         * @param id The primary key of the entry
         * @return The status returned by Postgres when removing the entry
         */
        @Modifying
        @Query("DELETE FROM Entry WHERE id = ?1")
        Integer removeFoodEntry(Long id);

        /**
         * Find the number of calories consumed by meal by summing all entries for that meal id
         * @param meal_id The primary key of the meal in the meal table
         * @return The number of calories consumed
         */
        @Query("SELECT SUM(calories) FROM Entry f WHERE f.meal_id = ?1")
        Integer findCaloriesByMeal(Long meal_id);
}
