package com.trinsic.diet3.entry;
import java.util.Optional;
import java.util.List;

import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

/**
 * Communicate between the REST API and {@link com.trinsic.diet3.entry.EntryRepository}
 * @author Matt Miller
 */
@Service
public class EntryService{

    /**
     * Communicate with the database
     */
    private final EntryRepository entryRepository;

    /**
     * Constructor initializing the entry repository
     * @param entryRepository Interface for communicating with the entry table
     */
    public EntryService(EntryRepository entryRepository){
        this.entryRepository = entryRepository;
    }

    /**
     * Get the Entry by meal ID
     * @param meal_id The meal primary key from the meal table
     * @return The entry found, null if no entry is found for this meal.
     */
	public Entry getFoodEntry(Long meal_id){
        Optional<Entry> food = entryRepository.findEntryByMeal(meal_id);
        if (food.isPresent()){
            return food.get();
        }
        return null;
	}

    /**
     * Add an entry to the database with the entry object specified
     * @param requestFood The entry to be added to the database
     * @return The entry added to the database, null if it could not be added.
     */
    @Transactional
    public Entry addFood(Entry requestFood){
        Integer entryStatus = entryRepository.addFoodEntry(requestFood.getMeal_Id(), requestFood.getFood_Id(), requestFood.getCalories());  
        if (entryStatus != 0){
            List<Entry> newEntry = entryRepository.findEntryById(requestFood.getMeal_Id(), requestFood.getFood_Id(), requestFood.getCalories());
            return newEntry.getFirst();
        }
        return null;
    }

    /**
     * Add an entry to the database without the need to create a new entry object
     * @param meal_id The primary key of the meal from the meal table
     * @param food_id The primary key of the food from the food table
     * @param servings The number of servings consumed (deprecated)
     * @param calories The number of calories consumed
     * @return The entry object created, null if it could not be created
     */
    @Transactional
    public Entry addFood(Long meal_id, Long food_id, Integer servings, Integer calories){
        Integer entryStatus = entryRepository.addFoodEntry(meal_id, food_id, calories);
        if (entryStatus != 0){
            List<Entry> newEntry = entryRepository.findEntryById(meal_id, food_id, calories);
            if (!newEntry.isEmpty()){
                return newEntry.getFirst();
            }
        }
        return null; 
    }

}
