package com.trinsic.diet3.entry;

import java.util.Optional;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Modifying;
import org.springframework.data.jpa.repository.Query;
import org.springframework.stereotype.Repository;

@Repository
public interface EntryRepository 
    extends JpaRepository<Entry, Long>{

        @Query("Select f FROM Entry f WHERE f.meal_id = ?1")
        Optional<Entry> findEntryByMeal(Long meal_id);

        @Query("Select f FROM Entry f WHERE f.meal_id = ?1 AND f.food_id = ?2 AND f.calories = ?3")
        Optional<Entry> findEntryById(Long meal_id, Long food_id, Integer calories);

        @Modifying
        @Query("INSERT INTO Entry (food_id, meal_id, calories) VALUES (?1, ?2, ?3)")
        Integer addFoodEntry(Long food_id, Long meal_id, Integer calories);

        @Modifying
        @Query("DELETE FROM Entry WHERE id = ?1")
        Integer removeFoodEntry(Long id);

        @Query("SELECT SUM(calories) FROM Entry f WHERE f.meal_id = ?1")
        Integer findCaloriesByMeal(Long meal_id);
}
