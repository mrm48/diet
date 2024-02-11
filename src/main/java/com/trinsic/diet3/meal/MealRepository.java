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
        
        @Query("SELECT m FROM Meal m WHERE m.name = ?1 AND m.day = ?2 AND m.dieterid = ?3")
        Optional<Meal> findMealByName(String name, LocalDate day, Long dieterid);

        @Query("SELECT SUM(calories) from Meal m WHERE m.dieter = ?1 AND m.day = ?2")
        Integer findCaloriesByDay(String name, LocalDate day);
        
        @Query("SELECT m FROM Meal m WHERE m.day = ?1 AND m.dieterid = ?2 AND m.name = ?3")
        Optional<Meal> findMealByDay(LocalDate day, Long dieterid, String mealname);

        @Modifying
        @Query("UPDATE Meal m SET m.calories = ?1, m.food = ?7 WHERE m.id = ?2 AND m.name = ?3 AND m.day = ?4 AND m.dieterid = ?5 AND m.dieter = ?6")
        Integer addFood(Integer cals, Long id, String name,  LocalDate day, Long dieterid, String dieter, Object[] food);

        @Modifying
        @Query("INSERT INTO Meal (calories, name, day, dieterid, dieter, food) VALUES (?1, ?2, ?3, ?4, ?5, ?6)")
        Integer addMeal(Integer cals, String name, LocalDate day, Long dieterid, String dieter, Object[] food);
}
