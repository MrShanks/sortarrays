package sortArray;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface ArrayRepository extends JpaRepository<ArrayObject, Long> {

}
