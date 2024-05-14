package com.trinsic.diet3.dieter;
import java.time.LocalDate;
import java.util.Optional;

import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import com.trinsic.diet3.meal.MealRepository;

/**
 * DieterService is intended to take requests from DieterController and pass back
 * Dieter objects or Integers.
 * See the {@link com.trinsic.diet3.dieter.Dieter} class for definitions of objects
 * passed
 * back to {@link com.trinsic.diet3.dieter.DieterController}
 * @author Matt Miller
 *
 */
@Service
public class DieterService{

    /**
     * A DieterRepository object to interact with the dieter table
     */
    private final DieterRepository dieterRepository;
    /**
     * A MealRepository object to interact with the meal table
     */
    private final MealRepository mealRepository;

    /**
     *  DieterService constructor accepting fields for each repository type
     *  @param mealRepository A MealRepository object to interact with the meal table
     *  @param dieterRepository A DieterRepository object to interact with the dieter table
     */
    public DieterService(DieterRepository dieterRepository, MealRepository mealRepository){
        this.dieterRepository = dieterRepository;
        this.mealRepository = mealRepository;
    }

    /**
     *  Add a dieter to the database if the name has not already been taken.
     *  @param requestDieter A dieter object to add to the dieter table
     *  @return Dieter added to the database, null if the dieter could not be added (name already exists).
     */
    @Transactional
    public Dieter addDieter(Dieter requestDieter){
        Optional<Dieter> dieter = dieterRepository.findDieterByName(requestDieter.getName());
        if (dieter.isEmpty()) {
            dieterRepository.addDieter(requestDieter.getName(), requestDieter.getCalories());  
            return getID(requestDieter);
        }
        return null;
    }

    /**
     *  Set the number of calories for a dieter.
     *  @param requestDieter A dieter object with the new number of calories.
     *  @return Modified dieter, null if the dieter could not be found.
     */
    @Transactional
    public Dieter setCalories(Dieter requestDieter){
        Optional<Dieter> dieter = dieterRepository.findDieterByName(requestDieter.getName());
        if (dieter.isPresent()){
            dieterRepository.addTotalCalories(requestDieter.getName(),requestDieter.getCalories());
            dieter.get().setCalories(requestDieter.getCalories());
            return dieter.get();
        }
        return null;
    }

    /**
     *  Get the number of calories consumed by a dieter during a day.
     *  @param requestDieter The requested dieter.
     *  @param requestDay The LocalDate object specifying the requested day
     *  @return Modified dieter, null if the dieter could not be found.
     */
    public Dieter getCaloriesByDay(Dieter requestDieter, LocalDate requestDay){
        Optional<Dieter> dieter = dieterRepository.findDieterByName(requestDieter.getName());
        Dieter foundDieter = null; 
        if(dieter.isPresent()){
            foundDieter = dieter.get();
            Integer currentCalories = mealRepository.findCaloriesByDay(foundDieter.getName(), requestDay);
            if (currentCalories != null) {
                foundDieter.setCalories(currentCalories);
            }
            else {
                foundDieter.setCalories(0);
            }
            return foundDieter;
        }
        return foundDieter;
    }

    /**
     *  Get the number of remaining calories for a dieter.
     *  @param requestDieter The requested dieter
     *  @return The dieter with the number of calories set to the remaining calories for the day.
     */
    public Dieter getRemainingCalories(Dieter requestDieter){
        LocalDate day = LocalDate.now();
        String dieterName = requestDieter.getName();
        Optional<Dieter> dieter = dieterRepository.findDieterByName(dieterName);
        Integer totalCalories = dieter.get().getCalories();
        Dieter responseDieter = new Dieter();
        responseDieter = getCaloriesByDay(requestDieter, day);
        if(dieter.isPresent()){
            responseDieter.setName(dieterName);
            if(responseDieter.getCalories() != null){
                responseDieter.setCalories(totalCalories - responseDieter.getCalories());
            }
            else{
                responseDieter.setCalories(totalCalories);
            }
            return responseDieter;
        }
        return responseDieter;
    }

    /**
     *  Get the ID for the dieter added to the database.
     *  @param requestDieter A dieter object with the name field specified.
     *  @return The dieter with the ID field populated, null if the dieter could not be found.
     */
    public Dieter getID(Dieter requestDieter){
        Optional<Dieter> dieter = dieterRepository.findDieterByName(requestDieter.getName());
        if (dieter.isPresent()){
            return dieter.get();
        }
        return null;
    }

    /**
     *  Get the id and target number of calories for the dieter added to the database by name.
     *  @param requestDieter A dieter object with the name field specified.
     *  @return The dieter with the ID and calories fields populated, null if the dieter could not be found.
     */
    public Dieter getDieterByName(Dieter requestDieter){
        Optional<Dieter> dieter = dieterRepository.findDieterByName(requestDieter.getName());
        if (dieter.isPresent()){
            return dieter.get();
        }
        return null;
    }
    

}
