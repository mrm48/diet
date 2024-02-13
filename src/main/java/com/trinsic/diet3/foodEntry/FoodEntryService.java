package com.trinsic.diet3.foodEntry;
import java.util.Optional;

import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

@Service
public class FoodEntryService{

    private final FoodEntryRepository foodEntryRepository;

    public FoodEntryService(FoodEntryRepository foodEntryRepository){
        this.foodEntryRepository = foodEntryRepository;
    }

	public FoodEntry getFoodEntry(Long meal_id){
        Optional<FoodEntry> food = foodEntryRepository.findEntryByMeal(meal_id);
        if (food.isPresent()){
            return food.get();
        }
        return null;
	}

    @Transactional
    public FoodEntry addFood(FoodEntry requestFood){
            return foodEntryRepository.addFoodEntry(requestFood.getMeal_Id(), requestFood.getFood_Id(), requestFood.getServings(), requestFood.getCalories());  
    }

    @Transactional
    public FoodEntry addFood(Long meal_id, Long food_id, Integer servings, Integer calories){
        return foodEntryRepository.addFoodEntry(meal_id, food_id, servings, calories);
    }

}
