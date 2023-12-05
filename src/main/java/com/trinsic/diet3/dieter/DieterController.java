package com.trinsic.diet3.dieter;

import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;

@RestController
@RequestMapping(path = "api/v1/dieter")
public class DieterController{

    private final DieterService dieterService;

    public DieterController(DieterService dieterService){
        this.dieterService = dieterService;
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

    @GetMapping("/id")
    @ResponseBody
    public Long getID(@RequestBody Dieter dieter){
        return dieterService.getID(dieter);
    }

    @GetMapping("/calories")
    public Integer getremainingcalories(@RequestBody Dieter dieter){
        return dieterService.getRemainingCalories(dieter);
    }
}
