package com.trinsic.diet3.entry;

import java.util.Optional;
import java.time.LocalDate;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Modifying;
import org.springframework.data.jpa.repository.Query;
import org.springframework.stereotype.Repository;

@Repository
public interface EntryRepository 
    extends JpaRepository<Entry, Long>{

        @Query("Select f FROM Entry f WHERE f.meal_id = ?1")
        Optional<Entry> findEntryByMeal(Long meal_id);

        @Modifying
        @Query("INSERT INTO Entry (food_id, meal_id, calories) VALUES (?1, ?2, ?3)")
        Entry addFoodEntry(Long food_id, Long meal_id, Integer calories);

        @Modifying
        @Query("DELETE FROM Entry WHERE id = ?1")
        Entry removeFoodEntry(Long id);

        @Query("SELECT SUM(calories) FROM Entry f WHERE f.meal_id = ?1")
        Integer findCaloriesByMeal(Long meal_id);

        //@Query("SELECT SUM(calories) FROM entry f, meal m")
        //Integer findCaloriesByDieter(String dieter, LocalDate day);
}
