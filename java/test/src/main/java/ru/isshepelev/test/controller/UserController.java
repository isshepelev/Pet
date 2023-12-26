package ru.isshepelev.test.controller;

import lombok.RequiredArgsConstructor;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;
import ru.isshepelev.test.domain.User;
import ru.isshepelev.test.service.UserService;

import java.util.List;

@RestController
@RequiredArgsConstructor
public class UserController {
    private final UserService userService;

    @PostMapping("/save")
    public User saveUser(@RequestBody User user){
        return userService.save(user);
    }
    @GetMapping("/users")
    public List<User> getUsers(){
        return userService.getUsers();
    }
}
