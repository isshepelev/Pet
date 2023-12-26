package ru.isshepelev.test.repository;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;
import ru.isshepelev.test.domain.User;

@Repository
public interface UserRepository extends JpaRepository<User,Long> {
}
