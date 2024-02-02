package com.trinsic.diet3.dieter;

import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;

import java.time.LocalDate;
import org.json.*;

@RestController
@RequestMapping(path = "api/v1/dieter")
public class DieterController{

    private final DieterService dieterService;

    public DieterController(DieterService dieterService){
        this.dieterService = dieterService;
    }

    @PostMapping("/")
    @ResponseBody
    public Dieter addDieter(@RequestBody Dieter dieter){
        return dieterService.addDieter(dieter);
    }

    @PostMapping("/calories")
    @ResponseBody
    public Dieter setCalories(@RequestBody Dieter dieter){
        return dieterService.setCalories(dieter);
    }

    @GetMapping("/")
    @ResponseBody
    public Dieter get(@RequestBody String dieterName){
        return dieterService.getDieterByName(dieterName);
    }

    @GetMapping("/id")
    @ResponseBody
    public Dieter getID(@RequestBody Dieter dieter){
        return dieterService.getID(dieter);
    }

    @GetMapping("/calories")
    public Dieter getremainingcalories(@RequestBody Dieter dieter){
        return dieterService.getRemainingCalories(dieter);
    }

    @GetMapping("/caloriesToday")
    public Dieter getCaloriesByDay(@RequestBody String dieterName){
        return dieterService.getCaloriesByDay(dieterName, LocalDate.now());
    }
}
