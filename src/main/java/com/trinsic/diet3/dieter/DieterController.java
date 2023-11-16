package com.trinsic.diet3.dieter;

import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;

import java.util.Optional;

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
        // call the Dieter constructor that uses the current date as the value
        return dieterService.addDieter(new Dieter(hDieter.getName(), hDieter.getTotalCalories()));
    }
}
