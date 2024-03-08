package com.trinsic.diet3.entry;
import java.util.Optional;

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
            return entryRepository.addFoodEntry(requestFood.getMeal_Id(), requestFood.getFood_Id(), requestFood.getCalories());  
    }

    @Transactional
    public Entry addFood(Long meal_id, Long food_id, Integer servings, Integer calories){
        return entryRepository.addFoodEntry(meal_id, food_id, calories);
    }

}
