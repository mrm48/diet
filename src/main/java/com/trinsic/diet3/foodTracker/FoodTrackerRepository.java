package com.trinsic.diet3.foodTracker;

import java.util.Optional;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Modifying;
import org.springframework.data.jpa.repository.Query;
import org.springframework.stereotype.Repository;

@Repository
public interface FoodTrackerRepository 
    extends JpaRepository<FoodTracker, Long>{

        @Query("SELECT f FROM foodtracker f WHERE f.meal_id = ?1 AND f.food_id = ?2")
        Optional<FoodTracker> findFoodByMeal(Long meal_id, Long food_id);

}
