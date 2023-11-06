package com.trinsic.diet3.meal;

import java.time.LocalDate;
import java.util.Optional;

import org.springframework.data.jpa.repository.Modifying;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface MealRepository 
        extends JpaRepository<Meal, Long>{
        
        @Query("Select f FROM Food f WHERE f.name = ?1 AND f.day = ?3")
        Optional<Meal> findMealByName(String name, LocalDate day);

        @Modifying
        @Query("UPDATE Meal f SET f.calories = ?2 WHERE f.name = ?1 AND f.day = ?3")
        Integer addFood(String name, Integer cals, LocalDate day);

        @Modifying
        @Query("INSERT INTO Meal (calories, name, day) VALUES (?2, ?1, ?3)")
        Integer addMeal(String name, Integer cals, LocalDate day);
}
