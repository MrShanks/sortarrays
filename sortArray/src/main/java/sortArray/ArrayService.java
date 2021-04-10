package sortArray;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.Collections;
import java.util.List;

@Service
public class ArrayService {

    private final ArrayRepository arrayRepository;

    @Autowired
    public ArrayService(ArrayRepository arrayRepository) {this.arrayRepository = arrayRepository;}

    public void addArray(ArrayObject arrayObject) {

        List<Integer> inputVector = arrayObject.getVector();
        Collections.sort(inputVector);
        arrayObject.setVector(inputVector);
        arrayRepository.save(arrayObject);
    }
}
