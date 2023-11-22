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
        
        @Query("SELECT m FROM Meal m WHERE m.day = ?1 AND m.dieterid = ?2")
        Optional<Meal> findMealByDay(LocalDate day, Long dieterid);

       @Query("SELECT m FROM Meal m WHERE m.name = ?1 AND m.day = ?2 AND m.dieterid = ?3")
        Optional<Meal> findMealByName(String name, LocalDate day, Long dieterid);

        @Modifying
        @Query("UPDATE Meal m SET m.calories = ?2 WHERE m.name = ?1 AND m.day = ?3 AND m.dieterid = ?4")
        Integer addFood(String name, Integer cals, LocalDate day, Long dieterid);

        @Modifying
        @Query("INSERT INTO Meal (calories, name, day, dieterid) VALUES (?2, ?1, ?3, ?4)")
        Integer addMeal(String name, Integer cals, LocalDate day, Long dieterid);
}
