package com.trinsic.diet3.dieter;

import java.util.Optional;

import org.springframework.data.jpa.repository.Modifying;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

/**
 * Specifies interactions with the dieter table
 * @author Matt Miller
 */
@Repository
public interface DieterRepository 
        extends JpaRepository<Dieter, Long>{

    /**
     * Returns a dieter object by name if found in the database.
     * @param name The string name of the dieter to find in the database.
     * @return The Dieter if found, the dieter will not be present if it is not.
     */
    @Query("SELECT d FROM Dieter d WHERE d.name = ?1")
    Optional<Dieter> findDieterByName(String name);

    /**
     * Returns the target number of calories for the dieter.
     * @param name The string name of the dieter to find in the database.
     * @return The targetted number of calories for the dieter.
     */
    @Query("SELECT d.calories FROM Dieter d WHERE d.name = ?1")
    Optional<Integer> findDieterCaloriesByDay(String name);

    /**
     * Adds a new dieter with the dieter name and target number of calories
     * @param name The string name of the dieter to create in the database.
     * @param cals The new target number of calories, represented by an Integer.
     * @return The Integer status of the command run on the database to add the new dieter.
     */
    @Modifying
    @Query("INSERT INTO Dieter (name, calories) VALUES (?1, ?2)")
    Integer addDieter(String name, Integer cals);

    /**
     * Updates the target number of calories for this dieter (by name)
     * @param name The string name of the dieter to find in the database.
     * @param cals The new target number of calories, represented by an Integer.
     * @return The Integer status of the command run on the database to set the number of calories.
     */
    @Modifying
    @Query("UPDATE Dieter SET calories = ?2 WHERE name = ?1")
    Integer addTotalCalories(String name, Integer cals);

}