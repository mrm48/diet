package com.trinsic.diet3.meal;

import java.time.LocalDate;
import java.util.Optional;

import org.springframework.data.jpa.repository.Modifying;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

/**
 * MealRepository is an interface that interacts with the postgresql database storing meal information
 *
 * @author Matt Miller
 *
 */
@Repository
public interface MealRepository 
        extends JpaRepository<Meal, Long>{

	/**
	 *  Find a meal with the name provided on the provided day with the provided dieterid
	 *  @param name A meal name provided by the user when adding the meal, if it exists
	 *  @param day A LocalDate object representing the day when the meal was entered
	 *  @param dieterid The Long primary key for the dieter in the dieter table
	 *  @return An Optional meal that only isPresent if a meal with the parameters provided is found
	 */
	@Query("SELECT m FROM Meal m WHERE m.name = ?1 AND m.day = ?2 AND m.dieterid = ?3")
	Optional<Meal> findMealByName(String name, LocalDate day, Long dieterid);

	/**
	 *  Find the number of calories consumed during a day by a dieter
	 *  @param name A dieter name
	 *  @param day A LocalDate object representing the day when the meal was entered
	 *  @return Integer representing the number of calories.
	 */
	@Query("SELECT SUM(calories) from Meal m WHERE m.dieter = ?1 AND m.day = ?2")
	Integer findCaloriesByDay(String name, LocalDate day);

	/**
	 *  Find a meal with the name provided on the provided day with the provided dieterid
	 *  @param day A LocalDate object representing the day when the meal was entered
	 *  @param dietername The String name for the dieter in the dieter table
	 *  @param mealname A meal name provided by the user when adding the meal, if it exists
	 *  @return An Optional meal that only isPresent if a meal with the parameters provided is found
	 */
	@Query("SELECT m FROM Meal m WHERE m.day = ?1 AND m.dieter = ?2 AND m.name = ?3")
	Optional<Meal> findMealByDieter(String mealname, LocalDate day, String dietername);

	/**
	 *  Add calories from a food to the meal specified.
	 *  @param cals The number of calories to add
	 *  @param id The meal primary key
	 *  @param name A meal name provided by the user when adding the meal, if it exists
	 *  @param day A LocalDate object representing the day when the meal was entered
	 *  @param dieterid The Long primary key for the dieter in the dieter table
	 *  @param dieter The String dieter name
	 *  @return An Integer representing the status of the update command (updated number of calories)
	 */
	@Modifying
	@Query("UPDATE Meal m SET m.calories = ?1 WHERE m.id = ?2 AND m.name = ?3 AND m.day = ?4 AND m.dieterid = ?5 AND m.dieter = ?6")
	Integer addFood(Integer cals, Long id, String name,  LocalDate day, Long dieterid, String dieter);

	/**
	 *  Add a new meal into the meal table
	 *  @param cals The number of calories to add
	 *  @param name A meal name provided by the user when adding the meal, if it exists
	 *  @param day A LocalDate object representing the day when the meal was entered
	 *  @param dieterid The Long primary key for the dieter in the dieter table
	 *  @param dieter The String dieter name
	 *  @return An Integer representing the status of the insert command
	 */
	@Modifying
	@Query("INSERT INTO Meal (calories, name, day, dieterid, dieter) VALUES (?1, ?2, ?3, ?4, ?5)")
	Integer addMeal(Integer cals, String name, LocalDate day, Long dieterid, String dieter);

	/**
	 * Delete a meal from the database.
	 * @param id Meal primary key from the meal table
	 * @return Integer reporting the status of the delete command from postgres
	 */
	@Modifying
	@Query("DELETE FROM Meal WHERE id = ?1")
	Integer deleteMealById(Long id);
}
