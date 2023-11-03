package com.trinsic.diet3.meal;
import java.util.List;
import org.springframework.stereotype.Service;

@Service
public class MealService{
	public List<Meal> listMeals(){
		return List.of(new Meal());
	}
}
