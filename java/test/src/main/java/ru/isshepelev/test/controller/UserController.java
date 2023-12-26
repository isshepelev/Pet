package ru.isshepelev.test.controller;

import lombok.RequiredArgsConstructor;
import org.springframework.web.bind.annotation.*;
import ru.isshepelev.test.domain.User;
import ru.isshepelev.test.service.UserService;

import java.util.List;
import java.util.Optional;

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

    @GetMapping("/user/{id}")
    public Optional<User> getUser(@PathVariable Long id){
        return userService.findByUser(id);
    }
}

