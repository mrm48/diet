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

	public Food listCalories(Food food){
        Optional<Food> searchFood = foodRepository.findFoodByName(food.getName());
        if (searchFood.isPresent()){
            return searchFood.get();
        }
        return null;
	}

	@Transactional
    public Food updateCaloriesByName(Food food){
        Food responseFood = this.listCalories(food);
        if(responseFood != null){
            responseFood.setCalories(food.getCalories());
            foodRepository.addCaloriesByName(responseFood.getName(), responseFood.getCalories());    
            return responseFood;
        }
        return null;
    }


	@Transactional
    public Food addCaloriesByName(Food food){
        Food responseFood = this.listCalories(food);
        if(responseFood != null){
            responseFood.setCalories(food.getCalories());
            foodRepository.addCaloriesByName(responseFood.getName(), responseFood.getCalories());    
            return responseFood;
        }
        else{
            foodRepository.addFood(food.getName(), food.getUnits(), food.getCalories());
        }
        return null;
    }

    @Transactional
    public Food addFood(Food f){
        Optional<Food> searchFood = foodRepository.findFoodByName(f.getName());
        if (searchFood.isEmpty()) {
            foodRepository.addFood(f.getName(), f.getUnits(), f.getCalories());  
            return listCalories(f);
        }
        return null;
    }

}
