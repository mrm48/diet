package com.trinsic.diet3.foodEntry;

import java.util.Optional;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Modifying;
import org.springframework.data.jpa.repository.Query;
import org.springframework.stereotype.Repository;

@Repository
public interface FoodEntryRepository 
    extends JpaRepository<FoodEntry, Long>{

        @Query("Select f FROM foodentry f WHERE f.meal_id = ?1")
        Optional<FoodEntry> findEntryByMeal(Long meal_id);

        @Modifying
        @Query("INSERT INTO foodentry (food_id, meal_id, servings, calories) VALUES (?1 ?2 ?3 ?4)")
        FoodEntry addFoodEntry(Long food_id, Long meal_id, Integer servings, Integer calories);

        @Modifying
        @Query("DELETE f FROM foodentry f WHERE f.id = ?1")
        FoodEntry removeFoodEntry(Long id);

}
