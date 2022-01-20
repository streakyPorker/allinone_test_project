package com.example.learn.anno;

import io.netty.channel.EventLoopGroup;
import lombok.extern.slf4j.Slf4j;
import org.springframework.context.annotation.Profile;
import org.springframework.stereotype.Component;

import javax.annotation.PostConstruct;

@Component
@Slf4j
@Profile("pre")
public class PreBean {

    @PostConstruct
    public void init() {
        log.warn(getClass().getName()+" is loaded in pre env");
    }

    public static void main(String[] args) {

    }

}
