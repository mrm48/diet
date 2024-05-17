package com.trinsic.diet3.dieter;

import org.springframework.web.bind.annotation.*;

import java.time.LocalDate;

/**
 * DieterController class representing the REST API endpoints interacting with the dieter table
 * @author Matt Miller
 */
@RestController
@RequestMapping(path = "api/v1/dieter")
public class DieterController{

    /**
     * Dieter service connecting the controller and the repository.
     */
    private final DieterService dieterService;

    /**
     * Creates a DieterController object with all specified parameters.
     * @param dieterService DieterService object to interact with the repository.
     */
    public DieterController(DieterService dieterService){
        this.dieterService = dieterService;
    }

    /**
     * Creates a new dieter in the database with the Dieter object passed.
     * @param dieter Dieter object to add to the database.
     * @return Dieter object added to the database.
     */
    @PostMapping("/")
    @ResponseBody
    public Dieter addDieter(@RequestBody Dieter dieter){
        return dieterService.addDieter(dieter);
    }

    /**
     * Sets the number of calories for the Dieter object passed, if it exists in the database.
     * @param dieter The Dieter object with the new number of calories specified.
     * @return Dieter object modified or null if the dieter could not be found.
     */
    @PostMapping("/calories")
    @ResponseBody
    public Dieter setCalories(@RequestBody Dieter dieter){
        return dieterService.setCalories(dieter);
    }

    /**
     * Get the dieter object from the database.
     * @param dieter The Dieter object specifying the name
     * @return Dieter object with the target number of calories and id, null if the dieter is not found.
     */
    @GetMapping("/")
    @ResponseBody
    public Dieter get(@RequestBody Dieter dieter){
        return dieterService.getDieterByName(dieter);
    }

    /**
     * Get the dieter ID by the name of the dieter passed.
     * @param dieter The Dieter object specifying the dieter name
     * @return Dieter object with the target number of calories and id, null if the dieter is not found.
     */
    @GetMapping("/id")
    @ResponseBody
    public Dieter getID(@RequestBody Dieter dieter){
        return dieterService.getID(dieter);
    }

    /**
     * Get the remaining number of calories by the name of the dieter passed.
     * @param dieter The Dieter object specifying the dieter name
     * @return Dieter object with the remaining number of calories before hitting the target
     */
    @GetMapping("/calories")
    public Dieter getRemainingCalories(@RequestBody Dieter dieter){
        return dieterService.getRemainingCalories(dieter);
    }

    /**
     * Get the number of calories consumed so far today by the name of the dieter passed.
     * @param dieter The Dieter object specifying the dieter name
     * @return Dieter object with the number of calories consumed
     */
    @GetMapping("/caloriesToday")
    public Dieter getCaloriesByDay(@RequestBody Dieter dieter){
        return dieterService.getCaloriesByDay(dieter, LocalDate.now());
    }

    /**
     * Remove the dieter by name that has been sent in the body of the request
     * @param dieter Dieter with at least the name field specified
     * @return The dieter that was deleted from the database or null if the dieter was not found
     */
    @DeleteMapping("/")
    public Dieter delete(@RequestBody Dieter dieter){
        return dieterService.removeDieterByName(dieter);
    }
}
