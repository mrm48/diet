package com.trinsic.diet3.dieter;

import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;

import java.time.LocalDate;

@RestController
@RequestMapping(path = "api/v1/dieter")
public class DieterController{

    private final DieterService dieterService;

    public DieterController(DieterService dieterService){
        this.dieterService = dieterService;
    }

    @PostMapping("/adddieter")
    @ResponseBody
    public Integer addDieter(@RequestBody Dieter hDieter){
        return dieterService.addDieter(hDieter);
    }

    @PostMapping("/setcalories")
    @ResponseBody
    public Integer setCalories(@RequestBody Dieter hDieter){
        return dieterService.setCalories(hDieter);
    }

    @GetMapping("/getcalories")
    @ResponseBody
    public Integer getCalories(@RequestBody Dieter hDieter){
        return dieterService.getCaloriesByDay(hDieter, LocalDate.now());
    }
}
