package com.trinsic.diet3.dieter;
import java.time.LocalDate;
import java.util.Optional;
import org.json.*;

import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import com.trinsic.diet3.meal.MealRepository;

@Service
public class DieterService{

    private final DieterRepository dieterRepository;
    private final MealRepository mealRepository;

    public DieterService(DieterRepository dieterRepository, MealRepository mealRepository){
        this.dieterRepository = dieterRepository;
        this.mealRepository = mealRepository;
    }

    @Transactional
    public Dieter addDieter(Dieter newDieter){
        Optional<Dieter> searchDieter = dieterRepository.findDieterByName(newDieter.getName());
        if (searchDieter.isEmpty()) {
            dieterRepository.addDieter(newDieter.getName(), newDieter.getCalories());  
            return getID(newDieter);
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

    public Dieter getCaloriesByDay(Dieter dieter, LocalDate day){
        Optional<Dieter> namedDieter = dieterRepository.findDieterByName(dieter.getName());
        Dieter foundDieter = null; 
        if(namedDieter.isPresent()){
            foundDieter = namedDieter.get();
            Integer currentCalories = mealRepository.findCaloriesByDay(foundDieter.getName(), day);
            foundDieter.setCalories(currentCalories);
            return foundDieter;
        }
        return foundDieter;
    }

    public Dieter getRemainingCalories(Dieter dieter){
        LocalDate day = LocalDate.now();
        Optional<Dieter> searchDieter = dieterRepository.findDieterByName(dieter.getName());
        String returnName = dieter.getName();
        Integer totalCalories = searchDieter.get().getCalories();
        Dieter responseDieter = new Dieter();
        responseDieter = getCaloriesByDay(dieter, day);
        if(searchDieter.isPresent()){
            if(responseDieter.getCalories() != null){
                responseDieter.setName(returnName);
                responseDieter.setCalories(totalCalories - responseDieter.getCalories());
            }
            else{
                responseDieter.setName(returnName);
                responseDieter.setCalories(totalCalories);
            }
            return responseDieter;
        }
        return responseDieter;
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
