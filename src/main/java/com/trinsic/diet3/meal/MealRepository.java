package com.trinsic.diet3.meal;

import java.util.Optional;

import org.springframework.data.jpa.repository.Modifying;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface MealRepository 
        extends JpaRepository<Meal, Long>{
        
        @Query("Select f FROM Food f WHERE f.name = ?1")
        Optional<Meal> findMealByName(String name);

        @Modifying
        @Query("UPDATE Meal f SET f.calories = ?2 WHERE f.name = ?1")
        Integer addFood(String name, Integer cals);

        @Modifying
        @Query("INSERT INTO Meal (calories, name) VALUES (?3, ?1)")
        Integer addMeal(String name, Integer cals);
}
