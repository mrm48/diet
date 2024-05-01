package com.trinsic.diet3.entry;
import java.util.Optional;
import java.util.List;

import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

@Service
public class EntryService{

    private final EntryRepository entryRepository;

    public EntryService(EntryRepository entryRepository){
        this.entryRepository = entryRepository;
    }

	public Entry getFoodEntry(Long meal_id){
        Optional<Entry> food = entryRepository.findEntryByMeal(meal_id);
        if (food.isPresent()){
            return food.get();
        }
        return null;
	}

    @Transactional
    public Entry addFood(Entry requestFood){
        Integer entryStatus = entryRepository.addFoodEntry(requestFood.getMeal_Id(), requestFood.getFood_Id(), requestFood.getCalories());  
        if (entryStatus != 0){
            List<Entry> newEntry = entryRepository.findEntryById(requestFood.getMeal_Id(), requestFood.getFood_Id(), requestFood.getCalories());
            return newEntry.get(0);
        }
        return null;
    }

    @Transactional
    public Entry addFood(Long meal_id, Long food_id, Integer servings, Integer calories){
        Integer entryStatus = entryRepository.addFoodEntry(meal_id, food_id, calories);
        if (entryStatus != 0){
            List<Entry> newEntry = entryRepository.findEntryById(meal_id, food_id, calories);
            if (!newEntry.isEmpty()){
                return newEntry.get(0);
            }
        }
        return null; 
    }

}
