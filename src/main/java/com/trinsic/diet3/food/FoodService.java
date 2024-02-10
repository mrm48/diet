package com.trinsic.diet3.food;
import java.util.Optional;

import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

@Service
public class FoodService{

    private final FoodRepository foodRepository;

    public FoodService(FoodRepository foodRepository){
        this.foodRepository = foodRepository;
    }

	public Food listCalories(Food requestFood){
        Optional<Food> food = foodRepository.findFoodByName(requestFood.getName());
        if (food.isPresent()){
            return food.get();
        }
        return null;
	}

	@Transactional
    public Food updateCaloriesByName(Food requestFood){
        Food responseFood = this.listCalories(requestFood);
        if(responseFood != null){
            responseFood.setCalories(requestFood.getCalories());
            foodRepository.addCaloriesByName(responseFood.getName(), responseFood.getCalories());    
            return responseFood;
        }
        return null;
    }


	@Transactional
    public Food addCaloriesByName(Food requestFood){
        Food responseFood = this.listCalories(requestFood);
        if(responseFood != null){
            responseFood.setCalories(requestFood.getCalories());
            foodRepository.addCaloriesByName(responseFood.getName(), responseFood.getCalories());    
            return responseFood;
        }
        else{
            foodRepository.addFood(requestFood.getName(), requestFood.getUnits(), requestFood.getCalories());
        }
        return null;
    }

    @Transactional
    public Food addFood(Food requestFood){
        Optional<Food> food = foodRepository.findFoodByName(requestFood.getName());
        if (food.isEmpty()) {
            foodRepository.addFood(requestFood.getName(), requestFood.getUnits(), requestFood.getCalories());  
            return listCalories(requestFood);
        }
        return null;
    }

}
