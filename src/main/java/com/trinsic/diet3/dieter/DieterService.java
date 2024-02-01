package com.trinsic.diet3.dieter;
import java.time.LocalDate;
import java.util.Optional;
import org.json.*;

import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import com.trinsic.diet3.meal.MealService;
import com.trinsic.diet3.meal.MealRepository;

@Service
public class DieterService{

    private final DieterRepository dieterRepository;
    private final MealRepository mealRepository;
    private final MealService mealService;

    public DieterService(DieterRepository dieterRepository, MealRepository mealRepository, MealService mealService){
        this.dieterRepository = dieterRepository;
        this.mealRepository = mealRepository;
        this.mealService = mealService;
    }

    @Transactional
    public Dieter addDieter(Dieter newDieter){
        Optional<Dieter> searchDieter = dieterRepository.findDieterByName(newDieter.getName());
        if (searchDieter.isEmpty()) {
            dieterRepository.addDieter(newDieter.getName(), newDieter.getCalories());  
            return newDieter;
        }
        return null;
    }

    @Transactional
    public Dieter setCalories(Dieter dieter){
        Optional<Dieter> searchDieter = dieterRepository.findDieterByName(dieter.getName());
        if (searchDieter.isPresent()){
            dieterRepository.addTotalCalories(dieter.getName(),dieter.getCalories());
            searchDieter.get().setCalories(dieter.getCalories());
            return searchDieter.get();
        }
        return null;
    }

    public Integer getCaloriesByDay(String req, LocalDate day){
        JSONObject requestBody = new JSONObject(req);
        String dieter = requestBody.get("name").toString();
        Optional<Dieter> searchDieter = dieterRepository.findDieterByName(dieter);
        if(searchDieter.isPresent()){
            Optional<Integer> currentCalories = mealRepository.findDieterCaloriesByDay(searchDieter.get().getName(), day);
            if (currentCalories.isPresent()){
                return currentCalories.get();
            }
        }
        return 0;
    }

    public Integer getRemainingCalories(Dieter dieter){
        LocalDate day = LocalDate.now();
        String dieterName = dieter.getName();
        Integer totalCalories = getCaloriesByDay(dieterName, day);
        Integer usedCalories = mealService.getCaloriesByDay(dieterName, day);
        return totalCalories - usedCalories;
    }

    public Long getID(Dieter dieter){
        Optional<Dieter> searchDieter = dieterRepository.findDieterByName(dieter.getName());
        if (searchDieter.isPresent()){
            return searchDieter.get().getId();
        }
        return Long.valueOf(-1);
    }

    public Dieter getDieterByName(String req){
        JSONObject requestBody = new JSONObject(req);
        String dieterName = requestBody.get("name").toString();
        Optional<Dieter> searchDieter = dieterRepository.findDieterByName(dieterName);
        if (searchDieter.isPresent()){
            return searchDieter.get();
        }
        return null;
    }
    

}
