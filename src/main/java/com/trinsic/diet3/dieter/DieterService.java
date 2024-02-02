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

    public Dieter getCaloriesByDay(String dieter, LocalDate day){
        Optional<Dieter> namedDieter = dieterRepository.findDieterByName(dieter);
        if(namedDieter.isPresent()){
            Integer currentCalories = mealRepository.findCaloriesByDay(namedDieter.get().getName(), day);
            namedDieter.get().setCalories(currentCalories);
            return namedDieter.get();
        }
        return null;
    }

    public Dieter getRemainingCalories(Dieter dieter){
        LocalDate day = LocalDate.now();
        Optional<Dieter> searchDieter = dieterRepository.findDieterByName(dieter.getName());
        dieter = getCaloriesByDay(dieter.getName(), day);
        Integer usedCalories = mealService.getCaloriesByDay(dieter.getName(), day);
        dieter.setCalories(usedCalories);
        if(searchDieter.isPresent()){
            dieter.setCalories(searchDieter.get().getCalories() - dieter.getCalories());
            return dieter;
        }
        return null;
    }

    public Dieter getID(Dieter dieter){
        Optional<Dieter> searchDieter = dieterRepository.findDieterByName(dieter.getName());
        if (searchDieter.isPresent()){
            return searchDieter.get();
        }
        return null;
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
