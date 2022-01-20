package com.example.learn.anno;

import lombok.extern.slf4j.Slf4j;
import org.springframework.context.annotation.Profile;
import org.springframework.stereotype.Component;

import javax.annotation.PostConstruct;

@Component
@Slf4j
@Profile("prod")
public class ProdBean {

    @PostConstruct
    public void init() {
        log.warn(getClass().getName()+" is loaded in prod env");
    }

}