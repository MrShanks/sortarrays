package sortArray;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping(path = "api/v1/array/default")
public class ArrayController {

    private final ArrayService arrayService;

    @Autowired
    public ArrayController(ArrayService arrayService) {this.arrayService = arrayService;}

    @PostMapping
    public void addArray(@RequestBody ArrayObject arrayObject){

        arrayService.addArray(arrayObject);
    }

}
