package com.trinsic.diet3.dieter;

import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;

import com.trinsic.diet3.meal.MealService;

import java.time.LocalDate;

@RestController
@RequestMapping(path = "api/v1/dieter")
public class DieterController{

    private final DieterService dieterService;
    private final MealService mealService;

    public DieterController(DieterService dieterService, MealService mealService){
        this.dieterService = dieterService;
        this.mealService = mealService;
    }

    @PostMapping("/")
    @ResponseBody
    public Integer addDieter(@RequestBody Dieter dieter){
        return dieterService.addDieter(dieter);
    }

    @PostMapping("/calories")
    @ResponseBody
    public Integer setCalories(@RequestBody Dieter dieter){
        return dieterService.setCalories(dieter.getName());
    }

    @GetMapping("/calories")
    @ResponseBody
    public Integer getCalories(@RequestBody Dieter dieter){
        return dieterService.getCaloriesByDay(dieter.getName(), LocalDate.now());
    }

    @GetMapping("/id")
    @ResponseBody
    public Long getID(@RequestBody Dieter dieter){
        return dieterService.getID(dieter);
    }

    @GetMapping("/remaining")
    public Integer getremainingcalories(@RequestBody Dieter dieter){

        String name = dieter.getName();

        Integer usedCalories = mealService.getCaloriesByDay(name, LocalDate.now());
        Integer totalCalories = dieterService.getCaloriesByDay(name, LocalDate.now());

        return totalCalories - usedCalories;
    }
}
