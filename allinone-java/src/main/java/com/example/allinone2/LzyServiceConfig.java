package com.example.allinone2;

import com.example.allinone2.selfrpc.TestService;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class LzyServiceConfig {

    @Bean
    public TestService getTestService(){
        return new TestService() {
            @Override
            public String foo(String arg1, String arg2) {
                return null;
            }
        };
    }
}
